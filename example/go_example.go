// GOP示例代码 - Go语言版本
// 功能: 注册插件 -> 输出版本号 -> 打开记事本 -> 后台绑定窗口 -> 输入文本 -> 延时5秒 -> 关闭窗口

package main

import (
	"fmt"
	"os/exec"
	"time"

	"github.com/yuan71058/GOP/libop"
)

func main() {
	fmt.Println("=== GOP Go 示例 ===")

	// 1. 注册插件(创建OP实例)
	fmt.Println("\n1. 注册插件...")
	op := libop.NewLibOP()
	fmt.Println("插件注册成功!")

	// 2. 输出版本号
	fmt.Println("\n2. 输出版本号...")
	version := op.Ver()
	fmt.Printf("GOP版本: %s\n", version)

	// 设置工作路径
	op.SetPath("C:\\")

	// 3. 打开记事本
	fmt.Println("\n3. 打开记事本...")
	cmd := exec.Command("notepad.exe")
	err := cmd.Start()
	if err != nil {
		fmt.Printf("启动记事本失败: %v\n", err)
		return
	}
	fmt.Println("正在启动记事本...")

	// 等待记事本启动
	time.Sleep(2 * time.Second)

	// 4. 查找记事本窗口
	fmt.Println("\n4. 查找记事本窗口...")
	hwnd := op.FindWindow("", "无标题 - 记事本")
	if hwnd == 0 {
		fmt.Println("未找到记事本窗口!")
		return
	}
	fmt.Printf("找到记事本窗口, 句柄: %d\n", hwnd)

	// 获取窗口信息
	var x1, y1, x2, y2 int
	op.GetWindowRect(hwnd, &x1, &y1, &x2, &y2)
	fmt.Printf("窗口位置: (%d, %d) - (%d, %d)\n", x1, y1, x2, y2)

	var width, height int
	op.GetClientSize(hwnd, &width, &height)
	fmt.Printf("客户区大小: %d x %d\n", width, height)

	// 5. 后台绑定窗口
	fmt.Println("\n5. 后台绑定窗口...")
	bindRet := op.BindWindow(hwnd, "normal", "normal", "normal", 0)
	if bindRet == 1 {
		fmt.Println("窗口绑定成功!")
	} else {
		fmt.Println("窗口绑定失败!")
		op.SetWindowState(hwnd, 6)
		return
	}

	// 6. 向记事本Edit控件输入文本
	fmt.Println("\n6. 向记事本Edit控件输入文本...")
	text := "Hello GOP! 这是从Go示例输入的文本。"
	// 查找记事本的Edit控件(类名为"Edit")
	editHwnd := op.FindWindowEx(hwnd, "Edit", "")
	if editHwnd == 0 {
		fmt.Println("未找到Edit控件!")
		op.SetWindowState(hwnd, 6)
		return
	}
	fmt.Printf("找到Edit控件, 句柄: %d\n", editHwnd)
	// 使用WM_SETTEXT消息直接设置Edit控件文本
	op.SendString(editHwnd, text)
	fmt.Printf("已输入: %s\n", text)

	// 7. 延时5秒
	fmt.Println("\n7. 延时5秒...")
	fmt.Println("等待中...")
	op.Delay(5000)
	fmt.Println("延时结束!")

	// 8. 解绑窗口
	fmt.Println("\n8. 解绑窗口...")
	op.UnBindWindow()
	fmt.Println("窗口已解绑!")

	// 9. 关闭记事本窗口
	fmt.Println("\n9. 关闭记事本窗口...")
	op.SetWindowState(hwnd, 6)
	fmt.Println("窗口已关闭!")

	fmt.Println("\n=== 示例完成 ===")
}
