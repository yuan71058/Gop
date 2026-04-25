// Package core 提供核心工具函数
// 该包包含字符串处理、键码映射等工具函数
package core

import (
	"strconv"
	"strings"
)

// KeyMap 按键名称到虚拟键码的映射表
// 用于将按键名称转换为Windows虚拟键码
var KeyMap = map[string]int{
	"backspace": 0x08,
	"tab":       0x09,
	"enter":     0x0D,
	"shift":     0x10,
	"ctrl":      0x11,
	"alt":       0x12,
	"pause":     0x13,
	"caps":      0x14,
	"esc":       0x1B,
	"space":     0x20,
	"pageup":    0x21,
	"pagedown":  0x22,
	"end":       0x23,
	"home":      0x24,
	"left":      0x25,
	"up":        0x26,
	"right":     0x27,
	"down":      0x28,
	"insert":    0x2D,
	"delete":    0x2E,
	"f1":        0x70,
	"f2":        0x71,
	"f3":        0x72,
	"f4":        0x73,
	"f5":        0x74,
	"f6":        0x75,
	"f7":        0x76,
	"f8":        0x77,
	"f9":        0x78,
	"f10":       0x79,
	"f11":       0x7A,
	"f12":       0x7B,
}

// GetKeycode 获取按键的虚拟键码
// 参数:
//
//	key: 按键名称
//
// 返回值:
//
//	int: 虚拟键码，如果未找到返回0
func GetKeycode(key string) int {
	key = strings.ToLower(key)
	if code, ok := KeyMap[key]; ok {
		return code
	}
	// 尝试解析为数字
	if len(key) == 1 {
		return int(key[0])
	}
	return 0
}

// ParsePoint 解析坐标字符串
// 参数:
//
//	str: 坐标字符串 (格式: "x,y")
//
// 返回值:
//
//	Point: 坐标点
//	error: 解析错误
func ParsePoint(str string) (Point, error) {
	parts := strings.Split(str, ",")
	if len(parts) != 2 {
		return Point{}, nil
	}
	x, err := strconv.Atoi(strings.TrimSpace(parts[0]))
	if err != nil {
		return Point{}, err
	}
	y, err := strconv.Atoi(strings.TrimSpace(parts[1]))
	if err != nil {
		return Point{}, err
	}
	return NewPoint(x, y), nil
}

// ParsePoints 解析多个坐标字符串
// 参数:
//
//	str: 坐标字符串 (格式: "x1,y1|x2,y2|...")
//
// 返回值:
//
//	[]Point: 坐标点列表
func ParsePoints(str string) []Point {
	var points []Point
	pairs := strings.Split(str, "|")
	for _, pair := range pairs {
		pair = strings.TrimSpace(pair)
		if pair == "" {
			continue
		}
		p, err := ParsePoint(pair)
		if err == nil {
			points = append(points, p)
		}
	}
	return points
}

// FormatPoints 格式化坐标点列表为字符串
// 参数:
//
//	points: 坐标点列表
//
// 返回值:
//
//	string: 格式化后的字符串 (格式: "x1,y1|x2,y2|...")
func FormatPoints(points []Point) string {
	var parts []string
	for _, p := range points {
		parts = append(parts, strconv.Itoa(p.X)+","+strconv.Itoa(p.Y))
	}
	return strings.Join(parts, "|")
}
