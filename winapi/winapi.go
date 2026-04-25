// Package winapi 提供Windows API封装
// 该包提供了窗口管理、进程管理等Windows API的Go封装
package winapi

import (
	"syscall"
	"unsafe"
)

// WinApi Windows API封装类
// 提供窗口、进程等Windows API的封装
type WinApi struct {
	user32   *syscall.LazyDLL
	kernel32 *syscall.LazyDLL
}

// NewWinApi 创建WinApi实例
// 返回值:
//
//	*WinApi: WinApi实例
func NewWinApi() *WinApi {
	return &WinApi{
		user32:   syscall.NewLazyDLL("user32.dll"),
		kernel32: syscall.NewLazyDLL("kernel32.dll"),
	}
}

// FindWindow 查找顶层窗口
// 参数:
//
//	className: 窗口类名（可为空）
//	title: 窗口标题（可为空）
//
// 返回值:
//
//	syscall.Handle: 窗口句柄，未找到返回0
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

// GetWindowRect 获取窗口在屏幕上的位置
// 参数:
//
//	hwnd: 窗口句柄
//
// 返回值:
//
//	int, int, int, int: x1, y1, x2, y2
//	bool: 是否成功
func (w *WinApi) GetWindowRect(hwnd syscall.Handle) (int, int, int, int, bool) {
	var rect [4]int32
	ret, _, _ := w.user32.NewProc("GetWindowRect").Call(
		uintptr(hwnd),
		uintptr(unsafe.Pointer(&rect[0])),
	)
	return int(rect[0]), int(rect[1]), int(rect[2]), int(rect[3]), ret != 0
}

// GetWindowText 获取窗口标题
// 参数:
//
//	hwnd: 窗口句柄
//
// 返回值:
//
//	string: 窗口标题
func (w *WinApi) GetWindowText(hwnd syscall.Handle) string {
	buf := make([]uint16, 256)
	w.user32.NewProc("GetWindowTextW").Call(
		uintptr(hwnd),
		uintptr(unsafe.Pointer(&buf[0])),
		uintptr(len(buf)),
	)
	return syscall.UTF16ToString(buf)
}

// SetCursorPos 设置鼠标位置
// 参数:
//
//	x, y: 坐标
//
// 返回值:
//
//	bool: 是否成功
func (w *WinApi) SetCursorPos(x, y int) bool {
	ret, _, _ := w.user32.NewProc("SetCursorPos").Call(
		uintptr(x),
		uintptr(y),
	)
	return ret != 0
}

// GetCursorPos 获取鼠标位置
// 返回值:
//
//	int, int: x, y坐标
//	bool: 是否成功
func (w *WinApi) GetCursorPos() (int, int, bool) {
	var point [2]int32
	ret, _, _ := w.user32.NewProc("GetCursorPos").Call(
		uintptr(unsafe.Pointer(&point[0])),
	)
	return int(point[0]), int(point[1]), ret != 0
}

// MouseEvent 模拟鼠标事件
// 参数:
//
//	flags: 事件标志
//	dx, dy: 坐标
//	data: 额外数据
func (w *WinApi) MouseEvent(flags, dx, dy, data uintptr) {
	w.user32.NewProc("mouse_event").Call(flags, dx, dy, data, 0)
}

// KeybdEvent 模拟键盘事件
// 参数:
//
//	bVk: 虚拟键码
//	bScan: 扫描码
//	dwFlags: 事件标志
//	dwExtraInfo: 额外信息
func (w *WinApi) KeybdEvent(bVk, bScan, dwFlags, dwExtraInfo uintptr) {
	w.user32.NewProc("keybd_event").Call(bVk, bScan, dwFlags, dwExtraInfo)
}

// ShowWindow 显示/隐藏窗口
// 参数:
//
//	hwnd: 窗口句柄
//	cmd: 显示命令
//
// 返回值:
//
//	bool: 是否成功
func (w *WinApi) ShowWindow(hwnd syscall.Handle, cmd int) bool {
	ret, _, _ := w.user32.NewProc("ShowWindow").Call(
		uintptr(hwnd),
		uintptr(cmd),
	)
	return ret != 0
}

// SetForegroundWindow 设置前台窗口
// 参数:
//
//	hwnd: 窗口句柄
//
// 返回值:
//
//	bool: 是否成功
func (w *WinApi) SetForegroundWindow(hwnd syscall.Handle) bool {
	ret, _, _ := w.user32.NewProc("SetForegroundWindow").Call(uintptr(hwnd))
	return ret != 0
}

// SendMessage 发送窗口消息
// 参数:
//
//	hwnd: 窗口句柄
//	msg: 消息
//	wParam, lParam: 参数
//
// 返回值:
//
//	uintptr: 返回值
func (w *WinApi) SendMessage(hwnd syscall.Handle, msg, wParam, lParam uintptr) uintptr {
	ret, _, _ := w.user32.NewProc("SendMessageW").Call(
		uintptr(hwnd),
		msg,
		wParam,
		lParam,
	)
	return ret
}

// PostMessage 投递窗口消息
// 参数:
//
//	hwnd: 窗口句柄
//	msg: 消息
//	wParam, lParam: 参数
//
// 返回值:
//
//	bool: 是否成功
func (w *WinApi) PostMessage(hwnd syscall.Handle, msg, wParam, lParam uintptr) bool {
	ret, _, _ := w.user32.NewProc("PostMessageW").Call(
		uintptr(hwnd),
		msg,
		wParam,
		lParam,
	)
	return ret != 0
}
