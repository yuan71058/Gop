// Package core 提供核心数据类型和工具函数
// 该包包含Point、Rect、Color等基础类型定义
package core

// Point 二维坐标点
// 用于表示屏幕坐标或图像中的位置
type Point struct {
	X int // X坐标
	Y int // Y坐标
}

// NewPoint 创建新的坐标点
// 参数:
//   x: X坐标
//   y: Y坐标
// 返回值:
//   Point: 坐标点
func NewPoint(x, y int) Point {
	return Point{X: x, Y: y}
}

// Equals 判断两个点是否相等
// 参数:
//   other: 另一个坐标点
// 返回值:
//   bool: 是否相等
func (p Point) Equals(other Point) bool {
	return p.X == other.X && p.Y == other.Y
}

// Rect 矩形区域
// 用于表示屏幕上的一个矩形区域
type Rect struct {
	X1 int // 左上角X坐标
	Y1 int // 左上角Y坐标
	X2 int // 右下角X坐标
	Y2 int // 右下角Y坐标
}

// NewRect 创建新的矩形区域
// 参数:
//   x1, y1: 左上角坐标
//   x2, y2: 右下角坐标
// 返回值:
//   Rect: 矩形区域
func NewRect(x1, y1, x2, y2 int) Rect {
	return Rect{X1: x1, Y1: y1, X2: x2, Y2: y2}
}

// Width 获取矩形宽度
// 返回值:
//   int: 宽度
func (r Rect) Width() int {
	return r.X2 - r.X1
}

// Height 获取矩形高度
// 返回值:
//   int: 高度
func (r Rect) Height() int {
	return r.Y2 - r.Y1
}

// Color 颜色值
// 用于表示RGB颜色
type Color struct {
	R uint8 // 红色分量
	G uint8 // 绿色分量
	B uint8 // 蓝色分量
}

// NewColor 创建新的颜色值
// 参数:
//   r, g, b: RGB分量 (0-255)
// 返回值:
//   Color: 颜色值
func NewColor(r, g, b uint8) Color {
	return Color{R: r, G: g, B: b}
}

// OcrRecResult OCR识别结果
// 包含识别的文字和位置信息
type OcrRecResult struct {
	Text string  // 识别的文字
	X    int     // X坐标
	Y    int     // Y坐标
	Conf float64 // 置信度
}

// ImageData 图像数据
// 用于存储图像像素数据
type ImageData struct {
	Data   []byte // 像素数据
	Width  int    // 宽度
	Height int    // 高度
	Format string // 格式 (bgra, rgb等)
}

// NewImageData 创建新的图像数据
// 参数:
//   data: 像素数据
//   width, height: 宽高
//   format: 格式
// 返回值:
//   *ImageData: 图像数据
func NewImageData(data []byte, width, height int, format string) *ImageData {
	return &ImageData{
		Data:   data,
		Width:  width,
		Height: height,
		Format: format,
	}
}

// OcrResult OCR识别结果集合
// 包含多个识别结果
type OcrResult struct {
	Results []OcrRecResult // 识别结果列表
	Raw     string         // 原始文本
}
