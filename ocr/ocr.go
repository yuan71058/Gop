// Package ocr 提供OCR识别功能
// 该包定义了OCR引擎接口和实现
package ocr

import (
	"github.com/yuan71058/GOP/core"
)

// OcrEngine OCR引擎接口
// 定义了OCR识别的基本操作
type OcrEngine interface {
	// Init 初始化OCR引擎
	// 参数:
	//   config: 配置参数
	// 返回值:
	//   error: 初始化错误
	Init(config map[string]interface{}) error

	// Recognize 识别图像中的文字
	// 参数:
	//   imageData: 图像数据
	// 返回值:
	//   core.OcrResult: OCR识别结果
	//   error: 识别错误
	Recognize(imageData []byte) (core.OcrResult, error)

	// Close 关闭OCR引擎
	// 释放相关资源
	Close()
}

// OcrManager OCR管理器
// 管理OCR引擎的创建和使用
type OcrManager struct {
	engine OcrEngine // OCR引擎实例
}

// NewOcrManager 创建OCR管理器实例
// 返回值:
//   *OcrManager: OCR管理器实例
func NewOcrManager() *OcrManager {
	return &OcrManager{}
}

// SetEngine 设置OCR引擎
// 参数:
//   engine: OCR引擎实例
func (om *OcrManager) SetEngine(engine OcrEngine) {
	om.engine = engine
}

// Init 初始化OCR引擎
// 参数:
//   config: 配置参数
// 返回值:
//   int: 1表示成功，0表示失败
func (om *OcrManager) Init(config map[string]interface{}) int {
	if om.engine == nil {
		return 0
	}
	err := om.engine.Init(config)
	if err != nil {
		return 0
	}
	return 1
}

// Recognize 识别图像中的文字
// 参数:
//   imageData: 图像数据
// 返回值:
//   core.OcrResult: OCR识别结果
//   int: 1表示成功，0表示失败
func (om *OcrManager) Recognize(imageData []byte) (core.OcrResult, int) {
	if om.engine == nil {
		return core.OcrResult{}, 0
	}
	result, err := om.engine.Recognize(imageData)
	if err != nil {
		return core.OcrResult{}, 0
	}
	return result, 1
}

// Close 关闭OCR引擎
func (om *OcrManager) Close() {
	if om.engine != nil {
		om.engine.Close()
	}
}
