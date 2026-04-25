// Package libop 提供主要的OP库接口
// 该包整合所有功能模块，提供统一的API
package libop

import (
	"github.com/yuan71058/GOP/background"
	"github.com/yuan71058/GOP/core"
	"github.com/yuan71058/GOP/imageproc"
	"github.com/yuan71058/GOP/ocr"
	"github.com/yuan71058/GOP/winapi"
	"syscall"
)

// LibOP 主要的OP库类
// 整合所有功能模块，提供统一的API
type LibOP struct {
	env      *core.Env          // 环境配置
	winapi   *winapi.WinApi     // Windows API封装
	bkproc   *background.Background // 后台操作
	imageproc *imageproc.ImageProc // 图像处理
	ocr      *ocr.OcrManager    // OCR管理器
}

// NewLibOP 创建OP库实例
// 返回值:
//   *LibOP: OP库实例
func NewLibOP() *LibOP {
	op := &LibOP{
		env:       core.NewEnv(),
		winapi:    winapi.NewWinApi(),
		bkproc:    background.NewBackground(),
		imageproc: imageproc.NewImageProc(),
		ocr:       ocr.NewOcrManager(),
	}
	return op
}

// SetPath 设置工作路径
// 参数:
//   path: 工作路径
// 返回值:
//   int: 1表示成功，0表示失败
func (op *LibOP) SetPath(path string) int {
	op.env.SetWorkPath(path)
	op.bkproc.CurrPath = path
	op.imageproc.CurrPath = path
	return 1
}

// GetPath 获取工作路径
// 返回值:
//   string: 工作路径
func (op *LibOP) GetPath() string {
	return op.env.GetWorkPath()
}

// FindWindow 查找窗口
// 参数:
//   className: 窗口类名
//   title: 窗口标题
// 返回值:
//   int: 窗口句柄
func (op *LibOP) FindWindow(className, title string) int {
	return int(op.winapi.FindWindow(className, title))
}

// BindWindow 绑定窗口
// 参数:
//   hwnd: 窗口句柄
//   display: 显示类型
//   mouse: 鼠标类型
//   keypad: 键盘类型
//   mode: 绑定模式
// 返回值:
//   int: 1表示成功，0表示失败
func (op *LibOP) BindWindow(hwnd int, display, mouse, keypad string, mode int) int {
	return op.bkproc.BindWindow(syscall.Handle(hwnd), display, mouse, keypad, mode)
}

// UnBindWindow 解绑窗口
// 返回值:
//   int: 1表示成功，0表示失败
func (op *LibOP) UnBindWindow() int {
	return op.bkproc.UnBindWindow()
}

// MoveTo 移动鼠标
// 参数:
//   x, y: 坐标
// 返回值:
//   int: 1表示成功，0表示失败
func (op *LibOP) MoveTo(x, y int) int {
	return op.bkproc.MoveTo(x, y)
}

// LeftClick 鼠标左键单击
// 返回值:
//   int: 1表示成功，0表示失败
func (op *LibOP) LeftClick() int {
	return op.bkproc.LeftClick()
}

// RightClick 鼠标右键单击
// 返回值:
//   int: 1表示成功，0表示失败
func (op *LibOP) RightClick() int {
	return op.bkproc.RightClick()
}

// KeyPress 按键
// 参数:
//   key: 虚拟键码
// 返回值:
//   int: 1表示成功，0表示失败
func (op *LibOP) KeyPress(key int) int {
	return op.bkproc.KeyPress(key)
}

// SendString 输入字符串
// 参数:
//   str: 字符串
// 返回值:
//   int: 1表示成功，0表示失败
func (op *LibOP) SendString(str string) int {
	return op.bkproc.SendString(str)
}

// FindPic 找图
// 参数:
//   x1, y1, x2, y2: 查找区域
//   picName: 图片名称
//   deltaColor: 颜色偏差
//   sim: 相似度
//   dir: 查找方向
// 返回值:
//   int, int: 找到的坐标
//   int: 1表示找到，0表示未找到
func (op *LibOP) FindPic(x1, y1, x2, y2 int, picName, deltaColor string, sim float64, dir int) (int, int, int) {
	return op.imageproc.FindPic(x1, y1, x2, y2, picName, deltaColor, sim, dir)
}

// FindColor 找色
// 参数:
//   x1, y1, x2, y2: 查找区域
//   color: 颜色值
//   sim: 相似度
//   dir: 查找方向
// 返回值:
//   int, int: 找到的坐标
//   int: 1表示找到，0表示未找到
func (op *LibOP) FindColor(x1, y1, x2, y2 int, color, sim string, dir int) (int, int, int) {
	// 解析相似度
	simVal := 1.0
	// 简化处理
	return op.imageproc.FindColor(x1, y1, x2, y2, color, simVal, dir)
}

// Ocr 识别文字
// 参数:
//   x1, y1, x2, y2: 识别区域
//   color: 颜色过滤
//   sim: 相似度
// 返回值:
//   string: 识别结果
func (op *LibOP) Ocr(x1, y1, x2, y2 int, color, sim string) string {
	// TODO: 实现OCR功能
	return ""
}

// GetCursorPos 获取鼠标位置
// 返回值:
//   int, int: x, y坐标
//   int: 1表示成功，0表示失败
func (op *LibOP) GetCursorPos() (int, int, int) {
	pos, ret := op.bkproc.GetCursorPos()
	return pos.X, pos.Y, ret
}

// GetClientSize 获取窗口客户区大小
// 参数:
//   hwnd: 窗口句柄
// 返回值:
//   int, int: 宽度, 高度
//   int: 1表示成功，0表示失败
func (op *LibOP) GetClientSize(hwnd int) (int, int, int) {
	// TODO: 实现获取窗口大小
	return 0, 0, 0
}

// Ver 获取版本号
// 返回值:
//   string: 版本号
func (op *LibOP) Ver() string {
	return "1.0.0"
}
