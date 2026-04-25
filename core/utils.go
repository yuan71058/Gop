// Package core 提供OP项目的核心工具函数
// 该包包含字符串处理、颜色转换等辅助功能
package core

import (
	"fmt"
	"strings"
	"syscall"
	"unicode"
	"unsafe"
)

// WString 表示Windows宽字符串（UTF-16）
// 用于与Windows API交互
type WString = *uint16

// StringToUTF16 将Go字符串转换为UTF-16指针
// 参数:
//   s: Go字符串
// 返回值:
//   *uint16: UTF-16编码的字符串指针
func StringToUTF16(s string) *uint16 {
	if s == "" {
		return nil
	}
	p, _ := syscall.UTF16PtrFromString(s)
	return p
}

// UTF16ToString 将UTF-16指针转换为Go字符串
// 参数:
//   p: UTF-16编码的字符串指针
// 返回值:
//   string: Go字符串
func UTF16ToString(p *uint16) string {
	if p == nil {
		return ""
	}
	// 找到字符串长度
	var length int
	for {
		if *(*uint16)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + uintptr(length)*2)) == 0 {
			break
		}
		length++
	}
	// 创建切片
	slice := unsafe.Slice(p, length)
	return syscall.UTF16ToString(slice)
}

// ReplaceW 替换字符串中的子串
// 参数:
//   s: 原始字符串
//   old: 要替换的子串
//   new: 替换后的子串
// 返回值:
//   string: 替换后的字符串
func ReplaceW(s, old, new string) string {
	return strings.ReplaceAll(s, old, new)
}

// ToLower 将字符串转换为小写
// 参数:
//   s: 原始字符串
// 返回值:
//   string: 小写字符串
func ToLower(s string) string {
	return strings.ToLower(s)
}

// Split 分割字符串
// 参数:
//   s: 要分割的字符串
//   sep: 分隔符
// 返回值:
//   []string: 分割后的字符串切片
func Split(s, sep string) []string {
	return strings.Split(s, sep)
}

// FormatPosition 格式化坐标为字符串
// 参数:
//   x, y: 坐标
// 返回值:
//   string: 格式化后的坐标字符串，如 "x,y"
func FormatPosition(x, y int) string {
	return fmt.Sprintf("%d,%d", x, y)
}

// ParsePosition 解析坐标字符串
// 参数:
//   s: 坐标字符串，如 "x,y"
// 返回值:
//   Point: 解析后的坐标
//   error: 解析错误
func ParsePosition(s string) (Point, error) {
	parts := strings.Split(s, ",")
	if len(parts) != 2 {
		return Point{}, fmt.Errorf("invalid position format: %s", s)
	}
	var x, y int
	_, err := fmt.Sscanf(s, "%d,%d", &x, &y)
	if err != nil {
		return Point{}, err
	}
	return NewPoint(x, y), nil
}

// ParsePositions 解析多个坐标
// 参数:
//   s: 坐标字符串，用 "|" 分隔，如 "x1,y1|x2,y2"
// 返回值:
//   []Point: 解析后的坐标切片
//   error: 解析错误
func ParsePositions(s string) ([]Point, error) {
	parts := Split(s, "|")
	points := make([]Point, 0, len(parts))
	for _, part := range parts {
		p, err := ParsePosition(part)
		if err != nil {
			return nil, err
		}
		points = append(points, p)
	}
	return points, nil
}

// FormatOcrResult 格式化OCR结果为字符串
// 参数:
//   result: OCR结果
// 返回值:
//   string: 格式化后的字符串
func FormatOcrResult(result []OcrRecResult) string {
	var sb strings.Builder
	for i, rec := range result {
		if i > 0 {
			sb.WriteString("|")
		}
		sb.WriteString(fmt.Sprintf("%s,%d,%d", rec.Text, rec.LeftTop.X, rec.LeftTop.Y))
	}
	return sb.String()
}

// IsPrintableChar 检查字符是否可打印
// 参数:
//   r: 要检查的字符
// 返回值:
//   bool: 如果可打印返回true，否则返回false
func IsPrintableChar(r rune) bool {
	return unicode.IsPrint(r)
}

// VirtualKeyCode 虚拟键码映射
type VirtualKeyCode int

// 常用虚拟键码常量
const (
	VK_BACK     VirtualKeyCode = 0x08 // 退格键
	VK_TAB      VirtualKeyCode = 0x09 // Tab键
	VK_RETURN   VirtualKeyCode = 0x0D // 回车键
	VK_SHIFT    VirtualKeyCode = 0x10 // Shift键
	VK_CONTROL  VirtualKeyCode = 0x11 // Ctrl键
	VK_MENU     VirtualKeyCode = 0x12 // Alt键
	VK_CAPITAL  VirtualKeyCode = 0x14 // 大写锁定键
	VK_ESCAPE   VirtualKeyCode = 0x1B // Esc键
	VK_SPACE    VirtualKeyCode = 0x20 // 空格键
	VK_LEFT     VirtualKeyCode = 0x25 // 左箭头键
	VK_UP       VirtualKeyCode = 0x26 // 上箭头键
	VK_RIGHT    VirtualKeyCode = 0x27 // 右箭头键
	VK_DOWN     VirtualKeyCode = 0x28 // 下箭头键
	VK_DELETE   VirtualKeyCode = 0x2E // Delete键
	VK_F1       VirtualKeyCode = 0x70 // F1键
	VK_F2       VirtualKeyCode = 0x71 // F2键
	VK_F3       VirtualKeyCode = 0x72 // F3键
	VK_F4       VirtualKeyCode = 0x73 // F4键
	VK_F5       VirtualKeyCode = 0x74 // F5键
	VK_F6       VirtualKeyCode = 0x75 // F6键
	VK_F7       VirtualKeyCode = 0x76 // F7键
	VK_F8       VirtualKeyCode = 0x77 // F8键
	VK_F9       VirtualKeyCode = 0x78 // F9键
	VK_F10      VirtualKeyCode = 0x79 // F10键
	VK_F11      VirtualKeyCode = 0x7A // F11键
	VK_F12      VirtualKeyCode = 0x7B // F12键
)

// VirtualKeyMap 虚拟键名到键码的映射
var VirtualKeyMap = map[string]VirtualKeyCode{
	"back":    VK_BACK,
	"ctrl":    VK_CONTROL,
	"alt":     VK_MENU,
	"shift":   VK_SHIFT,
	"win":     0x5B,
	"space":   VK_SPACE,
	"cap":     VK_CAPITAL,
	"tab":     VK_TAB,
	"esc":     VK_ESCAPE,
	"enter":   VK_RETURN,
	"up":      VK_UP,
	"down":    VK_DOWN,
	"left":    VK_LEFT,
	"right":   VK_RIGHT,
	"option":  0x5D,
	"print":   0x2C,
	"delete":  VK_DELETE,
	"home":    0x24,
	"end":     0x23,
	"pgup":    0x21,
	"pgdn":    0x22,
	"f1":      VK_F1,
	"f2":      VK_F2,
	"f3":      VK_F3,
	"f4":      VK_F4,
	"f5":      VK_F5,
	"f6":      VK_F6,
	"f7":      VK_F7,
	"f8":      VK_F8,
	"f9":      VK_F9,
	"f10":     VK_F10,
	"f11":     VK_F11,
	"f12":     VK_F12,
}

// GetVirtualKeyCode 从键名获取虚拟键码
// 参数:
//   keyName: 键名（如 "enter", "esc" 等）
// 返回值:
//   VirtualKeyCode: 虚拟键码
//   bool: 是否找到
func GetVirtualKeyCode(keyName string) (VirtualKeyCode, bool) {
	code, ok := VirtualKeyMap[ToLower(keyName)]
	return code, ok
}
