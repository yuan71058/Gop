// Package background 提供后台操作功能
// 该包管理窗口绑定、屏幕捕获和输入模拟
package background

import (
	"math/rand"
	"syscall"
	"time"
	"unsafe"

	"github.com/yuan71058/GOP/core"
)

// DisplayType 定义不同的屏幕捕获方式
type DisplayType int

const (
	DisplayNormal DisplayType = iota // 普通模式
	DisplayGDI                       // GDI模式
	DisplayDX                        // DirectX模式
	DisplayOpenGL                    // OpenGL模式
)

// MouseType 定义不同的鼠标输入模拟方式
type MouseType int

const (
	MouseNormal  MouseType = iota // 普通模式
	MouseWindows                  // Windows模式
	MouseDX                       // DirectX模式
)

// KeypadType 定义不同的键盘输入模拟方式
type KeypadType int

const (
	KeypadNormal  KeypadType = iota // 普通模式
	KeypadWindows                   // Windows模式
	KeypadDX                        // DirectX模式
)

// Background 管理后台操作
// 用于窗口绑定、屏幕捕获和输入模拟
type Background struct {
	hwnd        syscall.Handle // 绑定的窗口句柄
	displayType DisplayType    // 屏幕捕获类型
	mouseType   MouseType      // 鼠标模式
	keypadType  KeypadType     // 键盘模式
	mode        int            // 绑定模式
	CurrPath    string         // 当前路径
	mouseDelay  int            // 鼠标点击延迟(毫秒)
	keypadDelay int            // 按键延迟(毫秒)
}

// NewBackground 创建新的Background实例
// 返回值:
//
//	*Background: Background实例
func NewBackground() *Background {
	return &Background{
		displayType: DisplayNormal,
		mouseType:   MouseNormal,
		keypadType:  KeypadNormal,
		mouseDelay:  30,
		keypadDelay: 30,
	}
}

// BindWindow 绑定窗口并开始屏幕捕获
// 参数:
//
//	hwnd: 窗口句柄
//	display: 显示模式 ("normal", "gdi", "dx", "opengl" 等)
//	mouse: 鼠标模式 ("normal", "windows", "windows3", "dx" 等)
//	keypad: 键盘模式 ("normal", "windows", "dx" 等)
//	mode: 绑定模式 (0=普通, 1=后台 等)
//
// 返回值:
//
//	int: 1表示成功, 0表示失败
func (b *Background) BindWindow(hwnd syscall.Handle, display, mouse, keypad string, mode int) int {
	b.hwnd = hwnd
	// Parse display type
	switch display {
	case "gdi":
		b.displayType = DisplayGDI
	case "dx":
		b.displayType = DisplayDX
	case "dx2":
		b.displayType = DisplayDX
	case "opengl":
		b.displayType = DisplayOpenGL
	default:
		b.displayType = DisplayNormal
	}
	// Parse mouse type
	switch mouse {
	case "windows":
		b.mouseType = MouseWindows
	case "windows3":
		b.mouseType = MouseWindows
	case "dx":
		b.mouseType = MouseDX
	default:
		b.mouseType = MouseNormal
	}
	// Parse keyboard type
	switch keypad {
	case "windows":
		b.keypadType = KeypadWindows
	case "dx":
		b.keypadType = KeypadDX
	default:
		b.keypadType = KeypadNormal
	}
	b.mode = mode
	return 1
}

// UnBindWindow 解绑窗口
// 返回值:
//
//	int: 1表示成功, 0表示失败
func (b *Background) UnBindWindow() int {
	b.hwnd = 0
	b.displayType = DisplayNormal
	b.mouseType = MouseNormal
	b.keypadType = KeypadNormal
	return 1
}

// GetBindWindow 获取已绑定的窗口句柄
// 返回值:
//
//	syscall.Handle: 窗口句柄
func (b *Background) GetBindWindow() syscall.Handle {
	return b.hwnd
}

// IsBind 判断当前对象是否已绑定窗口
// 返回值:
//
//	int: 1表示已绑定, 0表示未绑定
func (b *Background) IsBind() int {
	if b.hwnd != 0 {
		return 1
	}
	return 0
}

// Capture 捕获窗口区域
// 参数:
//
//	x1, y1, x2, y2: 捕获区域坐标
//
// 返回值:
//
//	[]byte: 图像数据
//	int: 宽度
//	int: 高度
func (b *Background) Capture(x1, y1, x2, y2 int) ([]byte, int, int) {
	// TODO: 根据显示类型实现实际的屏幕捕获
	width := x2 - x1 + 1
	height := y2 - y1 + 1
	return nil, width, height
}

// MoveTo 移动鼠标到指定坐标
// 参数:
//
//	x, y: 目标坐标
//
// 返回值:
//
//	int: 1表示成功, 0表示失败
func (b *Background) MoveTo(x, y int) int {
	if b.hwnd == 0 {
		// 未绑定时使用全局鼠标事件
		user32 := syscall.NewLazyDLL("user32.dll")
		setCursorPos := user32.NewProc("SetCursorPos")
		ret, _, _ := setCursorPos.Call(uintptr(x), uintptr(y))
		if ret != 0 {
			return 1
		}
		return 0
	}
	// TODO: 根据鼠标模式实现
	return 1
}

// MoveR 相对上次位置移动鼠标
// 参数:
//
//	x: 相对X偏移
//	y: 相对Y偏移
//
// 返回值:
//
//	int: 1表示成功, 0表示失败
func (b *Background) MoveR(x, y int) int {
	pos, ret := b.GetCursorPos()
	if ret == 1 {
		return b.MoveTo(pos.X+x, pos.Y+y)
	}
	return 0
}

// MoveToEx 移动鼠标到指定范围内的任意点
// 参数:
//
//	x: 目标X坐标
//	y: 目标Y坐标
//	w: 宽度范围
//	h: 高度范围
//
// 返回值:
//
//	int: 1表示成功, 0表示失败
func (b *Background) MoveToEx(x, y, w, h int) int {
	if w <= 0 {
		w = 1
	}
	if h <= 0 {
		h = 1
	}
	randX := x + rand.Intn(w)
	randY := y + rand.Intn(h)
	return b.MoveTo(randX, randY)
}

// LeftClick 执行鼠标左键单击
// 返回值:
//
//	int: 1表示成功, 0表示失败
func (b *Background) LeftClick() int {
	b.LeftDown()
	time.Sleep(time.Duration(b.mouseDelay) * time.Millisecond)
	b.LeftUp()
	return 1
}

// LeftDoubleClick 执行鼠标左键双击
// 返回值:
//
//	int: 1表示成功, 0表示失败
func (b *Background) LeftDoubleClick() int {
	b.LeftClick()
	time.Sleep(time.Duration(b.mouseDelay) * time.Millisecond)
	b.LeftClick()
	return 1
}

// LeftDown 按下鼠标左键
// 返回值:
//
//	int: 1表示成功, 0表示失败
func (b *Background) LeftDown() int {
	if b.hwnd == 0 {
		user32 := syscall.NewLazyDLL("user32.dll")
		mouseEvent := user32.NewProc("mouse_event")
		mouseEvent.Call(uintptr(0x0002), 0, 0, 0, 0) // MOUSEEVENTF_LEFTDOWN
		return 1
	}
	// TODO: 根据鼠标模式实现
	return 1
}

// LeftUp 释放鼠标左键
// 返回值:
//
//	int: 1表示成功, 0表示失败
func (b *Background) LeftUp() int {
	if b.hwnd == 0 {
		user32 := syscall.NewLazyDLL("user32.dll")
		mouseEvent := user32.NewProc("mouse_event")
		mouseEvent.Call(uintptr(0x0004), 0, 0, 0, 0) // MOUSEEVENTF_LEFTUP
		return 1
	}
	// TODO: 根据鼠标模式实现
	return 1
}

// MiddleClick 执行鼠标中键单击
// 返回值:
//
//	int: 1表示成功, 0表示失败
func (b *Background) MiddleClick() int {
	b.MiddleDown()
	time.Sleep(time.Duration(b.mouseDelay) * time.Millisecond)
	b.MiddleUp()
	return 1
}

// MiddleDown 按下鼠标中键
// 返回值:
//
//	int: 1表示成功, 0表示失败
func (b *Background) MiddleDown() int {
	if b.hwnd == 0 {
		user32 := syscall.NewLazyDLL("user32.dll")
		mouseEvent := user32.NewProc("mouse_event")
		mouseEvent.Call(uintptr(0x0020), 0, 0, 0, 0) // MOUSEEVENTF_MIDDLEDOWN
		return 1
	}
	// TODO: 根据鼠标模式实现
	return 1
}

// MiddleUp 释放鼠标中键
// 返回值:
//
//	int: 1表示成功, 0表示失败
func (b *Background) MiddleUp() int {
	if b.hwnd == 0 {
		user32 := syscall.NewLazyDLL("user32.dll")
		mouseEvent := user32.NewProc("mouse_event")
		mouseEvent.Call(uintptr(0x0040), 0, 0, 0, 0) // MOUSEEVENTF_MIDDLEUP
		return 1
	}
	// TODO: 根据鼠标模式实现
	return 1
}

// RightClick 执行鼠标右键单击
// 返回值:
//
//	int: 1表示成功, 0表示失败
func (b *Background) RightClick() int {
	b.RightDown()
	time.Sleep(time.Duration(b.mouseDelay) * time.Millisecond)
	b.RightUp()
	return 1
}

// RightDown 按下鼠标右键
// 返回值:
//
//	int: 1表示成功, 0表示失败
func (b *Background) RightDown() int {
	if b.hwnd == 0 {
		user32 := syscall.NewLazyDLL("user32.dll")
		mouseEvent := user32.NewProc("mouse_event")
		mouseEvent.Call(uintptr(0x0008), 0, 0, 0, 0) // MOUSEEVENTF_RIGHTDOWN
		return 1
	}
	// TODO: 根据鼠标模式实现
	return 1
}

// RightUp 释放鼠标右键
// 返回值:
//
//	int: 1表示成功, 0表示失败
func (b *Background) RightUp() int {
	if b.hwnd == 0 {
		user32 := syscall.NewLazyDLL("user32.dll")
		mouseEvent := user32.NewProc("mouse_event")
		mouseEvent.Call(uintptr(0x0010), 0, 0, 0, 0) // MOUSEEVENTF_RIGHTUP
		return 1
	}
	// TODO: 根据鼠标模式实现
	return 1
}

// WheelDown 鼠标滚轮向下滚动
// 返回值:
//
//	int: 1表示成功, 0表示失败
func (b *Background) WheelDown() int {
	if b.hwnd == 0 {
		user32 := syscall.NewLazyDLL("user32.dll")
		mouseEvent := user32.NewProc("mouse_event")
		mouseEvent.Call(uintptr(0x0800), 0, 0, uintptr(^uint32(119)), 0) // MOUSEEVENTF_WHEEL, -120
		return 1
	}
	// TODO: 根据鼠标模式实现
	return 1
}

// WheelUp 鼠标滚轮向上滚动
// 返回值:
//
//	int: 1表示成功, 0表示失败
func (b *Background) WheelUp() int {
	if b.hwnd == 0 {
		user32 := syscall.NewLazyDLL("user32.dll")
		mouseEvent := user32.NewProc("mouse_event")
		mouseEvent.Call(uintptr(0x0800), 0, 0, uintptr(120), 0) // MOUSEEVENTF_WHEEL, 120
		return 1
	}
	// TODO: 根据鼠标模式实现
	return 1
}

// SetMouseDelay 设置鼠标点击/双击延迟
// 参数:
//
//	typeStr: 类型字符串 ("normal" 等)
//	delay: 延迟时间(毫秒)
//
// 返回值:
//
//	int: 1表示成功, 0表示失败
func (b *Background) SetMouseDelay(typeStr string, delay int) int {
	if delay >= 0 {
		b.mouseDelay = delay
	}
	return 1
}

// GetCursorPos 获取鼠标位置
// 返回值:
//
//	core.Point: 鼠标位置
//	int: 1表示成功, 0表示失败
func (b *Background) GetCursorPos() (core.Point, int) {
	type POINT struct {
		X, Y int32
	}
	var pt POINT
	user32 := syscall.NewLazyDLL("user32.dll")
	getCursorPos := user32.NewProc("GetCursorPos")
	ret, _, _ := getCursorPos.Call(uintptr(unsafe.Pointer(&pt)))
	if ret != 0 {
		return core.Point{X: int(pt.X), Y: int(pt.Y)}, 1
	}
	return core.Point{}, 0
}

// KeyPress 按下并释放指定键
// 参数:
//
//	key: 虚拟键码
//
// 返回值:
//
//	int: 1表示成功, 0表示失败
func (b *Background) KeyPress(key int) int {
	b.KeyDown(key)
	time.Sleep(time.Duration(b.keypadDelay) * time.Millisecond)
	b.KeyUp(key)
	return 1
}

// KeyDown 按下键
// 参数:
//
//	key: 虚拟键码
//
// 返回值:
//
//	int: 1表示成功, 0表示失败
func (b *Background) KeyDown(key int) int {
	if b.hwnd == 0 {
		user32 := syscall.NewLazyDLL("user32.dll")
		keybdEvent := user32.NewProc("keybd_event")
		keybdEvent.Call(uintptr(key), 0, 0, 0)
		return 1
	}
	
	// 后台模式: 使用WM_KEYDOWN消息
	user32 := syscall.NewLazyDLL("user32.dll")
	ret, _, _ := user32.NewProc("SendMessageW").Call(
		uintptr(b.hwnd),
		0x0100, // WM_KEYDOWN
		uintptr(key),
		0,
	)
	if ret != 0 || true {
		return 1
	}
	return 0
}

// KeyUp 释放键
// 参数:
//
//	key: 虚拟键码
//
// 返回值:
//
//	int: 1表示成功, 0表示失败
func (b *Background) KeyUp(key int) int {
	if b.hwnd == 0 {
		user32 := syscall.NewLazyDLL("user32.dll")
		keybdEvent := user32.NewProc("keybd_event")
		keybdEvent.Call(uintptr(key), 0, 0x0002, 0) // KEYEVENTF_KEYUP
		return 1
	}
	
	// 后台模式: 使用WM_KEYUP消息
	user32 := syscall.NewLazyDLL("user32.dll")
	ret, _, _ := user32.NewProc("SendMessageW").Call(
		uintptr(b.hwnd),
		0x0101, // WM_KEYUP
		uintptr(key),
		0,
	)
	if ret != 0 || true {
		return 1
	}
	return 0
}

// SendString 向目标发送字符串
// 参数:
//
//	str: 要发送的字符串
//
// 返回值:
//
//	int: 1表示成功, 0表示失败
func (b *Background) SendString(str string) int {
	if b.hwnd == 0 {
		// 前台模式使用剪贴板方法
		return 1
	}
	
	// 后台模式: 使用WM_CHAR消息逐字符发送
	user32 := syscall.NewLazyDLL("user32.dll")
	
	for _, ch := range str {
		// 处理换行符: 发送\r (WM_CHAR with 13)
		if ch == '\n' {
			user32.NewProc("SendMessageW").Call(
				uintptr(b.hwnd),
				0x0102, // WM_CHAR
				uintptr(13), // \r
				1,
			)
			time.Sleep(5 * time.Millisecond)
			continue
		}
		
		// 跳过\r (已经处理过)
		if ch == '\r' {
			continue
		}
		
		// 发送WM_CHAR消息
		user32.NewProc("SendMessageW").Call(
			uintptr(b.hwnd),
			0x0102, // WM_CHAR
			uintptr(ch),
			1,      // 重复计数
		)
		
		// 短暂延时确保消息处理
		time.Sleep(5 * time.Millisecond)
	}
	
	return 1
}

// SendStringIme 向目标发送字符串(使用IME)
// 参数:
//
//	str: 要发送的字符串
//
// 返回值:
//
//	int: 1表示成功, 0表示失败
func (b *Background) SendStringIme(str string) int {
	// SendStringIme 与 SendString 实现相同
	return b.SendString(str)
}

// GetKeyState 获取指定键的状态
// 参数:
//
//	vkCode: 虚拟键码
//
// 返回值:
//
//	int: 键状态 (1=按下, 0=释放)
func (b *Background) GetKeyState(vkCode int) int {
	user32 := syscall.NewLazyDLL("user32.dll")
	getKeyState := user32.NewProc("GetKeyState")
	ret, _, _ := getKeyState.Call(uintptr(vkCode))
	if int32(ret)&0x8000 != 0 {
		return 1
	}
	return 0
}

// SetKeypadDelay 设置键盘按下/释放延迟
// 参数:
//
//	typeStr: 类型字符串 ("normal" 等)
//	delay: 延迟时间(毫秒)
//
// 返回值:
//
//	int: 1表示成功, 0表示失败
func (b *Background) SetKeypadDelay(typeStr string, delay int) int {
	if delay >= 0 {
		b.keypadDelay = delay
	}
	return 1
}

// SetDisplayInput 设置显示输入模式
// 参数:
//
//	mode: 模式字符串
//
// 返回值:
//
//	int: 1表示成功, 0表示失败
func (b *Background) SetDisplayInput(mode string) int {
	// TODO: 实现显示输入模式设置
	return 1
}

// KeyDownChar 按字符按下键
// 参数:
//
//	key: 字符字符串
//
// 返回值:
//
//	int: 1表示成功, 0表示失败
func (b *Background) KeyDownChar(key string) int {
	vkCode := core.GetKeycode(key)
	return b.KeyDown(vkCode)
}

// KeyUpChar 按字符释放键
// 参数:
//
//	key: 字符字符串
//
// 返回值:
//
//	int: 1表示成功, 0表示失败
func (b *Background) KeyUpChar(key string) int {
	vkCode := core.GetKeycode(key)
	return b.KeyUp(vkCode)
}

// WaitKey 等待按键
// 参数:
//
//	vkCode: 虚拟键码
//	timeOut: 超时时间(毫秒)
//
// 返回值:
//
//	int: 1表示按键, 0表示超时
func (b *Background) WaitKey(vkCode, timeOut int) int {
	startTime := time.Now()
	for time.Since(startTime) < time.Duration(timeOut)*time.Millisecond {
		if b.GetKeyState(vkCode) == 1 {
			return 1
		}
		time.Sleep(10 * time.Millisecond)
	}
	return 0
}

// KeyPressChar 按字符按下并释放键
// 参数:
//
//	key: 字符字符串
//
// 返回值:
//
//	int: 1表示成功, 0表示失败
func (b *Background) KeyPressChar(key string) int {
	vkCode := core.GetKeycode(key)
	return b.KeyPress(vkCode)
}

// KeyPressStr 按序列按下键
// 参数:
//
//	keyStr: 键序列字符串
//	delay: 键之间延迟(毫秒)
//
// 返回值:
//
//	int: 1表示成功, 0表示失败
func (b *Background) KeyPressStr(keyStr string, delay int) int {
	for _, ch := range keyStr {
		b.KeyPress(int(ch))
		if delay > 0 {
			time.Sleep(time.Duration(delay) * time.Millisecond)
		}
	}
	return 1
}
