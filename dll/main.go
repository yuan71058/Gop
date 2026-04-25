// Package main 提供DLL导出接口
// 该文件用于将Go代码编译为DLL，供其他语言调用
package main

/*
#include <stdlib.h>
#include <stdio.h>
*/
import "C"
import (
	"fmt"
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

//export FindWindow
func FindWindow(className, title *C.char) C.int {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	goClassName := C.GoString(className)
	goTitle := C.GoString(title)
	return C.int(opInstance.FindWindow(goClassName, goTitle))
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

//export MoveTo
func MoveTo(x, y C.int) C.int {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	return C.int(opInstance.MoveTo(int(x), int(y)))
}

//export LeftClick
func LeftClick() C.int {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	return C.int(opInstance.LeftClick())
}

//export RightClick
func RightClick() C.int {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	return C.int(opInstance.RightClick())
}

//export KeyPress
func KeyPress(key C.int) C.int {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	return C.int(opInstance.KeyPress(int(key)))
}

//export SendString
func SendString(str *C.char) C.int {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	goStr := C.GoString(str)
	return C.int(opInstance.SendString(goStr))
}

//export FindPic
func FindPic(x1, y1, x2, y2 C.int, picName, deltaColor *C.char, sim C.double, dir C.int) *C.char {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	goPicName := C.GoString(picName)
	goDeltaColor := C.GoString(deltaColor)
	x, y, ret := opInstance.FindPic(int(x1), int(y1), int(x2), int(y2), goPicName, goDeltaColor, float64(sim), int(dir))
	result := C.CString(fmt.Sprintf("%d,%d|%d", x, y, ret))
	return result
}

//export FindColor
func FindColor(x1, y1, x2, y2 C.int, color, sim *C.char, dir C.int) *C.char {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	goColor := C.GoString(color)
	goSim := C.GoString(sim)
	x, y, ret := opInstance.FindColor(int(x1), int(y1), int(x2), int(y2), goColor, goSim, int(dir))
	result := C.CString(fmt.Sprintf("%d,%d|%d", x, y, ret))
	return result
}

//export Ocr
func Ocr(x1, y1, x2, y2 C.int, color, sim *C.char) *C.char {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	goColor := C.GoString(color)
	goSim := C.GoString(sim)
	result := opInstance.Ocr(int(x1), int(y1), int(x2), int(y2), goColor, goSim)
	return C.CString(result)
}

//export GetCursorPos
func GetCursorPos(x, y *C.int) C.int {
	if opInstance == nil {
		opInstance = libop.NewLibOP()
	}
	goX, goY, ret := opInstance.GetCursorPos()
	*x = C.int(goX)
	*y = C.int(goY)
	return C.int(ret)
}

//export FreeMemory
func FreeMemory(ptr unsafe.Pointer) {
	C.free(ptr)
}

func main() {
	// 该函数为空，因为这是一个DLL库
}
