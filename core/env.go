// Package core 提供环境配置和全局变量管理
// 该包管理 OP 插件的全局配置
package core

import (
	"sync"
	"sync/atomic"
	"time"
)

var (
	idCounter int32
	idMutex   sync.Mutex
)

// NextID 为每个 OP 实例生成唯一 ID
// 返回值:
//   int: 唯一 ID
func NextID() int {
	return int(atomic.AddInt32(&idCounter, 1))
}

// Env 管理环境配置和全局状态
// 存储和管理 OP 插件的全局配置信息
type Env struct {
	WorkPath        string // 工作目录路径
	BasePath        string // 插件基础目录路径
	ShowErrorMsg    int    // 错误消息显示模式
	LastError       int    // 最后错误代码
	ScreenDataMode  int    // 屏幕数据模式 (0=从上到下, 1=从下到上)
	PicCacheEnabled int    // 图片缓存启用标志
}

// NewEnv 创建新的环境配置实例
// 返回值:
//   *Env: 环境配置实例
func NewEnv() *Env {
	return &Env{
		WorkPath:        ".",
		BasePath:        ".",
		ShowErrorMsg:    0,
		LastError:       0,
		ScreenDataMode:  0,
		PicCacheEnabled: 1,
	}
}

// SetWorkPath 设置工作目录路径
// 参数:
//   path: 工作目录路径
func (e *Env) SetWorkPath(path string) {
	e.WorkPath = path
}

// GetWorkPath 获取工作目录路径
// 返回值:
//   string: 工作目录路径
func (e *Env) GetWorkPath() string {
	return e.WorkPath
}

// SetBasePath 设置插件基础目录路径
// 参数:
//   path: 插件基础目录路径
func (e *Env) SetBasePath(path string) {
	e.BasePath = path
}

// GetBasePath 获取插件基础目录路径
// 返回值:
//   string: 插件基础目录路径
func (e *Env) GetBasePath() string {
	return e.BasePath
}

// SetShowErrorMsg 设置错误消息显示模式
// 参数:
//   mode: 显示模式 (0=关闭, 1=消息框, 2=文件, 3=标准输出)
func (e *Env) SetShowErrorMsg(mode int) {
	e.ShowErrorMsg = mode
}

// GetShowErrorMsg 获取错误消息显示模式
// 返回值:
//   int: 显示模式
func (e *Env) GetShowErrorMsg() int {
	return e.ShowErrorMsg
}

// SetLastError 设置最后错误代码
// 参数:
//   code: 错误代码
func (e *Env) SetLastError(code int) {
	e.LastError = code
}

// GetLastError 获取最后错误代码
// 返回值:
//   int: 最后错误代码
func (e *Env) GetLastError() int {
	return e.LastError
}

// SetScreenDataMode 设置屏幕数据模式
// 参数:
//   mode: 屏幕数据模式 (0=从上到下, 1=从下到上)
func (e *Env) SetScreenDataMode(mode int) {
	e.ScreenDataMode = mode
}

// GetScreenDataMode 获取屏幕数据模式
// 返回值:
//   int: 屏幕数据模式
func (e *Env) GetScreenDataMode() int {
	return e.ScreenDataMode
}

// EnablePicCache 启用或禁用图片缓存
// 参数:
//   enable: 1 启用, 0 禁用
func (e *Env) EnablePicCache(enable int) {
	e.PicCacheEnabled = enable
}

// IsPicCacheEnabled 检查图片缓存是否启用
// 返回值:
//   int: 1 已启用, 0 未启用
func (e *Env) IsPicCacheEnabled() int {
	return e.PicCacheEnabled
}

// Sleep 延迟指定毫秒数（阻塞）
// 参数:
//   milliseconds: 延迟时间（毫秒）
func (e *Env) Sleep(milliseconds int) {
	time.Sleep(time.Duration(milliseconds) * time.Millisecond)
}

// Delay 延迟指定毫秒数（不阻塞 UI）
// 参数:
//   ms: 延迟时间（毫秒）
func (e *Env) Delay(ms int) {
	deadline := time.Now().Add(time.Duration(ms) * time.Millisecond)
	for time.Now().Before(deadline) {
		time.Sleep(1 * time.Millisecond)
	}
}

// Delays 在指定范围内随机延迟毫秒数（不阻塞 UI）
// 参数:
//   msMin: 最小延迟时间（毫秒）
//   msMax: 最大延迟时间（毫秒）
func (e *Env) Delays(msMin, msMax int) {
	if msMin > msMax {
		msMin, msMax = msMax, msMin
	}
	randomMs := msMin + int(time.Now().UnixNano()%(int64(msMax-msMin)+1))
	e.Delay(randomMs)
}
