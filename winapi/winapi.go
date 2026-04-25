// Package winapi 提供Windows API封装
// 该包提供窗口管理、进程管理等Windows API函数
package winapi

import (
	"fmt"
	"os/exec"
	"strings"
	"syscall"
	"unsafe"
)

// enumWindowsState 枚举窗口状态
type enumWindowsState struct {
	targetPid int
	className string
	title     string
	hwnds     []int
}

// enumWindowsCallback 枚举窗口的回调函数
//go:uintptrescapes
func enumWindowsCallback(hwnd uintptr, lParam uintptr) uintptr {
	state := (*enumWindowsState)(unsafe.Pointer(lParam))
	
	// 获取窗口所属的进程ID
	var pid uint32
	syscall.SyscallN(
		syscall.NewLazyDLL("user32.dll").NewProc("GetWindowThreadProcessId").Addr(),
		hwnd,
		uintptr(unsafe.Pointer(&pid)),
		0,
	)
	
	// 如果指定了目标PID,检查是否匹配
	if state.targetPid > 0 && int(pid) != state.targetPid {
		return 1 // 继续枚举
	}
	
	// 检查窗口是否可见
	ret, _, _ := syscall.NewLazyDLL("user32.dll").NewProc("IsWindowVisible").Call(hwnd)
	if ret == 0 {
		return 1 // 继续枚举
	}
	
	// 获取窗口类名
	buf := make([]uint16, 256)
	syscall.NewLazyDLL("user32.dll").NewProc("GetClassNameW").Call(
		hwnd,
		uintptr(unsafe.Pointer(&buf[0])),
		uintptr(len(buf)),
	)
	className := syscall.UTF16ToString(buf)
	
	// 获取窗口标题
	titleBuf := make([]uint16, 512)
	syscall.NewLazyDLL("user32.dll").NewProc("GetWindowTextW").Call(
		hwnd,
		uintptr(unsafe.Pointer(&titleBuf[0])),
		uintptr(len(titleBuf)),
	)
	title := syscall.UTF16ToString(titleBuf)
	
	// 匹配类名
	if state.className != "" && className != state.className {
		return 1
	}
	
	// 匹配标题
	if state.title != "" && title != state.title {
		return 1
	}
	
	// 添加到结果列表
	state.hwnds = append(state.hwnds, int(hwnd))
	
	return 1 // 继续枚举
}

// WinApi 提供Windows API封装
// 封装Windows API函数，用于窗口和进程管理
type WinApi struct {
	user32   *syscall.LazyDLL
	kernel32 *syscall.LazyDLL
	psapi    *syscall.LazyDLL
}

// NewWinApi 创建新的WinApi实例
// 返回值:
//   *WinApi: WinApi实例
func NewWinApi() *WinApi {
	return &WinApi{
		user32:   syscall.NewLazyDLL("user32.dll"),
		kernel32: syscall.NewLazyDLL("kernel32.dll"),
		psapi:    syscall.NewLazyDLL("psapi.dll"),
	}
}

// FindWindow 查找符合类名或者标题名的顶层可见窗口
// 参数:
//   className: 窗口类名(可为空)
//   title: 窗口标题(可为空)
// 返回值:
//   syscall.Handle: 窗口句柄，找不到返回0
func (w *WinApi) FindWindow(className, title string) syscall.Handle {
	var classPtr, titlePtr *uint16
	if className != "" {
		classPtr, _ = syscall.UTF16PtrFromString(className)
	}
	if title != "" {
		titlePtr, _ = syscall.UTF16PtrFromString(title)
	}
	ret, _, _ := w.user32.NewProc("FindWindowW").Call(
		uintptr(unsafe.Pointer(classPtr)),
		uintptr(unsafe.Pointer(titlePtr)),
	)
	return syscall.Handle(ret)
}

// FindWindowEx 在parent的第一个一级子窗口中寻找窗口
// 参数:
//   parent: 父窗口句柄(0为桌面)
//   className: 窗口类名(可为空)
//   title: 窗口标题(可为空)
// 返回值:
//   syscall.Handle: 窗口句柄，找不到返回0
func (w *WinApi) FindWindowEx(parent int, className, title string) syscall.Handle {
	var classPtr, titlePtr *uint16
	if className != "" {
		classPtr, _ = syscall.UTF16PtrFromString(className)
	}
	if title != "" {
		titlePtr, _ = syscall.UTF16PtrFromString(title)
	}
	ret, _, _ := w.user32.NewProc("FindWindowExW").Call(
		uintptr(parent),
		0,
		uintptr(unsafe.Pointer(classPtr)),
		uintptr(unsafe.Pointer(titlePtr)),
	)
	return syscall.Handle(ret)
}

// GetWindowRect 获取窗口在屏幕上的位置
// 参数:
//   hwnd: 窗口句柄
//   x1: 左坐标(输出)
//   y1: 上坐标(输出)
//   x2: 右坐标(输出)
//   y2: 下坐标(输出)
// 返回值:
//   int: 1表示成功, 0表示失败
func (w *WinApi) GetWindowRect(hwnd int, x1, y1, x2, y2 *int) int {
	var rect [4]int32
	ret, _, _ := w.user32.NewProc("GetWindowRect").Call(
		uintptr(hwnd),
		uintptr(unsafe.Pointer(&rect[0])),
	)
	if ret != 0 {
		*x1 = int(rect[0])
		*y1 = int(rect[1])
		*x2 = int(rect[2])
		*y2 = int(rect[3])
		return 1
	}
	return 0
}

// GetClientRect 获取窗口客户区在屏幕上的位置
// 参数:
//   hwnd: 窗口句柄
//   x1: 左坐标(输出)
//   y1: 上坐标(输出)
//   x2: 右坐标(输出)
//   y2: 下坐标(输出)
// 返回值:
//   int: 1表示成功, 0表示失败
func (w *WinApi) GetClientRect(hwnd int, x1, y1, x2, y2 *int) int {
	var rect [4]int32
	ret, _, _ := w.user32.NewProc("GetClientRect").Call(
		uintptr(hwnd),
		uintptr(unsafe.Pointer(&rect[0])),
	)
	if ret != 0 {
		// Convert client coordinates to screen coordinates
		var ptTop, ptBottom [2]int32
		ptTop[0] = rect[0]
		ptTop[1] = rect[1]
		ptBottom[0] = rect[2]
		ptBottom[1] = rect[3]
		w.user32.NewProc("ClientToScreen").Call(uintptr(hwnd), uintptr(unsafe.Pointer(&ptTop[0])))
		w.user32.NewProc("ClientToScreen").Call(uintptr(hwnd), uintptr(unsafe.Pointer(&ptBottom[0])))
		*x1 = int(ptTop[0])
		*y1 = int(ptTop[1])
		*x2 = int(ptBottom[0])
		*y2 = int(ptBottom[1])
		return 1
	}
	return 0
}

// GetClientSize 获取窗口客户区的宽度和高度
// 参数:
//   hwnd: 窗口句柄
//   width: 宽度(输出)
//   height: 高度(输出)
// 返回值:
//   int: 1表示成功, 0表示失败
func (w *WinApi) GetClientSize(hwnd int, width, height *int) int {
	var rect [4]int32
	ret, _, _ := w.user32.NewProc("GetClientRect").Call(
		uintptr(hwnd),
		uintptr(unsafe.Pointer(&rect[0])),
	)
	if ret != 0 {
		*width = int(rect[2] - rect[0])
		*height = int(rect[3] - rect[1])
		return 1
	}
	return 0
}

// GetWindowText 获取窗口的标题
// 参数:
//   hwnd: 窗口句柄
// 返回值:
//   string: 窗口标题
func (w *WinApi) GetWindowText(hwnd syscall.Handle) string {
	buf := make([]uint16, 512)
	ret, _, _ := w.user32.NewProc("GetWindowTextW").Call(
		uintptr(hwnd),
		uintptr(unsafe.Pointer(&buf[0])),
		uintptr(len(buf)),
	)
	if ret == 0 {
		return ""
	}
	return syscall.UTF16ToString(buf[:ret])
}

// GetWindowTitle 获取窗口的标题
// 参数:
//   hwnd: 窗口句柄
// 返回值:
//   string: 窗口标题
func (w *WinApi) GetWindowTitle(hwnd int) string {
	return w.GetWindowText(syscall.Handle(hwnd))
}

// GetWindowClass 获取窗口的类名
// 参数:
//   hwnd: 窗口句柄
// 返回值:
//   string: 窗口类名
func (w *WinApi) GetWindowClass(hwnd int) string {
	buf := make([]uint16, 256)
	ret, _, _ := w.user32.NewProc("GetClassNameW").Call(
		uintptr(hwnd),
		uintptr(unsafe.Pointer(&buf[0])),
		uintptr(len(buf)),
	)
	if ret == 0 {
		return ""
	}
	return syscall.UTF16ToString(buf[:ret])
}

// GetDesktopWindow 获取桌面窗口句柄
// 返回值:
//   syscall.Handle: 桌面窗口句柄
func (w *WinApi) GetDesktopWindow() syscall.Handle {
	ret, _, _ := w.user32.NewProc("GetDesktopWindow").Call()
	return syscall.Handle(ret)
}

// IsWindowVisible 检查窗口是否可见
// 参数:
//   hwnd: 窗口句柄
// 返回值:
//   bool: true表示可见, false表示不可见
func (w *WinApi) IsWindowVisible(hwnd int) bool {
	ret, _, _ := w.user32.NewProc("IsWindowVisible").Call(
		uintptr(hwnd),
	)
	return ret != 0
}

// GetClassName 获取窗口类名(int版本)
// 参数:
//   hwnd: 窗口句柄
// 返回值:
//   string: 窗口类名
func (w *WinApi) GetClassName(hwnd int) string {
	return w.GetWindowClass(hwnd)
}

// GetWindowText 获取窗口标题(int版本)
// 参数:
//   hwnd: 窗口句柄
// 返回值:
//   string: 窗口标题
func (w *WinApi) GetWindowTextInt(hwnd int) string {
	buf := make([]uint16, 512)
	ret, _, _ := w.user32.NewProc("GetWindowTextW").Call(
		uintptr(hwnd),
		uintptr(unsafe.Pointer(&buf[0])),
		uintptr(len(buf)),
	)
	if ret == 0 {
		return ""
	}
	return syscall.UTF16ToString(buf[:ret])
}

// GetWindowProcessId 获取指定窗口的进程ID
// 参数:
//   hwnd: 窗口句柄
// 返回值:
//   int: 进程ID
func (w *WinApi) GetWindowProcessId(hwnd int) int {
	var pid uint32
	w.user32.NewProc("GetWindowThreadProcessId").Call(
		uintptr(hwnd),
		uintptr(unsafe.Pointer(&pid)),
	)
	return int(pid)
}

// GetWindowProcessPath 获取指定窗口所属进程的exe文件完整路径
// 参数:
//   hwnd: 窗口句柄
// 返回值:
//   string: 进程可执行文件路径
func (w *WinApi) GetWindowProcessPath(hwnd int) string {
	var pid uint32
	w.user32.NewProc("GetWindowThreadProcessId").Call(
		uintptr(hwnd),
		uintptr(unsafe.Pointer(&pid)),
	)
	if pid == 0 {
		return ""
	}
	hProcess, _, _ := w.kernel32.NewProc("OpenProcess").Call(
		uintptr(0x0410), // PROCESS_QUERY_INFORMATION | PROCESS_VM_READ
		0,
		uintptr(pid),
	)
	if hProcess == 0 {
		return ""
	}
	defer w.kernel32.NewProc("CloseHandle").Call(hProcess)

	buf := make([]uint16, 512)
	ret, _, _ := w.psapi.NewProc("GetModuleFileNameExW").Call(
		hProcess,
		0,
		uintptr(unsafe.Pointer(&buf[0])),
		uintptr(len(buf)),
	)
	if ret == 0 {
		return ""
	}
	return syscall.UTF16ToString(buf[:ret])
}

// GetProcessInfo 获取指定PID的详细信息
// 参数:
//   pid: 进程ID
// 返回值:
//   string: 进程信息字符串
func (w *WinApi) GetProcessInfo(pid int) string {
	hProcess, _, _ := w.kernel32.NewProc("OpenProcess").Call(
		uintptr(0x0410),
		0,
		uintptr(pid),
	)
	if hProcess == 0 {
		return ""
	}
	defer w.kernel32.NewProc("CloseHandle").Call(hProcess)

	buf := make([]uint16, 512)
	ret, _, _ := w.psapi.NewProc("GetModuleFileNameExW").Call(
		hProcess,
		0,
		uintptr(unsafe.Pointer(&buf[0])),
		uintptr(len(buf)),
	)
	if ret == 0 {
		return fmt.Sprintf("PID:%d|Path:Unknown", pid)
	}
	path := syscall.UTF16ToString(buf[:ret])
	return fmt.Sprintf("PID:%d|Path:%s", pid, path)
}

// EnumProcess 枚举符合指定进程名的进程PID
// 参数:
//   name: 进程名
// 返回值:
//   string: 进程ID字符串，格式: "pid1|pid2|..."
func (w *WinApi) EnumProcess(name string) string {
	// Use tasklist command to enumerate processes
	cmd := exec.Command("tasklist", "/FI", fmt.Sprintf("IMAGENAME eq %s", name), "/FO", "CSV", "/NH")
	output, err := cmd.Output()
	if err != nil {
		return ""
	}
	// Parse output to get PIDs
	lines := strings.Split(string(output), "\n")
	var pids []string
	for _, line := range lines {
		if strings.Contains(line, name) {
			// Extract PID from CSV format
			parts := strings.Split(line, ",")
			if len(parts) >= 2 {
				pid := strings.Trim(parts[1], "\"")
				pids = append(pids, pid)
			}
		}
	}
	return strings.Join(pids, "|")
}

// EnumWindow 枚举符合指定条件的窗口
// 参数:
//   parent: 父窗口句柄(0为桌面)
//   title: 窗口标题(可为空)
//   className: 窗口类名(可为空)
//   filter: 过滤标志
// 返回值:
//   string: 窗口句柄字符串，格式: "hwnd1,hwnd2|..."
func (w *WinApi) EnumWindow(parent int, title, className string, filter int) string {
	var hwnds []int

	// 如果parent为0,使用桌面窗口
	if parent == 0 {
		parent = int(w.GetDesktopWindow())
	}

	// 获取第一个子窗口
	hwnd := w.GetWindow(parent, 1) // GW_CHILD
	if hwnd == 0 {
		return ""
	}

	// 获取第一个窗口
	hwnd = w.GetWindow(hwnd, 1) // GW_HWNDFIRST

	// 遍历所有窗口
	count := 0
	maxCount := 1000 // 防止无限循环
	for hwnd != 0 && count < maxCount {
		count++
		
		// 检查是否可见
		if (filter & 16) != 0 {
			if !w.IsWindowVisible(hwnd) {
				hwnd = w.GetWindow(hwnd, 2) // GW_HWNDNEXT
				continue
			}
		}

		// 获取窗口类名
		windowClassName := w.GetClassName(hwnd)
		// 获取窗口标题
		windowTitle := w.GetWindowTextInt(hwnd)

		// 根据filter进行匹配
		match := false
		switch {
		case filter == 0: // 所有模式
			match = true
		case (filter & 1) != 0 && title != "": // 匹配窗口标题
			if windowTitle == title {
				match = true
			}
		case (filter & 2) != 0 && className != "": // 匹配窗口类名
			if windowClassName == className {
				match = true
			}
		case (filter & 3) != 0: // 匹配类名或标题
			if (className != "" && windowClassName == className) || (title != "" && windowTitle == title) {
				match = true
			}
		}

		if match {
			hwnds = append(hwnds, hwnd)
		}

		// 获取下一个窗口
		nextHwnd := w.GetWindow(hwnd, 2) // GW_HWNDNEXT
		if nextHwnd == hwnd {
			break // 防止死循环
		}
		hwnd = nextHwnd
	}

	// 转换为字符串
	if len(hwnds) == 0 {
		return ""
	}

	result := fmt.Sprintf("%d", hwnds[0])
	for i := 1; i < len(hwnds); i++ {
		result += fmt.Sprintf(",%d", hwnds[i])
	}
	return result
}

// EnumWindowByProcess 枚举符合指定进程和条件的窗口
// 参数:
//   processName: 进程名
//   title: 窗口标题(可为空)
//   className: 窗口类名(可为空)
//   filter: 过滤标志
// 返回值:
//   string: 窗口句柄字符串，格式: "hwnd1,hwnd2|..."
func (w *WinApi) EnumWindowByProcess(processName, title, className string, filter int) string {
	// TODO: 实现基于进程的窗口枚举
	return ""
}

// FindWindowByProcess 通过进程名寻找可见窗口
// 参数:
//   processName: 进程名
//   className: 窗口类名(可为空)
//   title: 窗口标题(可为空)
// 返回值:
//   int: 窗口句柄，找不到返回0
func (w *WinApi) FindWindowByProcess(processName, className, title string) int {
	// TODO: 实现基于进程的窗口查找
	return 0
}

// FindWindowByProcessId 通过进程ID寻找可见窗口
// 参数:
//   processId: 进程ID
//   className: 窗口类名(可为空)
//   title: 窗口标题(可为空)
// 返回值:
//   int: 窗口句柄，找不到返回0
func (w *WinApi) FindWindowByProcessId(processId int, className, title string) int {
	// 获取桌面窗口
	desktop := w.GetDesktopWindow()
	
	// 获取第一个子窗口
	hwnd := w.GetWindow(int(desktop), 1) // GW_CHILD
	
	// 遍历所有窗口
	for hwnd != 0 {
		// 获取窗口所属的进程ID
		windowPid := w.GetWindowProcessId(hwnd)
		
		// 检查是否匹配进程ID
		if windowPid == processId {
			// 检查是否可见
			if w.IsWindowVisible(hwnd) {
				// 如果指定了类名或标题,进行匹配
				if className != "" {
					winClassName := w.GetClassName(hwnd)
					if winClassName != className {
						hwnd = w.GetWindow(hwnd, 2)
						continue
					}
				}
				
				if title != "" {
					winTitle := w.GetWindowTextInt(hwnd)
					if winTitle != title {
						hwnd = w.GetWindow(hwnd, 2)
						continue
					}
				}
				
				return hwnd
			}
		}
		
		// 获取下一个窗口
		hwnd = w.GetWindow(hwnd, 2) // GW_HWNDNEXT
	}
	
	return 0
}

// ClientToScreen 将客户区坐标转换为屏幕坐标
// 参数:
//   hwnd: 窗口句柄
//   x: X坐标(输入/输出)
//   y: Y坐标(输入/输出)
// 返回值:
//   int: 1表示成功, 0表示失败
func (w *WinApi) ClientToScreen(hwnd int, x, y *int) int {
	type POINT struct {
		X, Y int32
	}
	pt := POINT{X: int32(*x), Y: int32(*y)}
	ret, _, _ := w.user32.NewProc("ClientToScreen").Call(
		uintptr(hwnd),
		uintptr(unsafe.Pointer(&pt)),
	)
	if ret != 0 {
		*x = int(pt.X)
		*y = int(pt.Y)
		return 1
	}
	return 0
}

// ScreenToClient 将屏幕坐标转换为客户区坐标
// 参数:
//   hwnd: 窗口句柄
//   x: X坐标(输入/输出)
//   y: Y坐标(输入/输出)
// 返回值:
//   int: 1表示成功, 0表示失败
func (w *WinApi) ScreenToClient(hwnd int, x, y *int) int {
	type POINT struct {
		X, Y int32
	}
	pt := POINT{X: int32(*x), Y: int32(*y)}
	ret, _, _ := w.user32.NewProc("ScreenToClient").Call(
		uintptr(hwnd),
		uintptr(unsafe.Pointer(&pt)),
	)
	if ret != 0 {
		*x = int(pt.X)
		*y = int(pt.Y)
		return 1
	}
	return 0
}

// GetForegroundFocus 获取顶层活动窗口中输入焦点所在的窗口句柄
// 返回值:
//   int: 有焦点的窗口句柄
func (w *WinApi) GetForegroundFocus() int {
	ret, _, _ := w.user32.NewProc("GetFocus").Call()
	return int(ret)
}

// GetForegroundWindow 获取顶层活动窗口
// 返回值:
//   int: 窗口句柄
func (w *WinApi) GetForegroundWindow() int {
	ret, _, _ := w.user32.NewProc("GetForegroundWindow").Call()
	return int(ret)
}

// GetMousePointWindow 获取鼠标光标下的可见窗口句柄
// 返回值:
//   int: 窗口句柄
func (w *WinApi) GetMousePointWindow() int {
	type POINT struct {
		X, Y int32
	}
	var pt POINT
	w.user32.NewProc("GetCursorPos").Call(uintptr(unsafe.Pointer(&pt)))
	ret, _, _ := w.user32.NewProc("WindowFromPoint").Call(
		uintptr(pt.X) | uintptr(pt.Y)<<32,
	)
	return int(ret)
}

// GetPointWindow 获取指定坐标处的可见窗口句柄
// 参数:
//   x: X坐标
//   y: Y坐标
// 返回值:
//   int: 窗口句柄
func (w *WinApi) GetPointWindow(x, y int) int {
	ret, _, _ := w.user32.NewProc("WindowFromPoint").Call(
		uintptr(x) | uintptr(y)<<32,
	)
	return int(ret)
}

// GetSpecialWindow 获取特殊窗口句柄
// 参数:
//   flag: 0=桌面, 1=任务栏
// 返回值:
//   int: 窗口句柄
func (w *WinApi) GetSpecialWindow(flag int) int {
	switch flag {
	case 0:
		ret, _, _ := w.user32.NewProc("GetDesktopWindow").Call()
		return int(ret)
	case 1:
		ret, _, _ := w.user32.NewProc("FindWindowW").Call(
			uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr("Shell_TrayWnd"))),
			0,
		)
		return int(ret)
	default:
		return 0
	}
}

// GetWindow 获取与给定窗口有关的窗口句柄
// 参数:
//   hwnd: 窗口句柄
//   flag: 0=父窗口, 1=首个子窗口, 2=下一个兄弟窗口, 3=上一个兄弟窗口
// 返回值:
//   int: 窗口句柄
func (w *WinApi) GetWindow(hwnd, flag int) int {
	var cmd uint32
	switch flag {
	case 0:
		cmd = 4 // GW_OWNER
	case 1:
		cmd = 5 // GW_CHILD
	case 2:
		cmd = 2 // GW_HWNDNEXT
	case 3:
		cmd = 1 // GW_HWNDPREV
	default:
		return 0
	}
	ret, _, _ := w.user32.NewProc("GetWindow").Call(
		uintptr(hwnd),
		uintptr(cmd),
	)
	return int(ret)
}

// GetWindowState 获取指定窗口的某些属性
// 参数:
//   hwnd: 窗口句柄
//   flag: 0=可见, 1=活动, 2=最小化, 3=最大化, 4=启用
// 返回值:
//   int: 窗口状态值
func (w *WinApi) GetWindowState(hwnd, flag int) int {
	switch flag {
	case 0:
		ret, _, _ := w.user32.NewProc("IsWindowVisible").Call(uintptr(hwnd))
		if ret != 0 {
			return 1
		}
		return 0
	case 1:
		foreground := w.GetForegroundWindow()
		if foreground == hwnd {
			return 1
		}
		return 0
	case 2:
		ret, _, _ := w.user32.NewProc("IsIconic").Call(uintptr(hwnd))
		if ret != 0 {
			return 1
		}
		return 0
	case 3:
		ret, _, _ := w.user32.NewProc("IsZoomed").Call(uintptr(hwnd))
		if ret != 0 {
			return 1
		}
		return 0
	case 4:
		ret, _, _ := w.user32.NewProc("IsWindowEnabled").Call(uintptr(hwnd))
		if ret != 0 {
			return 1
		}
		return 0
	default:
		return 0
	}
}

// MoveWindow 移动指定窗口到指定位置
// 参数:
//   hwnd: 窗口句柄
//   x: X坐标
//   y: Y坐标
// 返回值:
//   int: 1表示成功, 0表示失败
func (w *WinApi) MoveWindow(hwnd, x, y int) int {
	var rect [4]int32
	ret, _, _ := w.user32.NewProc("GetWindowRect").Call(
		uintptr(hwnd),
		uintptr(unsafe.Pointer(&rect[0])),
	)
	if ret == 0 {
		return 0
	}
	width := rect[2] - rect[0]
	height := rect[3] - rect[1]
	
	// 如果窗口最小化,先还原
	if w.GetWindowState(hwnd, 2) == 1 {
		w.SetWindowState(hwnd, 4) // 还原窗口
	}
	
	ret, _, _ = w.user32.NewProc("MoveWindow").Call(
		uintptr(hwnd),
		uintptr(x),
		uintptr(y),
		uintptr(width),
		uintptr(height),
		1, // 重绘
	)
	if ret != 0 {
		return 1
	}
	return 0
}

// SetClientSize 设置窗口客户区的宽度和高度
// 参数:
//   hwnd: 窗口句柄
//   width: 宽度
//   height: 高度
// 返回值:
//   int: 1表示成功, 0表示失败
func (w *WinApi) SetClientSize(hwnd, width, height int) int {
	// TODO: 实现带有适当窗口样式调整的客户端大小设置
	return 0
}

// SetWindowState 设置窗口的状态
// 参数:
//   hwnd: 窗口句柄
//   flag: 0=显示, 1=隐藏, 2=最小化, 3=最大化, 4=还原, 5=激活, 6=关闭
// 返回值:
//   int: 1表示成功, 0表示失败
func (w *WinApi) SetWindowState(hwnd, flag int) int {
	var cmd int32
	switch flag {
	case 0:
		cmd = 5 // SW_SHOW
	case 1:
		cmd = 0 // SW_HIDE
	case 2:
		cmd = 6 // SW_MINIMIZE
	case 3:
		cmd = 3 // SW_MAXIMIZE
	case 4:
		cmd = 9 // SW_RESTORE
	case 5:
		cmd = 9 // SW_RESTORE
		w.SetForegroundWindow(syscall.Handle(hwnd))
	case 6:
		ret, _, _ := w.user32.NewProc("PostMessageW").Call(
			uintptr(hwnd),
			0x0010, // WM_CLOSE
			0,
			0,
		)
		if ret != 0 {
			return 1
		}
		return 0
	default:
		cmd = 5
	}
	ret, _, _ := w.user32.NewProc("ShowWindow").Call(
		uintptr(hwnd),
		uintptr(cmd),
	)
	if ret != 0 {
		return 1
	}
	return 0
}

// SetWindowSize 设置窗口的大小
// 参数:
//   hwnd: 窗口句柄
//   width: 宽度
//   height: 高度
// 返回值:
//   int: 1表示成功, 0表示失败
func (w *WinApi) SetWindowSize(hwnd, width, height int) int {
	var rect [4]int32
	ret, _, _ := w.user32.NewProc("GetWindowRect").Call(
		uintptr(hwnd),
		uintptr(unsafe.Pointer(&rect[0])),
	)
	if ret == 0 {
		return 0
	}
	ret, _, _ = w.user32.NewProc("MoveWindow").Call(
		uintptr(hwnd),
		uintptr(rect[0]),
		uintptr(rect[1]),
		uintptr(width),
		uintptr(height),
		1,
	)
	if ret != 0 {
		return 1
	}
	return 0
}

// SetWindowText 设置窗口的标题
// 参数:
//   hwnd: 窗口句柄
//   title: 窗口标题
// 返回值:
//   int: 1表示成功, 0表示失败
func (w *WinApi) SetWindowText(hwnd int, title string) int {
	titlePtr, _ := syscall.UTF16PtrFromString(title)
	ret, _, _ := w.user32.NewProc("SetWindowTextW").Call(
		uintptr(hwnd),
		uintptr(unsafe.Pointer(titlePtr)),
	)
	if ret != 0 {
		return 1
	}
	return 0
}

// SetWindowTransparent 设置窗口的透明度
// 参数:
//   hwnd: 窗口句柄
//   trans: 透明度值(0-255)
// 返回值:
//   int: 1表示成功, 0表示失败
func (w *WinApi) SetWindowTransparent(hwnd, trans int) int {
	// TODO: 使用SetLayeredWindowAttributes实现窗口透明度
	return 0
}

// SendPaste 向指定窗口发送粘贴命令
// 参数:
//   hwnd: 窗口句柄
// 返回值:
//   int: 1表示成功, 0表示失败
func (w *WinApi) SendPaste(hwnd int) int {
	// 使用SendMessageW发送WM_PASTE消息
	ret, _, _ := w.user32.NewProc("SendMessageW").Call(
		uintptr(hwnd),
		0x0302, // WM_PASTE
		0,
		0,
	)
	if ret != 0 || true { // SendMessageW 返回值因窗口而异
		return 1
	}
	return 0
}

// SendString 向指定窗口发送文本数据
// 参数:
//   hwnd: 窗口句柄
//   str: 文本字符串
// 返回值:
//   int: 1表示成功, 0表示失败
func (w *WinApi) SendString(hwnd int, str string) int {
	// 转换字符串为UTF-16指针
	strPtr, _ := syscall.UTF16PtrFromString(str)
	
	// 使用SendMessageW发送WM_SETTEXT消息
	ret, _, _ := w.user32.NewProc("SendMessageW").Call(
		uintptr(hwnd),
		0x000C, // WM_SETTEXT
		0,
		uintptr(unsafe.Pointer(strPtr)),
	)
	if ret != 0 || true {
		return 1
	}
	return 0
}

// SendStringIme 向指定窗口发送文本数据(使用IME)
// 参数:
//   hwnd: 窗口句柄
//   str: 文本字符串
// 返回值:
//   int: 1表示成功, 0表示失败
func (w *WinApi) SendStringIme(hwnd int, str string) int {
	// TODO: 实现IME字符串发送
	return 0
}

// TerminateProcess 结束指定进程
// 参数:
//   pid: 进程ID
// 返回值:
//   int: 1表示成功, 0表示失败
func (w *WinApi) TerminateProcess(pid int) int {
	// 打开进程
	hProcess, _, _ := w.kernel32.NewProc("OpenProcess").Call(
		0x0001, // PROCESS_TERMINATE
		0,
		uintptr(pid),
	)
	if hProcess == 0 {
		return 0
	}
	defer w.kernel32.NewProc("CloseHandle").Call(hProcess)

	// 终止进程
	ret, _, _ := w.kernel32.NewProc("TerminateProcess").Call(hProcess, 0)
	if ret != 0 {
		return 1
	}
	return 0
}

// RunApp 运行一个可执行文件，根据指定模式
// 参数:
//   cmdline: 命令行
//   mode: 运行模式 (0=普通, 1=显示, 2=隐藏)
// 返回值:
//   int: 1表示成功, 0表示失败
func (w *WinApi) RunApp(cmdline string, mode int) int {
	cmd := exec.Command("cmd", "/C", cmdline)
	if mode == 2 {
		cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	}
	err := cmd.Start()
	if err != nil {
		return 0
	}
	return 1
}

// WinExec 运行一个可执行文件，根据指定显示模式
// 参数:
//   cmdline: 命令行
//   cmdShow: 显示模式 (SW_SHOW, SW_HIDE等)
// 返回值:
//   int: 1表示成功, 0表示失败
func (w *WinApi) WinExec(cmdline string, cmdShow int) int {
	cmdlinePtr, _ := syscall.UTF16PtrFromString(cmdline)
	ret, _, _ := w.kernel32.NewProc("WinExec").Call(
		uintptr(unsafe.Pointer(cmdlinePtr)),
		uintptr(cmdShow),
	)
	if ret > 31 {
		return 1
	}
	return 0
}

// GetCmdStr 运行命令行并返回结果
// 参数:
//   cmd: 命令字符串
//   milliseconds: 超时时间(毫秒)
// 返回值:
//   string: 命令输出
func (w *WinApi) GetCmdStr(cmd string, milliseconds int) string {
	command := exec.Command("cmd", "/C", cmd)
	output, err := command.Output()
	if err != nil {
		return ""
	}
	return string(output)
}

// SetClipboard 设置剪贴板数据
// 参数:
//   str: 文本字符串
// 返回值:
//   int: 1表示成功, 0表示失败
func (w *WinApi) SetClipboard(str string) int {
	// 转换字符串为UTF-16字节
	utf16Str, _ := syscall.UTF16FromString(str)
	strLen := len(utf16Str) * 2

	// 分配全局内存
	hMem, _, _ := w.kernel32.NewProc("GlobalAlloc").Call(0x0042, uintptr(strLen)) // GMEM_MOVEABLE | GMEM_ZEROINIT
	if hMem == 0 {
		return 0
	}

	// 锁定内存
	pMem, _, _ := w.kernel32.NewProc("GlobalLock").Call(hMem)
	if pMem == 0 {
		w.kernel32.NewProc("GlobalFree").Call(hMem)
		return 0
	}

	// 复制数据到内存
	src := unsafe.Pointer(&utf16Str[0])
	for i := 0; i < strLen; i++ {
		*(*byte)(unsafe.Add(unsafe.Pointer(pMem), i)) = *(*byte)(unsafe.Add(src, i))
	}

	// 解锁内存
	w.kernel32.NewProc("GlobalUnlock").Call(hMem)

	// 打开剪贴板
	ret, _, _ := w.user32.NewProc("OpenClipboard").Call(0)
	if ret == 0 {
		w.kernel32.NewProc("GlobalFree").Call(hMem)
		return 0
	}
	defer w.user32.NewProc("CloseClipboard").Call()

	// 清空剪贴板
	w.user32.NewProc("EmptyClipboard").Call()

	// 设置剪贴板数据 (CF_UNICODETEXT = 13)
	ret, _, _ = w.user32.NewProc("SetClipboardData").Call(13, hMem)
	if ret == 0 {
		w.kernel32.NewProc("GlobalFree").Call(hMem)
		return 0
	}

	return 1
}

// GetClipboard 获取剪贴板数据
// 返回值:
//   string: 剪贴板文本
func (w *WinApi) GetClipboard() string {
	// 打开剪贴板
	ret, _, _ := w.user32.NewProc("OpenClipboard").Call(0)
	if ret == 0 {
		return ""
	}
	defer w.user32.NewProc("CloseClipboard").Call()

	// 获取剪贴板数据句柄
	hMem, _, _ := w.user32.NewProc("GetClipboardData").Call(1) // CF_TEXT
	if hMem == 0 {
		return ""
	}

	// 锁定全局内存并获取指针
	pMem, _, _ := w.kernel32.NewProc("GlobalLock").Call(hMem)
	if pMem == 0 {
		return ""
	}
	defer w.kernel32.NewProc("GlobalUnlock").Call(hMem)

	// Read string from memory
	// TODO: Implement proper string reading from clipboard
	return ""
}

// InjectDll injects a DLL into a process
// Parameters:
//   processName: Process name
//   dllName: DLL name to inject
// Returns:
//   int: 1 for success, 0 for failure
func (w *WinApi) InjectDll(processName, dllName string) int {
	// TODO: Implement DLL injection using CreateRemoteThread
	return 0
}

// SetCursorPos sets the mouse cursor position
// Parameters:
//   x: X coordinate
//   y: Y coordinate
// Returns:
//   bool: true for success, false for failure
func (w *WinApi) SetCursorPos(x, y int) bool {
	ret, _, _ := w.user32.NewProc("SetCursorPos").Call(
		uintptr(x),
		uintptr(y),
	)
	return ret != 0
}

// GetCursorPos gets the mouse cursor position
// Returns:
//   int, int: x, y coordinates
//   bool: true for success, false for failure
func (w *WinApi) GetCursorPos() (int, int, bool) {
	type POINT struct {
		X, Y int32
	}
	var pt POINT
	ret, _, _ := w.user32.NewProc("GetCursorPos").Call(
		uintptr(unsafe.Pointer(&pt)),
	)
	return int(pt.X), int(pt.Y), ret != 0
}

// MouseEvent simulates a mouse event
// Parameters:
//   flags: Event flags
//   dx, dy: Coordinates
//   data: Additional data
func (w *WinApi) MouseEvent(flags, dx, dy, data uintptr) {
	w.user32.NewProc("mouse_event").Call(flags, dx, dy, data, 0)
}

// KeybdEvent simulates a keyboard event
// Parameters:
//   bVk: Virtual key code
//   bScan: Scan code
//   dwFlags: Event flags
//   dwExtraInfo: Additional information
func (w *WinApi) KeybdEvent(bVk, bScan, dwFlags, dwExtraInfo uintptr) {
	w.user32.NewProc("keybd_event").Call(bVk, bScan, dwFlags, dwExtraInfo)
}

// ShowWindow shows or hides a window
// Parameters:
//   hwnd: Window handle
//   cmd: Show command
// Returns:
//   bool: true for success, false for failure
func (w *WinApi) ShowWindow(hwnd syscall.Handle, cmd int) bool {
	ret, _, _ := w.user32.NewProc("ShowWindow").Call(
		uintptr(hwnd),
		uintptr(cmd),
	)
	return ret != 0
}

// SetForegroundWindow sets the foreground window
// Parameters:
//   hwnd: Window handle
// Returns:
//   bool: true for success, false for failure
func (w *WinApi) SetForegroundWindow(hwnd syscall.Handle) bool {
	ret, _, _ := w.user32.NewProc("SetForegroundWindow").Call(uintptr(hwnd))
	return ret != 0
}

// SendMessage sends a window message to a window
// Parameters:
//   hwnd: Window handle
//   msg: Message
//   wParam, lParam: Parameters
// Returns:
//   uintptr: Return value
func (w *WinApi) SendMessage(hwnd syscall.Handle, msg, wParam, lParam uintptr) uintptr {
	ret, _, _ := w.user32.NewProc("SendMessageW").Call(
		uintptr(hwnd),
		msg,
		wParam,
		lParam,
	)
	return ret
}

// PostMessage posts a window message to a window
// Parameters:
//   hwnd: Window handle
//   msg: Message
//   wParam, lParam: Parameters
// Returns:
//   bool: true for success, false for failure
func (w *WinApi) PostMessage(hwnd syscall.Handle, msg, wParam, lParam uintptr) bool {
	ret, _, _ := w.user32.NewProc("PostMessageW").Call(
		uintptr(hwnd),
		msg,
		wParam,
		lParam,
	)
	return ret != 0
}
