// Package core 提供OP项目的核心类型定义和基础数据结构
// 该包包含了整个项目的基础类型，如点、矩形、颜色等
package core

// Point 表示二维空间中的一个点
type Point struct {
	X int // X坐标
	Y int // Y坐标
}

// NewPoint 创建一个新的点
// 参数:
//   x: X坐标
//   y: Y坐标
// 返回值:
//   Point: 新创建的点
func NewPoint(x, y int) Point {
	return Point{X: x, Y: y}
}

// Rect 表示一个矩形区域
// 用于定义屏幕上的一个矩形区域，常用于截图、找图等操作
type Rect struct {
	X1 int // 左上角X坐标
	Y1 int // 左上角Y坐标
	X2 int // 右下角X坐标
	Y2 int // 右下角Y坐标
}

// NewRect 创建一个新的矩形
// 参数:
//   x1, y1: 左上角坐标
//   x2, y2: 右下角坐标
// 返回值:
//   Rect: 新创建的矩形
func NewRect(x1, y1, x2, y2 int) Rect {
	return Rect{X1: x1, Y1: y1, X2: x2, Y2: y2}
}

// Width 计算矩形的宽度
// 返回值:
//   int: 矩形的宽度（像素）
func (r Rect) Width() int {
	return r.X2 - r.X1
}

// Height 计算矩形的高度
// 返回值:
//   int: 矩形的高度（像素）
func (r Rect) Height() int {
	return r.Y2 - r.Y1
}

// Area 计算矩形的面积
// 返回值:
//   int: 矩形的面积（像素数）
func (r Rect) Area() int {
	return r.Width() * r.Height()
}

// Valid 检查矩形是否有效
// 返回值:
//   bool: 如果矩形有效（坐标正确）返回true，否则返回false
func (r Rect) Valid() bool {
	return r.X1 >= 0 && r.X1 < r.X2 && r.Y1 >= 0 && r.Y1 < r.Y2
}

// Color 表示一个RGB颜色值
type Color struct {
	R uint8 // 红色分量 (0-255)
	G uint8 // 绿色分量 (0-255)
	B uint8 // 蓝色分量 (0-255)
}

// NewColor 创建一个新的颜色
// 参数:
//   r, g, b: 红、绿、蓝分量 (0-255)
// 返回值:
//   Color: 新创建的颜色
func NewColor(r, g, b uint8) Color {
	return Color{R: r, G: g, B: b}
}

// ParseColor 从字符串解析颜色
// 支持格式: "RRGGBB" 或 "R,G,B"
// 参数:
//   s: 颜色字符串
// 返回值:
//   Color: 解析后的颜色
//   error: 解析错误（如果有）
func ParseColor(s string) (Color, error) {
	// TODO: 实现颜色解析逻辑
	return Color{}, nil
}

// OcrRecResult 表示OCR识别结果
type OcrRecResult struct {
	LeftTop      Point  // 左上角坐标
	RightBottom  Point  // 右下角坐标
	Text         string // 识别的文本
	Confidence   float32 // 识别置信度 (0.0-1.0)
}

// OcrResult 表示完整的OCR识别结果
type OcrResult struct {
	Records []OcrRecResult // 所有识别结果
}

// ImageData 表示图像数据
type ImageData struct {
	Data   []byte // 像素数据
	Width  int    // 图像宽度
	Height int    // 图像高度
	Format string // 图像格式 (如 "bgra", "bgr")
}

// NewImageData 创建一个新的图像数据
// 参数:
//   data: 像素数据
//   width: 图像宽度
//   height: 图像高度
//   format: 图像格式
// 返回值:
//   *ImageData: 新创建的图像数据
func NewImageData(data []byte, width, height int, format string) *ImageData {
	return &ImageData{
		Data:   data,
		Width:  width,
		Height: height,
		Format: format,
	}
}
