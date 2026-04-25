// 示例程序展示如何使用GOP库
package main

import (
	"fmt"
	"github.com/yuan71058/GOP/libop"
)

func main() {
	// 创建OP实例
	op := libop.NewLibOP()

	// 打印版本号
	fmt.Println("GOP Version:", op.Ver())

	// 设置工作路径
	op.SetPath("C:\\game")

	// 查找窗口
	hwnd := op.FindWindow("", "记事本")
	if hwnd == 0 {
		fmt.Println("未找到窗口")
		return
	}

	fmt.Println("找到窗口:", hwnd)

	// 绑定窗口
	ret := op.BindWindow(hwnd, "gdi", "windows", "windows", 0)
	if ret == 0 {
		fmt.Println("绑定窗口失败")
		return
	}

	fmt.Println("绑定窗口成功")

	// 移动鼠标
	op.MoveTo(100, 100)

	// 点击
	op.LeftClick()

	// 输入文字
	op.SendString("Hello GOP!")

	// 解绑窗口
	op.UnBindWindow()

	fmt.Println("操作完成")
}
