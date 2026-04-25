// Package libop 提供主OP库接口
// 该包集成所有功能模块并提供统一的API
package libop

import (
	"math/rand"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/yuan71058/GOP/algorithm"
	"github.com/yuan71058/GOP/background"
	"github.com/yuan71058/GOP/core"
	"github.com/yuan71058/GOP/imageproc"
	"github.com/yuan71058/GOP/ocr"
	"github.com/yuan71058/GOP/winapi"
)

// LibOP 是主OP库类
// 集成所有功能模块并提供统一的API
type LibOP struct {
	env       *core.Env              // 环境配置
	winapi    *winapi.WinApi         // Windows API封装
	bkproc    *background.Background // 后台操作
	imageproc *imageproc.ImageProc   // 图像处理
	ocr       *ocr.OcrManager        // OCR管理器
	algorithm *algorithm.AStar       // A*算法
	id        int                    // 唯一实例ID
}

// NewLibOP 创建新的OP库实例
// 返回值:
//   *LibOP: OP库实例
func NewLibOP() *LibOP {
	op := &LibOP{
		env:       core.NewEnv(),
		winapi:    winapi.NewWinApi(),
		bkproc:    background.NewBackground(),
		imageproc: imageproc.NewImageProc(),
		ocr:       ocr.NewOcrManager(),
		algorithm: algorithm.NewAStar(),
		id:        core.NextID(),
	}
	return op
}

// ==================== 基础设置/属性 ====================

// Ver 返回库的版本号
// 返回值:
//   string: 版本字符串
func (op *LibOP) Ver() string {
	return "1.0.0"
}

// SetPath 设置工作路径
// 参数:
//   path: 工作目录路径
func (op *LibOP) SetPath(path string) int {
	op.env.SetWorkPath(path)
	op.bkproc.CurrPath = path
	op.imageproc.CurrPath = path
	return 1
}

// GetPath 获取工作路径
// 返回值:
//   string: 工作目录路径
func (op *LibOP) GetPath() string {
	return op.env.GetWorkPath()
}

// GetBasePath 返回插件基础目录
// 返回值:
//   string: 插件基础目录路径
func (op *LibOP) GetBasePath() string {
	return op.env.GetBasePath()
}

// GetID 返回此对象实例的唯一ID
// 此值对每个对象都是唯一的,可用于判断两个对象是否相同
// 返回值:
//   int: 唯一实例ID
func (op *LibOP) GetID() int {
	return op.id
}

// GetLastError 返回最后一次错误代码
// 返回值:
//   int: 最后一次错误代码
func (op *LibOP) GetLastError() int {
	return op.env.GetLastError()
}

// SetShowErrorMsg 设置是否显示错误信息
// 参数:
//   showType: 0=关闭, 1=消息框, 2=保存到文件, 3=输出到标准输出
// 返回值:
//   int: 1表示成功, 0表示失败
func (op *LibOP) SetShowErrorMsg(showType int) int {
	op.env.SetShowErrorMsg(showType)
	return 1
}

// Sleep 延迟指定毫秒数
// 参数:
//   milliseconds: 延迟时间(毫秒)
// 返回值:
//   int: 1表示成功, 0表示失败
func (op *LibOP) Sleep(milliseconds int) int {
	op.env.Sleep(milliseconds)
	return 1
}

// Delay 延迟指定毫秒数(不阻塞UI)
// 参数:
//   ms: 延迟时间(毫秒)
// 返回值:
//   int: 1表示成功, 0表示失败
func (op *LibOP) Delay(ms int) int {
	op.env.Delay(ms)
	return 1
}

// Delays 在指定范围内随机延迟毫秒数(不阻塞UI)
// 参数:
//   msMin: 最小延迟时间(毫秒)
//   msMax: 最大延迟时间(毫秒)
// 返回值:
//   int: 1表示成功, 0表示失败
func (op *LibOP) Delays(msMin, msMax int) int {
	op.env.Delays(msMin, msMax)
	return 1
}

// InjectDll 向进程注入DLL
// 参数:
//   processName: 进程名
//   dllName: 要注入的DLL名
// 返回值:
//   int: 1表示成功, 0表示失败
func (op *LibOP) InjectDll(processName, dllName string) int {
	return op.winapi.InjectDll(processName, dllName)
}

// EnablePicCache 启用或禁用图像缓存
// 参数:
//   enable: 1启用, 0禁用
// 返回值:
//   int: 1表示成功, 0表示失败
func (op *LibOP) EnablePicCache(enable int) int {
	op.imageproc.EnablePicCache(enable)
	return 1
}

// CapturePre 将最后一次捕获的屏幕区域保存到文件(24位位图)
// 参数:
//   fileName: 输出文件名
// 返回值:
//   int: 1表示成功, 0表示失败
func (op *LibOP) CapturePre(fileName string) int {
	return op.imageproc.CapturePre(fileName)
}

// SetScreenDataMode 设置屏幕数据模式
// 参数:
//   mode: 0=从上到下(默认), 1=从下到上
// 返回值:
//   int: 1表示成功, 0表示失败
func (op *LibOP) SetScreenDataMode(mode int) int {
	op.env.SetScreenDataMode(mode)
	return 1
}

// ==================== 算法 ====================

// AStarFindPath 使用A*算法查找路径
// 参数:
//   mapWidth: 地图宽度
//   mapHeight: 地图高度
//   disablePoints: 障碍点字符串, 格式: "x1,y1|x2,y2|..."
//   beginX: 起点X坐标
//   beginY: 起点Y坐标
//   endX: 终点X坐标
//   endY: 终点Y坐标
// 返回值:
//   string: 路径结果, 格式: "x1,y1|x2,y2|..."
func (op *LibOP) AStarFindPath(mapWidth, mapHeight int, disablePoints string, beginX, beginY, endX, endY int) string {
	return op.algorithm.AStarFindPath(mapWidth, mapHeight, disablePoints, beginX, beginY, endX, endY)
}

// FindNearestPos 从所有位置中查找最近的位置
// 参数:
//   allPos: 所有位置字符串, 格式: "x1,y1|x2,y2|..."
//   posType: 位置类型
//   x: 目标X坐标
//   y: 目标Y坐标
// 返回值:
//   string: 最近位置, 格式: "x,y"
func (op *LibOP) FindNearestPos(allPos string, posType, x, y int) string {
	return op.algorithm.FindNearestPos(allPos, posType, x, y)
}

// ==================== Windows API ====================

// EnumWindow 枚举符合指定条件的窗口
// 参数:
//   parent: 父窗口句柄(0为桌面)
//   title: 窗口标题(可为空)
//   className: 窗口类名(可为空)
//   filter: 过滤标志
// 返回值:
//   string: 窗口句柄字符串, 格式: "hwnd1,hwnd2|..."
func (op *LibOP) EnumWindow(parent int, title, className string, filter int) string {
	return op.winapi.EnumWindow(parent, title, className, filter)
}

// EnumWindowByProcess 枚举符合指定进程和条件的窗口
// 参数:
//   processName: 进程名
//   title: 窗口标题(可为空)
//   className: 窗口类名(可为空)
//   filter: 过滤标志
// 返回值:
//   string: 窗口句柄字符串, 格式: "hwnd1,hwnd2|..."
func (op *LibOP) EnumWindowByProcess(processName, title, className string, filter int) string {
	return op.winapi.EnumWindowByProcess(processName, title, className, filter)
}

// EnumProcess 枚举符合指定进程名的进程PID
// 参数:
//   name: 进程名
// 返回值:
//   string: 进程ID字符串, 格式: "pid1|pid2|..."
func (op *LibOP) EnumProcess(name string) string {
	return op.winapi.EnumProcess(name)
}

// ClientToScreen 将客户区坐标转换为屏幕坐标
// 参数:
//   hwnd: 窗口句柄
//   x: X坐标(输入/输出)
//   y: Y坐标(输入/输出)
// 返回值:
//   int: 1表示成功, 0表示失败
func (op *LibOP) ClientToScreen(hwnd int, x, y *int) int {
	return op.winapi.ClientToScreen(hwnd, x, y)
}

// FindWindow 查找符合类名或者标题名的顶层可见窗口
// 参数:
//   className: 窗口类名(可为空)
//   title: 窗口标题(可为空)
// 返回值:
//   int: 窗口句柄, 未找到返回0
func (op *LibOP) FindWindow(className, title string) int {
	return int(op.winapi.FindWindow(className, title))
}

// FindWindowByProcess 通过进程名查找可见窗口
// 参数:
//   processName: 进程名
//   className: 窗口类名(可为空)
//   title: 窗口标题(可为空)
// 返回值:
//   int: 窗口句柄, 未找到返回0
func (op *LibOP) FindWindowByProcess(processName, className, title string) int {
	return op.winapi.FindWindowByProcess(processName, className, title)
}

// FindWindowByProcessId 通过进程ID查找可见窗口
// 参数:
//   processId: 进程ID
//   className: 窗口类名(可为空)
//   title: 窗口标题(可为空)
// 返回值:
//   int: 窗口句柄, 未找到返回0
func (op *LibOP) FindWindowByProcessId(processId int, className, title string) int {
	return op.winapi.FindWindowByProcessId(processId, className, title)
}

// FindWindowEx 查找符合类名或者标题名的顶层可见窗口,如果指定parent则在parent的一级子窗口中查找
// 参数:
//   parent: 父窗口句柄(0为桌面)
//   className: 窗口类名(可为空)
//   title: 窗口标题(可为空)
// 返回值:
//   int: 窗口句柄, 未找到返回0
func (op *LibOP) FindWindowEx(parent int, className, title string) int {
	return int(op.winapi.FindWindowEx(parent, className, title))
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
func (op *LibOP) GetClientRect(hwnd int, x1, y1, x2, y2 *int) int {
	return op.winapi.GetClientRect(hwnd, x1, y1, x2, y2)
}

// GetClientSize 获取窗口客户区的宽度和高度
// 参数:
//   hwnd: 窗口句柄
//   width: 宽度(输出)
//   height: 高度(输出)
// 返回值:
//   int: 1表示成功, 0表示失败
func (op *LibOP) GetClientSize(hwnd int, width, height *int) int {
	return op.winapi.GetClientSize(hwnd, width, height)
}

// GetForegroundFocus 获取顶层活动窗口中具有输入焦点的窗口句柄
// 返回值:
//   int: 具有焦点的窗口句柄
func (op *LibOP) GetForegroundFocus() int {
	return op.winapi.GetForegroundFocus()
}

// GetForegroundWindow 获取顶层活动窗口
// 返回值:
//   int: 窗口句柄
func (op *LibOP) GetForegroundWindow() int {
	return op.winapi.GetForegroundWindow()
}

// GetMousePointWindow 获取鼠标光标指向的可见窗口句柄
// 返回值:
//   int: 窗口句柄
func (op *LibOP) GetMousePointWindow() int {
	return op.winapi.GetMousePointWindow()
}

// GetPointWindow 获取指定坐标处的可见窗口句柄
// 参数:
//   x: X坐标
//   y: Y坐标
// 返回值:
//   int: 窗口句柄
func (op *LibOP) GetPointWindow(x, y int) int {
	return op.winapi.GetPointWindow(x, y)
}

// GetProcessInfo 获取指定PID的详细信息
// 参数:
//   pid: 进程ID
// 返回值:
//   string: 进程信息字符串
func (op *LibOP) GetProcessInfo(pid int) string {
	return op.winapi.GetProcessInfo(pid)
}

// GetSpecialWindow 获取特殊窗口句柄
// 参数:
//   flag: 0=桌面, 1=任务栏
// 返回值:
//   int: 窗口句柄
func (op *LibOP) GetSpecialWindow(flag int) int {
	return op.winapi.GetSpecialWindow(flag)
}

// GetWindow 获取与给定窗口有关的窗口句柄
// 参数:
//   hwnd: 窗口句柄
//   flag: 0=父窗口, 1=第一个子窗口, 2=下一个兄弟窗口, 3=上一个兄弟窗口
// 返回值:
//   int: 窗口句柄
func (op *LibOP) GetWindow(hwnd, flag int) int {
	return op.winapi.GetWindow(hwnd, flag)
}

// GetWindowClass 获取窗口的类名
// 参数:
//   hwnd: 窗口句柄
// 返回值:
//   string: 窗口类名
func (op *LibOP) GetWindowClass(hwnd int) string {
	return op.winapi.GetWindowClass(hwnd)
}

// GetWindowProcessId 获取指定窗口的进程ID
// 参数:
//   hwnd: 窗口句柄
// 返回值:
//   int: 进程ID
func (op *LibOP) GetWindowProcessId(hwnd int) int {
	return op.winapi.GetWindowProcessId(hwnd)
}

// GetWindowProcessPath 获取指定窗口所在进程的exe文件全路径
// 参数:
//   hwnd: 窗口句柄
// 返回值:
//   string: 进程可执行文件路径
func (op *LibOP) GetWindowProcessPath(hwnd int) string {
	return op.winapi.GetWindowProcessPath(hwnd)
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
func (op *LibOP) GetWindowRect(hwnd int, x1, y1, x2, y2 *int) int {
	return op.winapi.GetWindowRect(hwnd, x1, y1, x2, y2)
}

// GetWindowState 获取指定窗口的某些属性
// 参数:
//   hwnd: 窗口句柄
//   flag: 0=可见, 1=活动, 2=最小化, 3=最大化, 4=启用
// 返回值:
//   int: 窗口状态值
func (op *LibOP) GetWindowState(hwnd, flag int) int {
	return op.winapi.GetWindowState(hwnd, flag)
}

// GetWindowTitle 获取窗口的标题
// 参数:
//   hwnd: 窗口句柄
// 返回值:
//   string: 窗口标题
func (op *LibOP) GetWindowTitle(hwnd int) string {
	return op.winapi.GetWindowTitle(hwnd)
}

// MoveWindow 移动指定窗口到指定位置
// 参数:
//   hwnd: 窗口句柄
//   x: X坐标
//   y: Y坐标
// 返回值:
//   int: 1表示成功, 0表示失败
func (op *LibOP) MoveWindow(hwnd, x, y int) int {
	return op.winapi.MoveWindow(hwnd, x, y)
}

// ScreenToClient 将屏幕坐标转换为客户区坐标
// 参数:
//   hwnd: 窗口句柄
//   x: X坐标(输入/输出)
//   y: Y坐标(输入/输出)
// 返回值:
//   int: 1表示成功, 0表示失败
func (op *LibOP) ScreenToClient(hwnd int, x, y *int) int {
	return op.winapi.ScreenToClient(hwnd, x, y)
}

// SendPaste 向指定窗口发送粘贴命令
// 参数:
//   hwnd: 窗口句柄
// 返回值:
//   int: 1表示成功, 0表示失败
func (op *LibOP) SendPaste(hwnd int) int {
	return op.winapi.SendPaste(hwnd)
}

// SetClientSize 设置窗口客户区的宽度和高度
// 参数:
//   hwnd: 窗口句柄
//   width: 宽度
//   height: 高度
// 返回值:
//   int: 1表示成功, 0表示失败
func (op *LibOP) SetClientSize(hwnd, width, height int) int {
	return op.winapi.SetClientSize(hwnd, width, height)
}

// SetWindowState 设置窗口的状态
// 参数:
//   hwnd: 窗口句柄
//   flag: 0=显示, 1=隐藏, 2=最小化, 3=最大化, 4=还原, 5=激活, 6=关闭
// 返回值:
//   int: 1表示成功, 0表示失败
func (op *LibOP) SetWindowState(hwnd, flag int) int {
	return op.winapi.SetWindowState(hwnd, flag)
}

// SetWindowSize 设置窗口的大小
// 参数:
//   hwnd: 窗口句柄
//   width: 宽度
//   height: 高度
// 返回值:
//   int: 1表示成功, 0表示失败
func (op *LibOP) SetWindowSize(hwnd, width, height int) int {
	return op.winapi.SetWindowSize(hwnd, width, height)
}

// SetWindowText 设置窗口的标题
// 参数:
//   hwnd: 窗口句柄
//   title: 窗口标题
// 返回值:
//   int: 1表示成功, 0表示失败
func (op *LibOP) SetWindowText(hwnd int, title string) int {
	return op.winapi.SetWindowText(hwnd, title)
}

// SetWindowTransparent 设置窗口的透明度
// 参数:
//   hwnd: 窗口句柄
//   trans: 透明度值(0-255)
// 返回值:
//   int: 1表示成功, 0表示失败
func (op *LibOP) SetWindowTransparent(hwnd, trans int) int {
	return op.winapi.SetWindowTransparent(hwnd, trans)
}

// SendString 向指定窗口发送文本数据
// 参数:
//   hwnd: 窗口句柄
//   str: 文本字符串
// 返回值:
//   int: 1表示成功, 0表示失败
func (op *LibOP) SendString(hwnd int, str string) int {
	return op.winapi.SendString(hwnd, str)
}

// SendStringIme 使用IME向指定窗口发送文本数据
// 参数:
//   hwnd: 窗口句柄
//   str: 文本字符串
// 返回值:
//   int: 1表示成功, 0表示失败
func (op *LibOP) SendStringIme(hwnd int, str string) int {
	return op.winapi.SendStringIme(hwnd, str)
}

// RunApp 以指定模式运行可执行文件
// 参数:
//   cmdline: 命令行
//   mode: 运行模式(0=普通, 1=显示, 2=隐藏)
// 返回值:
//   int: 1表示成功, 0表示失败
func (op *LibOP) RunApp(cmdline string, mode int) int {
	return op.winapi.RunApp(cmdline, mode)
}

// WinExec 以指定显示模式运行可执行文件
// 参数:
//   cmdline: 命令行
//   cmdShow: 显示模式(SW_SHOW, SW_HIDE等)
// 返回值:
//   int: 1表示成功, 0表示失败
func (op *LibOP) WinExec(cmdline string, cmdShow int) int {
	return op.winapi.WinExec(cmdline, cmdShow)
}

// GetCmdStr 运行命令行并返回结果
// 参数:
//   cmd: 命令字符串
//   milliseconds: 超时时间(毫秒)
// 返回值:
//   string: 命令输出
func (op *LibOP) GetCmdStr(cmd string, milliseconds int) string {
	return op.winapi.GetCmdStr(cmd, milliseconds)
}

// SetClipboard 设置剪贴板数据
// 参数:
//   str: 文本字符串
// 返回值:
//   int: 1表示成功, 0表示失败
func (op *LibOP) SetClipboard(str string) int {
	return op.winapi.SetClipboard(str)
}

// GetClipboard 获取剪贴板数据
// 返回值:
//   string: 剪贴板文本
func (op *LibOP) GetClipboard() string {
	return op.winapi.GetClipboard()
}

// ==================== 后台操作 ====================

// BindWindow 绑定窗口并开始屏幕捕获
// 参数:
//   hwnd: 窗口句柄
//   display: 显示模式("normal", "gdi", "dx", "dx2", "opengl"等)
//   mouse: 鼠标模式("normal", "windows", "windows3", "dx"等)
//   keypad: 键盘模式("normal", "windows", "dx"等)
//   mode: 绑定模式(0=普通, 1=后台等)
// 返回值:
//   int: 1表示成功, 0表示失败
func (op *LibOP) BindWindow(hwnd int, display, mouse, keypad string, mode int) int {
	return op.bkproc.BindWindow(syscall.Handle(hwnd), display, mouse, keypad, mode)
}

// UnBindWindow 解绑窗口
// 返回值:
//   int: 1表示成功, 0表示失败
func (op *LibOP) UnBindWindow() int {
	return op.bkproc.UnBindWindow()
}

// GetBindWindow 获取此对象当前绑定的窗口句柄
// 返回值:
//   int: 窗口句柄, 未绑定返回0
func (op *LibOP) GetBindWindow() int {
	return int(op.bkproc.GetBindWindow())
}

// IsBind 判断当前对象是否绑定了窗口
// 返回值:
//   int: 1表示已绑定, 0表示未绑定
func (op *LibOP) IsBind() int {
	return op.bkproc.IsBind()
}

// ==================== 鼠标与键盘 ====================

// GetCursorPos 获取鼠标位置
// 参数:
//   x: X坐标(输出)
//   y: Y坐标(输出)
// 返回值:
//   int: 1表示成功, 0表示失败
func (op *LibOP) GetCursorPos(x, y *int) int {
	pos, ret := op.bkproc.GetCursorPos()
	if ret == 1 {
		*x = pos.X
		*y = pos.Y
	}
	return ret
}

// MoveR 相对上次位置移动鼠标
// 参数:
//   x: 相对X偏移
//   y: 相对Y偏移
// 返回值:
//   int: 1表示成功, 0表示失败
func (op *LibOP) MoveR(x, y int) int {
	return op.bkproc.MoveR(x, y)
}

// MoveTo 移动鼠标到目标点(x, y)
// 参数:
//   x: X坐标
//   y: Y坐标
// 返回值:
//   int: 1表示成功, 0表示失败
func (op *LibOP) MoveTo(x, y int) int {
	return op.bkproc.MoveTo(x, y)
}

// MoveToEx 移动鼠标到指定范围内的任意点
// 参数:
//   x: X坐标
//   y: Y坐标
//   w: 宽度范围
//   h: 高度范围
// 返回值:
//   int: 1表示成功, 0表示失败
func (op *LibOP) MoveToEx(x, y, w, h int) int {
	return op.bkproc.MoveToEx(x, y, w, h)
}

// LeftClick 点击鼠标左键
// 返回值:
//   int: 1表示成功, 0表示失败
func (op *LibOP) LeftClick() int {
	return op.bkproc.LeftClick()
}

// LeftDoubleClick 双击鼠标左键
// 返回值:
//   int: 1表示成功, 0表示失败
func (op *LibOP) LeftDoubleClick() int {
	return op.bkproc.LeftDoubleClick()
}

// LeftDown 按住鼠标左键
// 返回值:
//   int: 1表示成功, 0表示失败
func (op *LibOP) LeftDown() int {
	return op.bkproc.LeftDown()
}

// LeftUp 释放鼠标左键
// 返回值:
//   int: 1表示成功, 0表示失败
func (op *LibOP) LeftUp() int {
	return op.bkproc.LeftUp()
}

// MiddleClick 点击鼠标中键
// 返回值:
//   int: 1表示成功, 0表示失败
func (op *LibOP) MiddleClick() int {
	return op.bkproc.MiddleClick()
}

// MiddleDown 按住鼠标中键
// 返回值:
//   int: 1表示成功, 0表示失败
func (op *LibOP) MiddleDown() int {
	return op.bkproc.MiddleDown()
}

// MiddleUp 释放鼠标中键
// 返回值:
//   int: 1表示成功, 0表示失败
func (op *LibOP) MiddleUp() int {
	return op.bkproc.MiddleUp()
}

// RightClick 点击鼠标右键
// 返回值:
//   int: 1表示成功, 0表示失败
func (op *LibOP) RightClick() int {
	return op.bkproc.RightClick()
}

// RightDown 按住鼠标右键
// 返回值:
//   int: 1表示成功, 0表示失败
func (op *LibOP) RightDown() int {
	return op.bkproc.RightDown()
}

// RightUp 释放鼠标右键
// 返回值:
//   int: 1表示成功, 0表示失败
func (op *LibOP) RightUp() int {
	return op.bkproc.RightUp()
}

// WheelDown 向下滚动鼠标滚轮
// 返回值:
//   int: 1表示成功, 0表示失败
func (op *LibOP) WheelDown() int {
	return op.bkproc.WheelDown()
}

// WheelUp 向上滚动鼠标滚轮
// 返回值:
//   int: 1表示成功, 0表示失败
func (op *LibOP) WheelUp() int {
	return op.bkproc.WheelUp()
}

// SetMouseDelay 设置鼠标点击的延迟时间
// 参数:
//   typeStr: 类型字符串
//   delay: 延迟时间(毫秒)
// 返回值:
//   int: 1表示成功, 0表示失败
func (op *LibOP) SetMouseDelay(typeStr string, delay int) int {
	return op.bkproc.SetMouseDelay(typeStr, delay)
}

// GetKeyState 获取指定按键的状态
// 参数:
//   vkCode: 虚拟键码
// 返回值:
//   int: 按键状态(1=按下, 0=释放)
func (op *LibOP) GetKeyState(vkCode int) int {
	return op.bkproc.GetKeyState(vkCode)
}

// KeyDown 按住指定的虚拟键
// 参数:
//   vkCode: 虚拟键码
// 返回值:
//   int: 1表示成功, 0表示失败
func (op *LibOP) KeyDown(vkCode int) int {
	return op.bkproc.KeyDown(vkCode)
}

// KeyDownChar 通过字符按住指定的键
// 参数:
//   vkCode: 键字符字符串
// 返回值:
//   int: 1表示成功, 0表示失败
func (op *LibOP) KeyDownChar(vkCode string) int {
	return op.bkproc.KeyDownChar(vkCode)
}

// KeyUp 释放指定的虚拟键
// 参数:
//   vkCode: 虚拟键码
// 返回值:
//   int: 1表示成功, 0表示失败
func (op *LibOP) KeyUp(vkCode int) int {
	return op.bkproc.KeyUp(vkCode)
}

// KeyUpChar 通过字符释放指定的键
// 参数:
//   vkCode: 键字符字符串
// 返回值:
//   int: 1表示成功, 0表示失败
func (op *LibOP) KeyUpChar(vkCode string) int {
	return op.bkproc.KeyUpChar(vkCode)
}

// WaitKey 等待指定的键被按下
// 参数:
//   vkCode: 虚拟键码
//   timeOut: 超时时间(毫秒)
// 返回值:
//   int: 1表示成功, 0表示失败
func (op *LibOP) WaitKey(vkCode, timeOut int) int {
	return op.bkproc.WaitKey(vkCode, timeOut)
}

// KeyPress 按下并释放指定的键
// 参数:
//   vkCode: 虚拟键码
// 返回值:
//   int: 1表示成功, 0表示失败
func (op *LibOP) KeyPress(vkCode int) int {
	return op.bkproc.KeyPress(vkCode)
}

// KeyPressChar 通过字符按下并释放指定的键
// 参数:
//   vkCode: 键字符字符串
// 返回值:
//   int: 1表示成功, 0表示失败
func (op *LibOP) KeyPressChar(vkCode string) int {
	return op.bkproc.KeyPressChar(vkCode)
}

// SetKeypadDelay 设置键盘输入的延迟时间
// 参数:
//   typeStr: 类型字符串
//   delay: 延迟时间(毫秒)
// 返回值:
//   int: 1表示成功, 0表示失败
func (op *LibOP) SetKeypadDelay(typeStr string, delay int) int {
	return op.bkproc.SetKeypadDelay(typeStr, delay)
}

// KeyPressStr 按照指定的字符串依次按键
// 参数:
//   keyStr: 键字符串
//   delay: 延迟时间(毫秒)
// 返回值:
//   int: 1表示成功, 0表示失败
func (op *LibOP) KeyPressStr(keyStr string, delay int) int {
	return op.bkproc.KeyPressStr(keyStr, delay)
}

// ==================== 图像与颜色 ====================

// Capture 捕获屏幕区域并保存到文件
// 参数:
//   x1, y1, x2, y2: 捕获区域坐标
//   fileName: 输出文件名
// 返回值:
//   int: 1表示成功, 0表示失败
func (op *LibOP) Capture(x1, y1, x2, y2 int, fileName string) int {
	return op.imageproc.Capture(x1, y1, x2, y2, fileName)
}

// CmpColor 比较指定坐标处的颜色
// 参数:
//   x, y: 坐标
//   color: 要比较的颜色字符串
//   sim: 相似度 (0.1-1.0)
// 返回值:
//   int: 1表示匹配, 0表示不匹配
func (op *LibOP) CmpColor(x, y int, color string, sim float64) int {
	return op.imageproc.CmpColor(x, y, color, sim)
}

// FindColor 在屏幕区域中查找指定颜色
// 参数:
//   x1, y1, x2, y2: 搜索区域坐标
//   color: 颜色值
//   sim: 相似度 (0.1-1.0)
//   dir: 搜索方向
//   x: X坐标(输出)
//   y: Y坐标(输出)
// 返回值:
//   int: 1表示找到, 0表示未找到
func (op *LibOP) FindColor(x1, y1, x2, y2 int, color string, sim float64, dir int, x, y *int) int {
	retX, retY, ret := op.imageproc.FindColor(x1, y1, x2, y2, color, sim, dir)
	if ret == 1 {
		*x = retX
		*y = retY
	}
	return ret
}

// FindColorEx 在屏幕区域中查找所有匹配的颜色
// 参数:
//   x1, y1, x2, y2: 搜索区域坐标
//   color: 颜色值
//   sim: 相似度 (0.1-1.0)
//   dir: 搜索方向
// 返回值:
//   string: 所有找到的坐标(格式: "x1,y1|x2,y2|...")
func (op *LibOP) FindColorEx(x1, y1, x2, y2 int, color string, sim float64, dir int) string {
	return op.imageproc.FindColorEx(x1, y1, x2, y2, color, sim, dir)
}

// GetColorNum 获取区域中匹配颜色的数量
// 参数:
//   x1, y1, x2, y2: 搜索区域坐标
//   color: 颜色值
//   sim: 相似度 (0.1-1.0)
// 返回值:
//   int: 颜色数量
func (op *LibOP) GetColorNum(x1, y1, x2, y2 int, color string, sim float64) int {
	return op.imageproc.GetColorNum(x1, y1, x2, y2, color, sim)
}

// FindMultiColor 在屏幕区域中查找多点颜色
// 参数:
//   x1, y1, x2, y2: 搜索区域坐标
//   firstColor: 第一个颜色
//   offsetColor: 偏移颜色
//   sim: 相似度 (0.1-1.0)
//   dir: 搜索方向
//   x: X坐标(输出)
//   y: Y坐标(输出)
// 返回值:
//   int: 1表示找到, 0表示未找到
func (op *LibOP) FindMultiColor(x1, y1, x2, y2 int, firstColor, offsetColor string, sim float64, dir int, x, y *int) int {
	retX, retY, ret := op.imageproc.FindMultiColor(x1, y1, x2, y2, firstColor, offsetColor, sim, dir)
	if ret == 1 {
		*x = retX
		*y = retY
	}
	return ret
}

// FindMultiColorEx 查找所有匹配多点颜色的坐标
// 参数:
//   x1, y1, x2, y2: 搜索区域坐标
//   firstColor: 第一个颜色
//   offsetColor: 偏移颜色
//   sim: 相似度 (0.1-1.0)
//   dir: 搜索方向
// 返回值:
//   string: 所有找到的坐标
func (op *LibOP) FindMultiColorEx(x1, y1, x2, y2 int, firstColor, offsetColor string, sim float64, dir int) string {
	return op.imageproc.FindMultiColorEx(x1, y1, x2, y2, firstColor, offsetColor, sim, dir)
}

// FindPic 在屏幕区域中查找图像
// 参数:
//   x1, y1, x2, y2: 搜索区域坐标
//   files: 图像文件名
//   deltaColor: 颜色偏差
//   sim: 相似度 (0.1-1.0)
//   dir: 搜索方向
//   x: X坐标(输出)
//   y: Y坐标(输出)
// 返回值:
//   int: 1表示找到, 0表示未找到
func (op *LibOP) FindPic(x1, y1, x2, y2 int, files, deltaColor string, sim float64, dir int, x, y *int) int {
	retX, retY, ret := op.imageproc.FindPic(x1, y1, x2, y2, files, deltaColor, sim, dir)
	if ret == 1 {
		*x = retX
		*y = retY
	}
	return ret
}

// FindPicEx 在屏幕区域中查找所有匹配的图像
// 参数:
//   x1, y1, x2, y2: 搜索区域坐标
//   files: 图像文件名
//   deltaColor: 颜色偏差
//   sim: 相似度 (0.1-1.0)
//   dir: 搜索方向
// 返回值:
//   string: 所有找到的坐标(格式: "x1,y1|x2,y2|...")
func (op *LibOP) FindPicEx(x1, y1, x2, y2 int, files, deltaColor string, sim float64, dir int) string {
	return op.imageproc.FindPicEx(x1, y1, x2, y2, files, deltaColor, sim, dir)
}

// FindPicExS 查找多个图像并返回所有找到的图像坐标
// 参数:
//   x1, y1, x2, y2: 搜索区域坐标
//   files: 图像文件名
//   deltaColor: 颜色偏差
//   sim: 相似度 (0.1-1.0)
//   dir: 搜索方向
// 返回值:
//   string: 所有找到的坐标(格式: "file1,x,y|file2,x,y|...")
func (op *LibOP) FindPicExS(x1, y1, x2, y2 int, files, deltaColor string, sim float64, dir int) string {
	return op.imageproc.FindPicExS(x1, y1, x2, y2, files, deltaColor, sim, dir)
}

// FindColorBlock 在屏幕区域中查找色块
// 参数:
//   x1, y1, x2, y2: 搜索区域坐标
//   color: 颜色值(格式: "RRGGBB-DRDGDB")
//   sim: 相似度 (0.1-1.0)
//   count: 最小颜色数量
//   height: 色块高度
//   width: 色块宽度
//   x: X坐标(输出)
//   y: Y坐标(输出)
// 返回值:
//   int: 1表示找到, 0表示未找到
func (op *LibOP) FindColorBlock(x1, y1, x2, y2 int, color string, sim float64, count, height, width int, x, y *int) int {
	return op.imageproc.FindColorBlock(x1, y1, x2, y2, color, sim, count, height, width, x, y)
}

// FindColorBlockEx 在屏幕区域中查找所有色块
// 参数:
//   x1, y1, x2, y2: 搜索区域坐标
//   color: 颜色值(格式: "RRGGBB-DRDGDB")
//   sim: 相似度 (0.1-1.0)
//   count: 最小颜色数量
//   height: 色块高度
//   width: 色块宽度
// 返回值:
//   string: 所有找到的色块坐标
func (op *LibOP) FindColorBlockEx(x1, y1, x2, y2 int, color string, sim float64, count, height, width int) string {
	return op.imageproc.FindColorBlockEx(x1, y1, x2, y2, color, sim, count, height, width)
}

// GetColor 获取指定坐标处的颜色
// 参数:
//   x, y: 坐标
// 返回值:
//   string: 颜色值(格式: "RRGGBB")
func (op *LibOP) GetColor(x, y int) string {
	return op.imageproc.GetColor(x, y)
}

// SetDisplayInput 设置图像输入模式
// 参数:
//   mode: 模式字符串
// 返回值:
//   int: 1表示成功, 0表示失败
func (op *LibOP) SetDisplayInput(mode string) int {
	return op.imageproc.SetDisplayInput(mode)
}

// LoadPic 加载图像
// 参数:
//   fileName: 图像文件名
// 返回值:
//   int: 1表示成功, 0表示失败
func (op *LibOP) LoadPic(fileName string) int {
	return op.imageproc.LoadPic(fileName)
}

// FreePic 从内存中释放图像
// 参数:
//   fileName: 图像文件名
// 返回值:
//   int: 1表示成功, 0表示失败
func (op *LibOP) FreePic(fileName string) int {
	return op.imageproc.FreePic(fileName)
}

// LoadMemPic 从内存加载图像
// 参数:
//   fileName: 图像名
//   data: 图像数据
//   size: 数据大小
// 返回值:
//   int: 1表示成功, 0表示失败
func (op *LibOP) LoadMemPic(fileName string, data []byte, size int) int {
	return op.imageproc.LoadMemPic(fileName, data, size)
}

// GetPicSize 获取指定图像的大小
// 参数:
//   picName: 图像名
//   width: 宽度(输出)
//   height: 高度(输出)
// 返回值:
//   int: 1表示成功, 0表示失败
func (op *LibOP) GetPicSize(picName string, width, height *int) int {
	return op.imageproc.GetPicSize(picName, width, height)
}

// GetScreenData 获取指定区域的屏幕数据
// 参数:
//   x1, y1, x2, y2: 区域坐标
// 返回值:
//   uintptr: 数据指针
//   int: 1表示成功, 0表示失败
func (op *LibOP) GetScreenData(x1, y1, x2, y2 int) (uintptr, int) {
	var data uintptr
	ret := op.imageproc.GetScreenData(x1, y1, x2, y2, &data)
	return data, ret
}

// GetScreenDataBmp 获取BMP格式的屏幕数据
// 参数:
//   x1, y1, x2, y2: 区域坐标
// 返回值:
//   uintptr: 数据指针
//   int: 数据大小
//   int: 1表示成功, 0表示失败
func (op *LibOP) GetScreenDataBmp(x1, y1, x2, y2 int) (uintptr, int, int) {
	var data uintptr
	var size int
	ret := op.imageproc.GetScreenDataBmp(x1, y1, x2, y2, &data, &size)
	return data, size, ret
}

// GetScreenFrameInfo 获取屏幕帧信息
// 返回值:
//   int: 帧ID
//   int: 帧时间
func (op *LibOP) GetScreenFrameInfo() (int, int) {
	var frameID, frameTime int
	op.imageproc.GetScreenFrameInfo(&frameID, &frameTime)
	return frameID, frameTime
}

// MatchPicName 通过模式匹配图像名
// 参数:
//   picName: 图像名模式
// 返回值:
//   string: 匹配的图像名
func (op *LibOP) MatchPicName(picName string) string {
	return op.imageproc.MatchPicName(picName)
}

// ==================== OCR识别 ====================

// SetOcrEngine 设置OCR引擎路径和配置
// 参数:
//   pathOfEngine: OCR引擎路径
//   dllName: DLL名称
//   argv: 参数
// 返回值:
//   int: 1表示成功, 0表示失败
func (op *LibOP) SetOcrEngine(pathOfEngine, dllName, argv string) int {
	return op.ocr.SetOcrEngine(pathOfEngine, dllName, argv)
}

// SetDict 设置指定索引的字库文件
// 参数:
//   idx: 字库索引 (0-9)
//   fileName: 字库文件名
// 返回值:
//   int: 1表示成功, 0表示失败
func (op *LibOP) SetDict(idx int, fileName string) int {
	return op.ocr.SetDict(idx, fileName)
}

// GetDict 获取指定索引和字体条目的字库信息
// 参数:
//   idx: 字库索引
//   fontIndex: 字体条目索引
// 返回值:
//   string: 字库信息字符串
func (op *LibOP) GetDict(idx, fontIndex int) string {
	return op.ocr.GetDict(idx, fontIndex)
}

// SetMemDict 设置指定索引的内存字库
// 参数:
//   idx: 字库索引 (0-9)
//   data: 字库数据
//   size: 数据大小
// 返回值:
//   int: 1表示成功, 0表示失败
func (op *LibOP) SetMemDict(idx int, data []byte, size int) int {
	return op.ocr.SetMemDict(idx, data, size)
}

// UseDict 设置当前使用的字库
// 参数:
//   idx: 字库索引
// 返回值:
//   int: 1表示成功, 0表示失败
func (op *LibOP) UseDict(idx int) int {
	return op.ocr.UseDict(idx)
}

// AddDict 向指定索引添加字库信息
// 参数:
//   idx: 字库索引
//   dictInfo: 字库信息字符串
// 返回值:
//   int: 1表示成功, 0表示失败
func (op *LibOP) AddDict(idx int, dictInfo string) int {
	return op.ocr.AddDict(idx, dictInfo)
}

// SaveDict 保存指定字库到文件
// 参数:
//   idx: 字库索引
//   fileName: 输出文件名
// 返回值:
//   int: 1表示成功, 0表示失败
func (op *LibOP) SaveDict(idx int, fileName string) int {
	return op.ocr.SaveDict(idx, fileName)
}

// ClearDict 清空指定字库
// 参数:
//   idx: 字库索引
// 返回值:
//   int: 1表示成功, 0表示失败
func (op *LibOP) ClearDict(idx int) int {
	return op.ocr.ClearDict(idx)
}

// GetDictCount 获取指定字库中的字符数量
// 参数:
//   idx: 字库索引
// 返回值:
//   int: 字符数量
func (op *LibOP) GetDictCount(idx int) int {
	return op.ocr.GetDictCount(idx)
}

// GetNowDict 获取当前字库索引
// 返回值:
//   int: 当前字库索引
func (op *LibOP) GetNowDict() int {
	return op.ocr.GetNowDict()
}

// FetchWord 提取指定字符的点阵信息
// 参数:
//   x1, y1, x2, y2: 屏幕区域坐标
//   color: 颜色格式
//   word: 要提取的字符串
// 返回值:
//   string: 提取的点阵信息
func (op *LibOP) FetchWord(x1, y1, x2, y2 int, color, word string) string {
	return op.ocr.FetchWord(x1, y1, x2, y2, color, word)
}

// GetWordsNoDict 不使用字库识别指定区域的词组
// 参数:
//   x1, y1, x2, y2: 屏幕区域坐标
//   color: 颜色格式
// 返回值:
//   string: 识别的词组信息
func (op *LibOP) GetWordsNoDict(x1, y1, x2, y2 int, color string) string {
	return op.ocr.GetWordsNoDict(x1, y1, x2, y2, color)
}

// GetWordResultCount 获取识别词组结果的数量
// 参数:
//   result: 识别结果字符串
// 返回值:
//   int: 词组数量
func (op *LibOP) GetWordResultCount(result string) int {
	return op.ocr.GetWordResultCount(result)
}

// GetWordResultPos 获取识别词组的坐标
// 参数:
//   result: 识别结果字符串
//   index: 词组索引
//   x: X坐标(输出)
//   y: Y坐标(输出)
// 返回值:
//   int: 1表示成功, 0表示失败
func (op *LibOP) GetWordResultPos(result string, index int, x, y *int) int {
	return op.ocr.GetWordResultPos(result, index, x, y)
}

// GetWordResultStr 获取识别词组的内容
// 参数:
//   result: 识别结果字符串
//   index: 词组索引
// 返回值:
//   string: 词组内容
func (op *LibOP) GetWordResultStr(result string, index int) string {
	return op.ocr.GetWordResultStr(result, index)
}

// Ocr 识别屏幕区域中匹配颜色格式的字符串
// 参数:
//   x1, y1, x2, y2: 屏幕区域坐标
//   color: 颜色格式
//   sim: 相似度 (0.1-1.0)
// 返回值:
//   string: 识别的字符串
func (op *LibOP) Ocr(x1, y1, x2, y2 int, color string, sim float64) string {
	return op.ocr.Ocr(x1, y1, x2, y2, color, sim)
}

// OcrEx 识别字符串并返回每个字符的坐标
// 参数:
//   x1, y1, x2, y2: 屏幕区域坐标
//   color: 颜色格式
//   sim: 相似度 (0.1-1.0)
// 返回值:
//   string: 带坐标的识别结果
func (op *LibOP) OcrEx(x1, y1, x2, y2 int, color string, sim float64) string {
	return op.ocr.OcrEx(x1, y1, x2, y2, color, sim)
}

// FindStr 在屏幕区域中查找字符串并返回坐标
// 参数:
//   x1, y1, x2, y2: 屏幕区域坐标
//   strs: 要查找的字符串(可以是多个字符串)
//   color: 颜色格式
//   sim: 相似度 (0.1-1.0)
//   retX: X坐标(输出)
//   retY: Y坐标(输出)
// 返回值:
//   int: 1表示成功, 0表示失败
func (op *LibOP) FindStr(x1, y1, x2, y2 int, strs, color string, sim float64, retX, retY *int) int {
	return op.ocr.FindStr(x1, y1, x2, y2, strs, color, sim, retX, retY)
}

// FindStrEx 查找所有匹配字符串的坐标
// 参数:
//   x1, y1, x2, y2: 屏幕区域坐标
//   strs: 要查找的字符串
//   color: 颜色格式
//   sim: 相似度 (0.1-1.0)
// 返回值:
//   string: 所有匹配的坐标
func (op *LibOP) FindStrEx(x1, y1, x2, y2 int, strs, color string, sim float64) string {
	return op.ocr.FindStrEx(x1, y1, x2, y2, strs, color, sim)
}

// OcrAuto 使用自动二值化识别屏幕区域中的字符串
// 参数:
//   x1, y1, x2, y2: 屏幕区域坐标
//   sim: 相似度 (0.1-1.0)
// 返回值:
//   string: 识别的字符串
func (op *LibOP) OcrAuto(x1, y1, x2, y2 int, sim float64) string {
	return op.ocr.OcrAuto(x1, y1, x2, y2, sim)
}

// OcrFromFile 从图像文件识别文本
// 参数:
//   fileName: 图像文件名
//   colorFormat: 颜色格式
//   sim: 相似度 (0.1-1.0)
// 返回值:
//   string: 识别的文本
func (op *LibOP) OcrFromFile(fileName, colorFormat string, sim float64) string {
	return op.ocr.OcrFromFile(fileName, colorFormat, sim)
}

// OcrAutoFromFile 从图像文件识别文本(不指定颜色)
// 参数:
//   fileName: 图像文件名
//   sim: 相似度 (0.1-1.0)
// 返回值:
//   string: 识别的文本
func (op *LibOP) OcrAutoFromFile(fileName string, sim float64) string {
	return op.ocr.OcrAutoFromFile(fileName, sim)
}

// FindLine 在屏幕区域中查找线条
// 参数:
//   x1, y1, x2, y2: 屏幕区域坐标
//   color: 颜色格式
//   sim: 相似度 (0.1-1.0)
// 返回值:
//   string: 找到的线条信息
func (op *LibOP) FindLine(x1, y1, x2, y2 int, color string, sim float64) string {
	return op.ocr.FindLine(x1, y1, x2, y2, color, sim)
}

// ==================== 内存操作 ====================

// WriteData 向进程内存写入数据
// 参数:
//   hwnd: 进程句柄
//   address: 内存地址
//   data: 要写入的数据
//   size: 数据大小
// 返回值:
//   int: 1表示成功, 0表示失败
func (op *LibOP) WriteData(hwnd int, address, data string, size int) int {
	// TODO: 实现内存写入
	return 0
}

// ReadData 从进程内存读取数据
// 参数:
//   hwnd: 进程句柄
//   address: 内存地址
//   size: 数据大小
// 返回值:
//   string: 读取的数据
func (op *LibOP) ReadData(hwnd int, address string, size int) string {
	// TODO: 实现内存读取
	return ""
}

// ==================== 工具函数 ====================

// parsePositions 解析位置字符串
// 参数:
//   posStr: 位置字符串, 格式: "x1,y1|x2,y2|..."
// 返回值:
//   []core.Point: 解析后的位置数组
func parsePositions(posStr string) []core.Point {
	var positions []core.Point
	parts := strings.Split(posStr, "|")
	for _, part := range parts {
		coords := strings.Split(part, ",")
		if len(coords) == 2 {
			x, errX := strconv.Atoi(strings.TrimSpace(coords[0]))
			y, errY := strconv.Atoi(strings.TrimSpace(coords[1]))
			if errX == nil && errY == nil {
				positions = append(positions, core.Point{X: x, Y: y})
			}
		}
	}
	return positions
}

// abs 返回绝对值
// 参数:
//   x: 输入值
// 返回值:
//   int: 绝对值
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// ==================== Random Utility ====================

func init() {
	rand.Seed(time.Now().UnixNano())
}
