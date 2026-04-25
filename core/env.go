// Package core 提供OP项目的全局变量和环境配置
// 该包管理全局配置，如错误消息显示模式、延迟设置等
package core

import (
	"os"
	"path/filepath"
	"sync"
)

// OpEnv OP环境管理类
// 用于管理全局配置和状态
type OpEnv struct {
	mu             sync.RWMutex // 读写锁，保护并发访问
	instance       interface{}  // 实例句柄（保留用于兼容性）
	basePath       string       // 基础路径
	opName         string       // OP名称
	showErrorMsg   int          // 错误消息显示模式: 0=关闭, 1=消息框, 2=文件, 3=标准输出
}

// 全局环境变量实例
var envInstance = &OpEnv{
	basePath:     getDefaultBasePath(),
	opName:       "op",
	showErrorMsg: 1,
}

// getDefaultBasePath 获取默认基础路径
// 返回值:
//   string: 默认基础路径（可执行文件所在目录）
func getDefaultBasePath() string {
	exe, err := os.Executable()
	if err != nil {
		// 如果获取失败，返回当前工作目录
		cwd, _ := os.Getwd()
		return cwd
	}
	return filepath.Dir(exe)
}

// GetInstance 获取环境变量实例
// 返回值:
//   *OpEnv: 环境变量实例
func GetInstance() *OpEnv {
	return envInstance
}

// SetInstance 设置实例句柄
// 参数:
//   instance: 要设置的实例句柄
func SetInstance(instance interface{}) {
	envInstance.mu.Lock()
	defer envInstance.mu.Unlock()
	envInstance.instance = instance
}

// GetBasePath 获取基础路径
// 返回值:
//   string: 基础路径
func GetBasePath() string {
	envInstance.mu.RLock()
	defer envInstance.mu.RUnlock()
	return envInstance.basePath
}

// SetBasePath 设置基础路径
// 参数:
//   path: 要设置的基础路径
func SetBasePath(path string) {
	envInstance.mu.Lock()
	defer envInstance.mu.Unlock()
	envInstance.basePath = path
}

// GetOpName 获取OP名称
// 返回值:
//   string: OP名称
func GetOpName() string {
	envInstance.mu.RLock()
	defer envInstance.mu.RUnlock()
	return envInstance.opName
}

// SetOpName 设置OP名称
// 参数:
//   name: 要设置的OP名称
func SetOpName(name string) {
	envInstance.mu.Lock()
	defer envInstance.mu.Unlock()
	envInstance.opName = name
}

// GetShowErrorMsg 获取错误消息显示模式
// 返回值:
//   int: 错误消息显示模式 (0=关闭, 1=消息框, 2=文件, 3=标准输出)
func GetShowErrorMsg() int {
	envInstance.mu.RLock()
	defer envInstance.mu.RUnlock()
	return envInstance.showErrorMsg
}

// SetShowErrorMsg 设置错误消息显示模式
// 参数:
//   mode: 错误消息显示模式 (0=关闭, 1=消息框, 2=文件, 3=标准输出)
func SetShowErrorMsg(mode int) {
	envInstance.mu.Lock()
	defer envInstance.mu.Unlock()
	envInstance.showErrorMsg = mode
}

// 全局延迟配置（用于鼠标和键盘操作）
var (
	// MouseNormalDelay 普通鼠标模式延迟（毫秒）
	MouseNormalDelay int = 30
	// MouseWindowsDelay Windows鼠标模式延迟（毫秒）
	MouseWindowsDelay int = 30
	// MouseDxDelay DirectX鼠标模式延迟（毫秒）
	MouseDxDelay int = 30
	// KeypadNormalDelay 普通键盘模式延迟（毫秒）
	KeypadNormalDelay int = 30
	// KeypadNormal2Delay 普通键盘模式2延迟（毫秒）
	KeypadNormal2Delay int = 30
	// KeypadWindowsDelay Windows键盘模式延迟（毫秒）
	KeypadWindowsDelay int = 30
	// KeypadDxDelay DirectX键盘模式延迟（毫秒）
	KeypadDxDelay int = 30
)
