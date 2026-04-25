// Package main 提供DLL导出接口
// 该文件用于将Go代码编译为DLL,供其他语言调用
package main

/*
#include <stdlib.h>
#include <stdio.h>
*/
import "C"
import (
	"sync"
	"unsafe"

	"github.com/yuan71058/GOP/libop"
)

var (
	opInstance *libop.LibOP
	opMutex    sync.Mutex
)

//export CreateOp
func CreateOp() C.int {
	opMutex.Lock()
	defer opMutex.Unlock()
	opInstance = libop.NewLibOP()
	return 1
}

//export Ver
func Ver() *C.char {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	return C.CString(opInstance.Ver())
}

//export SetPath
func SetPath(path *C.char) C.int {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	goPath := C.GoString(path)
	return C.int(opInstance.SetPath(goPath))
}

//export GetPath
func GetPath() *C.char {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	return C.CString(opInstance.GetPath())
}

//export GetBasePath
func GetBasePath() *C.char {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	return C.CString(opInstance.GetBasePath())
}

//export GetID
func GetID() C.int {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	return C.int(opInstance.GetID())
}

//export OpGetLastError
func OpGetLastError() C.int {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	return C.int(opInstance.GetLastError())
}

//export SetShowErrorMsg
func SetShowErrorMsg(showType C.int) C.int {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	return C.int(opInstance.SetShowErrorMsg(int(showType)))
}

//export OpSleep
func OpSleep(ms C.int) C.int {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	return C.int(opInstance.Sleep(int(ms)))
}

//export Delay
func Delay(ms C.int) C.int {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	return C.int(opInstance.Delay(int(ms)))
}

//export Delays
func Delays(msMin, msMax C.int) C.int {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	return C.int(opInstance.Delays(int(msMin), int(msMax)))
}

//export InjectDll
func InjectDll(processName, dllName *C.char) C.int {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	goProcessName := C.GoString(processName)
	goDllName := C.GoString(dllName)
	return C.int(opInstance.InjectDll(goProcessName, goDllName))
}

//export EnablePicCache
func EnablePicCache(enable C.int) C.int {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	return C.int(opInstance.EnablePicCache(int(enable)))
}

//export CapturePre
func CapturePre(fileName *C.char) C.int {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	goFileName := C.GoString(fileName)
	return C.int(opInstance.CapturePre(goFileName))
}

//export SetScreenDataMode
func SetScreenDataMode(mode C.int) C.int {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	return C.int(opInstance.SetScreenDataMode(int(mode)))
}

//export AStarFindPath
func AStarFindPath(mapWidth, mapHeight, beginX, beginY, endX, endY C.int, disablePoints *C.char) *C.char {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	goDisablePoints := C.GoString(disablePoints)
	result := opInstance.AStarFindPath(int(mapWidth), int(mapHeight), goDisablePoints, int(beginX), int(beginY), int(endX), int(endY))
	return C.CString(result)
}

//export FindNearestPos
func FindNearestPos(allPos *C.char, posType, x, y C.int) *C.char {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	goAllPos := C.GoString(allPos)
	result := opInstance.FindNearestPos(goAllPos, int(posType), int(x), int(y))
	return C.CString(result)
}

//export EnumWindow
func EnumWindow(parent C.int, title, className *C.char, filter C.int) *C.char {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	goTitle := C.GoString(title)
	goClassName := C.GoString(className)
	result := opInstance.EnumWindow(int(parent), goTitle, goClassName, int(filter))
	return C.CString(result)
}

//export EnumWindowByProcess
func EnumWindowByProcess(processName, title, className *C.char, filter C.int) *C.char {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	goProcessName := C.GoString(processName)
	goTitle := C.GoString(title)
	goClassName := C.GoString(className)
	result := opInstance.EnumWindowByProcess(goProcessName, goTitle, goClassName, int(filter))
	return C.CString(result)
}

//export EnumProcess
func EnumProcess(name *C.char) *C.char {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	goName := C.GoString(name)
	result := opInstance.EnumProcess(goName)
	return C.CString(result)
}

//export ClientToScreen
func ClientToScreen(hwnd C.int, x, y *C.int) C.int {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	goX := int(*x)
	goY := int(*y)
	ret := opInstance.ClientToScreen(int(hwnd), &goX, &goY)
	*x = C.int(goX)
	*y = C.int(goY)
	return C.int(ret)
}

//export FindWindow
func FindWindow(className, title *C.char) C.int {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	goClassName := C.GoString(className)
	goTitle := C.GoString(title)
	return C.int(opInstance.FindWindow(goClassName, goTitle))
}

//export FindWindowByProcess
func FindWindowByProcess(processName, className, title *C.char) C.int {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	goProcessName := C.GoString(processName)
	goClassName := C.GoString(className)
	goTitle := C.GoString(title)
	return C.int(opInstance.FindWindowByProcess(goProcessName, goClassName, goTitle))
}

//export FindWindowByProcessId
func FindWindowByProcessId(processId C.int, className, title *C.char) C.int {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	goClassName := C.GoString(className)
	goTitle := C.GoString(title)
	return C.int(opInstance.FindWindowByProcessId(int(processId), goClassName, goTitle))
}

//export FindWindowEx
func FindWindowEx(parent C.int, className, title *C.char) C.int {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	goClassName := C.GoString(className)
	goTitle := C.GoString(title)
	return C.int(opInstance.FindWindowEx(int(parent), goClassName, goTitle))
}

//export GetClientRect
func GetClientRect(hwnd C.int, x1, y1, x2, y2 *C.int) C.int {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	goX1, goY1, goX2, goY2 := 0, 0, 0, 0
	ret := opInstance.GetClientRect(int(hwnd), &goX1, &goY1, &goX2, &goY2)
	*x1 = C.int(goX1)
	*y1 = C.int(goY1)
	*x2 = C.int(goX2)
	*y2 = C.int(goY2)
	return C.int(ret)
}

//export GetClientSize
func GetClientSize(hwnd C.int, width, height *C.int) C.int {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	goWidth, goHeight := 0, 0
	ret := opInstance.GetClientSize(int(hwnd), &goWidth, &goHeight)
	*width = C.int(goWidth)
	*height = C.int(goHeight)
	return C.int(ret)
}

//export GetForegroundFocus
func GetForegroundFocus() C.int {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	return C.int(opInstance.GetForegroundFocus())
}

//export GetForegroundWindow
func GetForegroundWindow() C.int {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	return C.int(opInstance.GetForegroundWindow())
}

//export GetMousePointWindow
func GetMousePointWindow() C.int {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	return C.int(opInstance.GetMousePointWindow())
}

//export GetPointWindow
func GetPointWindow(x, y C.int) C.int {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	return C.int(opInstance.GetPointWindow(int(x), int(y)))
}

//export GetProcessInfo
func GetProcessInfo(pid C.int) *C.char {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	result := opInstance.GetProcessInfo(int(pid))
	return C.CString(result)
}

//export GetSpecialWindow
func GetSpecialWindow(flag C.int) C.int {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	return C.int(opInstance.GetSpecialWindow(int(flag)))
}

//export GetWindow
func GetWindow(hwnd C.int, flag C.int) C.int {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	return C.int(opInstance.GetWindow(int(hwnd), int(flag)))
}

//export GetWindowClass
func GetWindowClass(hwnd C.int) *C.char {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	result := opInstance.GetWindowClass(int(hwnd))
	return C.CString(result)
}

//export GetWindowProcessId
func GetWindowProcessId(hwnd C.int) C.int {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	return C.int(opInstance.GetWindowProcessId(int(hwnd)))
}

//export GetWindowProcessPath
func GetWindowProcessPath(hwnd C.int) *C.char {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	result := opInstance.GetWindowProcessPath(int(hwnd))
	return C.CString(result)
}

//export GetWindowRect
func GetWindowRect(hwnd C.int, x1, y1, x2, y2 *C.int) C.int {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	goX1, goY1, goX2, goY2 := 0, 0, 0, 0
	ret := opInstance.GetWindowRect(int(hwnd), &goX1, &goY1, &goX2, &goY2)
	*x1 = C.int(goX1)
	*y1 = C.int(goY1)
	*x2 = C.int(goX2)
	*y2 = C.int(goY2)
	return C.int(ret)
}

//export GetWindowState
func GetWindowState(hwnd, flag C.int) C.int {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	return C.int(opInstance.GetWindowState(int(hwnd), int(flag)))
}

//export GetWindowTitle
func GetWindowTitle(hwnd C.int) *C.char {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	result := opInstance.GetWindowTitle(int(hwnd))
	return C.CString(result)
}

//export MoveWindow
func MoveWindow(hwnd, x, y C.int) C.int {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	return C.int(opInstance.MoveWindow(int(hwnd), int(x), int(y)))
}

//export ScreenToClient
func ScreenToClient(hwnd C.int, x, y *C.int) C.int {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	goX := int(*x)
	goY := int(*y)
	ret := opInstance.ScreenToClient(int(hwnd), &goX, &goY)
	*x = C.int(goX)
	*y = C.int(goY)
	return C.int(ret)
}

//export SendPaste
func SendPaste(hwnd C.int) C.int {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	return C.int(opInstance.SendPaste(int(hwnd)))
}

//export SetClientSize
func SetClientSize(hwnd, width, height C.int) C.int {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	return C.int(opInstance.SetClientSize(int(hwnd), int(width), int(height)))
}

//export SetWindowState
func SetWindowState(hwnd, flag C.int) C.int {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	return C.int(opInstance.SetWindowState(int(hwnd), int(flag)))
}

//export SetWindowSize
func SetWindowSize(hwnd, width, height C.int) C.int {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	return C.int(opInstance.SetWindowSize(int(hwnd), int(width), int(height)))
}

//export SetWindowText
func SetWindowText(hwnd C.int, title *C.char) C.int {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	goTitle := C.GoString(title)
	return C.int(opInstance.SetWindowText(int(hwnd), goTitle))
}

//export SetWindowTransparent
func SetWindowTransparent(hwnd, trans C.int) C.int {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	return C.int(opInstance.SetWindowTransparent(int(hwnd), int(trans)))
}

//export SendString
func SendString(hwnd C.int, str *C.char) C.int {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	goStr := C.GoString(str)
	return C.int(opInstance.SendString(int(hwnd), goStr))
}

//export SendStringIme
func SendStringIme(hwnd C.int, str *C.char) C.int {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	goStr := C.GoString(str)
	return C.int(opInstance.SendStringIme(int(hwnd), goStr))
}

//export RunApp
func RunApp(cmdline *C.char, mode C.int) C.int {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	goCmdline := C.GoString(cmdline)
	return C.int(opInstance.RunApp(goCmdline, int(mode)))
}

//export WinExec
func WinExec(cmdline *C.char, cmdShow C.int) C.int {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	goCmdline := C.GoString(cmdline)
	return C.int(opInstance.WinExec(goCmdline, int(cmdShow)))
}

//export GetCmdStr
func GetCmdStr(cmd *C.char, milliseconds C.int) *C.char {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	goCmd := C.GoString(cmd)
	result := opInstance.GetCmdStr(goCmd, int(milliseconds))
	return C.CString(result)
}

//export SetClipboard
func SetClipboard(str *C.char) C.int {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	goStr := C.GoString(str)
	return C.int(opInstance.SetClipboard(goStr))
}

//export GetClipboard
func GetClipboard() *C.char {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	result := opInstance.GetClipboard()
	return C.CString(result)
}

//export BindWindow
func BindWindow(hwnd C.int, display, mouse, keypad *C.char, mode C.int) C.int {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	goDisplay := C.GoString(display)
	goMouse := C.GoString(mouse)
	goKeypad := C.GoString(keypad)
	return C.int(opInstance.BindWindow(int(hwnd), goDisplay, goMouse, goKeypad, int(mode)))
}

//export UnBindWindow
func UnBindWindow() C.int {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	return C.int(opInstance.UnBindWindow())
}

//export GetBindWindow
func GetBindWindow() C.int {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	return C.int(opInstance.GetBindWindow())
}

//export IsBind
func IsBind() C.int {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	return C.int(opInstance.IsBind())
}

//export GetCursorPos
func GetCursorPos(x, y *C.int) C.int {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	goX, goY := 0, 0
	ret := opInstance.GetCursorPos(&goX, &goY)
	*x = C.int(goX)
	*y = C.int(goY)
	return C.int(ret)
}

//export MoveR
func MoveR(x, y C.int) C.int {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	return C.int(opInstance.MoveR(int(x), int(y)))
}

//export MoveTo
func MoveTo(x, y C.int) C.int {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	return C.int(opInstance.MoveTo(int(x), int(y)))
}

//export MoveToEx
func MoveToEx(x, y, w, h C.int) C.int {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	return C.int(opInstance.MoveToEx(int(x), int(y), int(w), int(h)))
}

//export LeftClick
func LeftClick() C.int {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	return C.int(opInstance.LeftClick())
}

//export LeftDoubleClick
func LeftDoubleClick() C.int {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	return C.int(opInstance.LeftDoubleClick())
}

//export LeftDown
func LeftDown() C.int {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	return C.int(opInstance.LeftDown())
}

//export LeftUp
func LeftUp() C.int {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	return C.int(opInstance.LeftUp())
}

//export MiddleClick
func MiddleClick() C.int {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	return C.int(opInstance.MiddleClick())
}

//export MiddleDown
func MiddleDown() C.int {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	return C.int(opInstance.MiddleDown())
}

//export MiddleUp
func MiddleUp() C.int {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	return C.int(opInstance.MiddleUp())
}

//export RightClick
func RightClick() C.int {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	return C.int(opInstance.RightClick())
}

//export RightDown
func RightDown() C.int {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	return C.int(opInstance.RightDown())
}

//export RightUp
func RightUp() C.int {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	return C.int(opInstance.RightUp())
}

//export WheelDown
func WheelDown() C.int {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	return C.int(opInstance.WheelDown())
}

//export WheelUp
func WheelUp() C.int {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	return C.int(opInstance.WheelUp())
}

//export SetMouseDelay
func SetMouseDelay(typeStr *C.char, delay C.int) C.int {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	goTypeStr := C.GoString(typeStr)
	return C.int(opInstance.SetMouseDelay(goTypeStr, int(delay)))
}

//export GetKeyState
func GetKeyState(vkCode C.int) C.int {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	return C.int(opInstance.GetKeyState(int(vkCode)))
}

//export KeyDown
func KeyDown(vkCode C.int) C.int {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	return C.int(opInstance.KeyDown(int(vkCode)))
}

//export KeyDownChar
func KeyDownChar(vkCode *C.char) C.int {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	goVkCode := C.GoString(vkCode)
	return C.int(opInstance.KeyDownChar(goVkCode))
}

//export KeyUp
func KeyUp(vkCode C.int) C.int {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	return C.int(opInstance.KeyUp(int(vkCode)))
}

//export KeyUpChar
func KeyUpChar(vkCode *C.char) C.int {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	goVkCode := C.GoString(vkCode)
	return C.int(opInstance.KeyUpChar(goVkCode))
}

//export WaitKey
func WaitKey(vkCode, timeOut C.int) C.int {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	return C.int(opInstance.WaitKey(int(vkCode), int(timeOut)))
}

//export KeyPress
func KeyPress(vkCode C.int) C.int {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	return C.int(opInstance.KeyPress(int(vkCode)))
}

//export KeyPressChar
func KeyPressChar(vkCode *C.char) C.int {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	goVkCode := C.GoString(vkCode)
	return C.int(opInstance.KeyPressChar(goVkCode))
}

//export SetKeypadDelay
func SetKeypadDelay(typeStr *C.char, delay C.int) C.int {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	goTypeStr := C.GoString(typeStr)
	return C.int(opInstance.SetKeypadDelay(goTypeStr, int(delay)))
}

//export KeyPressStr
func KeyPressStr(keyStr *C.char, delay C.int) C.int {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	goKeyStr := C.GoString(keyStr)
	return C.int(opInstance.KeyPressStr(goKeyStr, int(delay)))
}

//export Capture
func Capture(x1, y1, x2, y2 C.int, fileName *C.char) C.int {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	goFileName := C.GoString(fileName)
	return C.int(opInstance.Capture(int(x1), int(y1), int(x2), int(y2), goFileName))
}

//export CmpColor
func CmpColor(x, y C.int, color *C.char, sim C.double) C.int {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	goColor := C.GoString(color)
	return C.int(opInstance.CmpColor(int(x), int(y), goColor, float64(sim)))
}

//export FindColor
func FindColor(x1, y1, x2, y2 C.int, color *C.char, sim C.double, dir C.int, x, y *C.int) C.int {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	goColor := C.GoString(color)
	goX, goY := 0, 0
	ret := opInstance.FindColor(int(x1), int(y1), int(x2), int(y2), goColor, float64(sim), int(dir), &goX, &goY)
	*x = C.int(goX)
	*y = C.int(goY)
	return C.int(ret)
}

//export FindColorEx
func FindColorEx(x1, y1, x2, y2 C.int, color *C.char, sim C.double, dir C.int) *C.char {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	goColor := C.GoString(color)
	result := opInstance.FindColorEx(int(x1), int(y1), int(x2), int(y2), goColor, float64(sim), int(dir))
	return C.CString(result)
}

//export GetColorNum
func GetColorNum(x1, y1, x2, y2 C.int, color *C.char, sim C.double) C.int {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	goColor := C.GoString(color)
	return C.int(opInstance.GetColorNum(int(x1), int(y1), int(x2), int(y2), goColor, float64(sim)))
}

//export FindMultiColor
func FindMultiColor(x1, y1, x2, y2 C.int, firstColor, offsetColor *C.char, sim C.double, dir C.int, x, y *C.int) C.int {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	goFirstColor := C.GoString(firstColor)
	goOffsetColor := C.GoString(offsetColor)
	goX, goY := 0, 0
	ret := opInstance.FindMultiColor(int(x1), int(y1), int(x2), int(y2), goFirstColor, goOffsetColor, float64(sim), int(dir), &goX, &goY)
	*x = C.int(goX)
	*y = C.int(goY)
	return C.int(ret)
}

//export FindMultiColorEx
func FindMultiColorEx(x1, y1, x2, y2 C.int, firstColor, offsetColor *C.char, sim C.double, dir C.int) *C.char {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	goFirstColor := C.GoString(firstColor)
	goOffsetColor := C.GoString(offsetColor)
	result := opInstance.FindMultiColorEx(int(x1), int(y1), int(x2), int(y2), goFirstColor, goOffsetColor, float64(sim), int(dir))
	return C.CString(result)
}

//export FindPic
func FindPic(x1, y1, x2, y2 C.int, files, deltaColor *C.char, sim C.double, dir C.int, x, y *C.int) C.int {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	goFiles := C.GoString(files)
	goDeltaColor := C.GoString(deltaColor)
	goX, goY := 0, 0
	ret := opInstance.FindPic(int(x1), int(y1), int(x2), int(y2), goFiles, goDeltaColor, float64(sim), int(dir), &goX, &goY)
	*x = C.int(goX)
	*y = C.int(goY)
	return C.int(ret)
}

//export FindPicEx
func FindPicEx(x1, y1, x2, y2 C.int, files, deltaColor *C.char, sim C.double, dir C.int) *C.char {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	goFiles := C.GoString(files)
	goDeltaColor := C.GoString(deltaColor)
	result := opInstance.FindPicEx(int(x1), int(y1), int(x2), int(y2), goFiles, goDeltaColor, float64(sim), int(dir))
	return C.CString(result)
}

//export FindPicExS
func FindPicExS(x1, y1, x2, y2 C.int, files, deltaColor *C.char, sim C.double, dir C.int) *C.char {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	goFiles := C.GoString(files)
	goDeltaColor := C.GoString(deltaColor)
	result := opInstance.FindPicExS(int(x1), int(y1), int(x2), int(y2), goFiles, goDeltaColor, float64(sim), int(dir))
	return C.CString(result)
}

//export FindColorBlock
func FindColorBlock(x1, y1, x2, y2 C.int, color *C.char, sim C.double, count, height, width C.int, x, y *C.int) C.int {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	goColor := C.GoString(color)
	goX, goY := 0, 0
	ret := opInstance.FindColorBlock(int(x1), int(y1), int(x2), int(y2), goColor, float64(sim), int(count), int(height), int(width), &goX, &goY)
	*x = C.int(goX)
	*y = C.int(goY)
	return C.int(ret)
}

//export FindColorBlockEx
func FindColorBlockEx(x1, y1, x2, y2 C.int, color *C.char, sim C.double, count, height, width C.int) *C.char {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	goColor := C.GoString(color)
	result := opInstance.FindColorBlockEx(int(x1), int(y1), int(x2), int(y2), goColor, float64(sim), int(count), int(height), int(width))
	return C.CString(result)
}

//export GetColor
func GetColor(x, y C.int) *C.char {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	result := opInstance.GetColor(int(x), int(y))
	return C.CString(result)
}

//export SetDisplayInput
func SetDisplayInput(mode *C.char) C.int {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	goMode := C.GoString(mode)
	return C.int(opInstance.SetDisplayInput(goMode))
}

//export LoadPic
func LoadPic(fileName *C.char) C.int {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	goFileName := C.GoString(fileName)
	return C.int(opInstance.LoadPic(goFileName))
}

//export FreePic
func FreePic(fileName *C.char) C.int {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	goFileName := C.GoString(fileName)
	return C.int(opInstance.FreePic(goFileName))
}

//export LoadMemPic
func LoadMemPic(fileName *C.char, data unsafe.Pointer, size C.int) C.int {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	goFileName := C.GoString(fileName)
	goData := C.GoBytes(data, size)
	return C.int(opInstance.LoadMemPic(goFileName, goData, int(size)))
}

//export GetPicSize
func GetPicSize(picName *C.char, width, height *C.int) C.int {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	goPicName := C.GoString(picName)
	goWidth, goHeight := 0, 0
	ret := opInstance.GetPicSize(goPicName, &goWidth, &goHeight)
	*width = C.int(goWidth)
	*height = C.int(goHeight)
	return C.int(ret)
}

//export GetScreenData
func GetScreenData(x1, y1, x2, y2 C.int) unsafe.Pointer {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	data, _ := opInstance.GetScreenData(int(x1), int(y1), int(x2), int(y2))
	return unsafe.Pointer(data)
}

//export GetScreenDataBmp
func GetScreenDataBmp(x1, y1, x2, y2 C.int, width, height *C.int) unsafe.Pointer {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	data, size, _ := opInstance.GetScreenDataBmp(int(x1), int(y1), int(x2), int(y2))
	*width = C.int(size)
	return unsafe.Pointer(data)
}

//export GetScreenFrameInfo
func GetScreenFrameInfo(frameID, frameTime *C.int) C.int {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	goFrameID, goFrameTime := opInstance.GetScreenFrameInfo()
	*frameID = C.int(goFrameID)
	*frameTime = C.int(goFrameTime)
	return 1
}

//export MatchPicName
func MatchPicName(picName *C.char) *C.char {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	goPicName := C.GoString(picName)
	result := opInstance.MatchPicName(goPicName)
	return C.CString(result)
}

//export SetOcrEngine
func SetOcrEngine(pathOfEngine, dllName, argv *C.char) C.int {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	goPath := C.GoString(pathOfEngine)
	goDllName := C.GoString(dllName)
	goArgv := C.GoString(argv)
	return C.int(opInstance.SetOcrEngine(goPath, goDllName, goArgv))
}

//export SetDict
func SetDict(idx C.int, fileName *C.char) C.int {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	goFileName := C.GoString(fileName)
	return C.int(opInstance.SetDict(int(idx), goFileName))
}

//export GetDict
func GetDict(idx, fontIndex C.int) *C.char {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	result := opInstance.GetDict(int(idx), int(fontIndex))
	return C.CString(result)
}

//export SetMemDict
func SetMemDict(idx C.int, data unsafe.Pointer, size C.int) C.int {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	goData := C.GoBytes(data, size)
	return C.int(opInstance.SetMemDict(int(idx), goData, int(size)))
}

//export UseDict
func UseDict(idx C.int) C.int {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	return C.int(opInstance.UseDict(int(idx)))
}

//export AddDict
func AddDict(idx C.int, dictInfo *C.char) C.int {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	goDictInfo := C.GoString(dictInfo)
	return C.int(opInstance.AddDict(int(idx), goDictInfo))
}

//export SaveDict
func SaveDict(idx C.int, fileName *C.char) C.int {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	goFileName := C.GoString(fileName)
	return C.int(opInstance.SaveDict(int(idx), goFileName))
}

//export ClearDict
func ClearDict(idx C.int) C.int {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	return C.int(opInstance.ClearDict(int(idx)))
}

//export GetDictCount
func GetDictCount(idx C.int) C.int {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	return C.int(opInstance.GetDictCount(int(idx)))
}

//export GetNowDict
func GetNowDict() C.int {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	return C.int(opInstance.GetNowDict())
}

//export FetchWord
func FetchWord(x1, y1, x2, y2 C.int, color, word *C.char) *C.char {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	goColor := C.GoString(color)
	goWord := C.GoString(word)
	result := opInstance.FetchWord(int(x1), int(y1), int(x2), int(y2), goColor, goWord)
	return C.CString(result)
}

//export GetWordsNoDict
func GetWordsNoDict(x1, y1, x2, y2 C.int, color *C.char) *C.char {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	goColor := C.GoString(color)
	result := opInstance.GetWordsNoDict(int(x1), int(y1), int(x2), int(y2), goColor)
	return C.CString(result)
}

//export GetWordResultCount
func GetWordResultCount(result *C.char) C.int {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	goResult := C.GoString(result)
	return C.int(opInstance.GetWordResultCount(goResult))
}

//export GetWordResultPos
func GetWordResultPos(result *C.char, index C.int, x, y *C.int) C.int {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	goResult := C.GoString(result)
	goX, goY := 0, 0
	ret := opInstance.GetWordResultPos(goResult, int(index), &goX, &goY)
	*x = C.int(goX)
	*y = C.int(goY)
	return C.int(ret)
}

//export GetWordResultStr
func GetWordResultStr(result *C.char, index C.int) *C.char {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	goResult := C.GoString(result)
	str := opInstance.GetWordResultStr(goResult, int(index))
	return C.CString(str)
}

//export Ocr
func Ocr(x1, y1, x2, y2 C.int, color *C.char, sim C.double) *C.char {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	goColor := C.GoString(color)
	result := opInstance.Ocr(int(x1), int(y1), int(x2), int(y2), goColor, float64(sim))
	return C.CString(result)
}

//export OcrEx
func OcrEx(x1, y1, x2, y2 C.int, color *C.char, sim C.double) *C.char {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	goColor := C.GoString(color)
	result := opInstance.OcrEx(int(x1), int(y1), int(x2), int(y2), goColor, float64(sim))
	return C.CString(result)
}

//export FindStr
func FindStr(x1, y1, x2, y2 C.int, strs, color *C.char, sim C.double, retX, retY *C.int) C.int {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	goStrs := C.GoString(strs)
	goColor := C.GoString(color)
	goX, goY := 0, 0
	ret := opInstance.FindStr(int(x1), int(y1), int(x2), int(y2), goStrs, goColor, float64(sim), &goX, &goY)
	*retX = C.int(goX)
	*retY = C.int(goY)
	return C.int(ret)
}

//export FindStrEx
func FindStrEx(x1, y1, x2, y2 C.int, strs, color *C.char, sim C.double) *C.char {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	goStrs := C.GoString(strs)
	goColor := C.GoString(color)
	result := opInstance.FindStrEx(int(x1), int(y1), int(x2), int(y2), goStrs, goColor, float64(sim))
	return C.CString(result)
}

//export OcrAuto
func OcrAuto(x1, y1, x2, y2 C.int, sim C.double) *C.char {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	result := opInstance.OcrAuto(int(x1), int(y1), int(x2), int(y2), float64(sim))
	return C.CString(result)
}

//export OcrFromFile
func OcrFromFile(fileName, colorFormat *C.char, sim C.double) *C.char {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	goFileName := C.GoString(fileName)
	goColorFormat := C.GoString(colorFormat)
	result := opInstance.OcrFromFile(goFileName, goColorFormat, float64(sim))
	return C.CString(result)
}

//export OcrAutoFromFile
func OcrAutoFromFile(fileName *C.char, sim C.double) *C.char {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	goFileName := C.GoString(fileName)
	result := opInstance.OcrAutoFromFile(goFileName, float64(sim))
	return C.CString(result)
}

//export FindLine
func FindLine(x1, y1, x2, y2 C.int, color *C.char, sim C.double) *C.char {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	goColor := C.GoString(color)
	result := opInstance.FindLine(int(x1), int(y1), int(x2), int(y2), goColor, float64(sim))
	return C.CString(result)
}

//export WriteData
func WriteData(hwnd C.int, address, data *C.char, size C.int) C.int {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	goAddress := C.GoString(address)
	goData := C.GoString(data)
	return C.int(opInstance.WriteData(int(hwnd), goAddress, goData, int(size)))
}

//export ReadData
func ReadData(hwnd C.int, address *C.char, size C.int) *C.char {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	goAddress := C.GoString(address)
	result := opInstance.ReadData(int(hwnd), goAddress, int(size))
	return C.CString(result)
}

//export FreeMemory
func FreeMemory(ptr unsafe.Pointer) {
	C.free(ptr)
}

func main() {
	// This function is empty because this is a DLL library
}
