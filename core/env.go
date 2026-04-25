// Package core 提供环境配置和全局变量管理
// 该包用于管理OP插件的全局配置
package core

// Env 环境配置类
// 存储和管理全局配置信息
type Env struct {
	WorkPath     string // 工作路径
	ShowErrorMsg int    // 错误消息显示模式
}

// NewEnv 创建环境配置实例
// 返回值:
//   *Env: 环境配置实例
func NewEnv() *Env {
	return &Env{
		WorkPath:     ".",
		ShowErrorMsg: 0,
	}
}

// SetWorkPath 设置工作路径
// 参数:
//   path: 工作路径
func (e *Env) SetWorkPath(path string) {
	e.WorkPath = path
}

// GetWorkPath 获取工作路径
// 返回值:
//   string: 工作路径
func (e *Env) GetWorkPath() string {
	return e.WorkPath
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
