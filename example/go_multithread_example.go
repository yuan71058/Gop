// GOP多线程示例 - Go语言版本
// 功能: 打开3个记事本,分别后台绑定并输入文本

package main

import (
	"fmt"
	"os/exec"
	"sync"
	"time"

	"github.com/yuan71058/GOP/libop"
)

// TaskInfo 任务信息
type TaskInfo struct {
	ID     int
	Notepad string
	Text   string
}

// NotepadTask 记事本任务
func NotepadTask(wg *sync.WaitGroup, task TaskInfo) {
	defer wg.Done()

	fmt.Printf("[任务%d] 开始执行...\n", task.ID)

	// 1. 创建独立的OP对象
	op := libop.NewLibOP()
	fmt.Printf("[任务%d] 创建OP对象成功\n", task.ID)

	// 2. 打开记事本
	fmt.Printf("[任务%d] 打开记事本...\n", task.ID)
	cmd := exec.Command("notepad.exe")
	err := cmd.Start()
	if err != nil {
		fmt.Printf("[任务%d] 启动记事本失败: %v\n", task.ID, err)
		return
	}

	// 等待记事本启动
	time.Sleep(2 * time.Second)

	// 3. 查找记事本窗口
	fmt.Printf("[任务%d] 查找记事本窗口...\n", task.ID)
	hwnd := op.FindWindow("", task.Notepad)
	if hwnd == 0 {
		fmt.Printf("[任务%d] 未找到记事本窗口!\n", task.ID)
		return
	}
	fmt.Printf("[任务%d] 找到记事本窗口, 句柄: %d\n", task.ID, hwnd)

	// 4. 后台绑定窗口(GDI模式)
	fmt.Printf("[任务%d] 后台绑定窗口(GDI模式)...\n", task.ID)
	bindRet := op.BindWindow(hwnd, "gdi", "normal", "normal", 0)
	if bindRet == 1 {
		fmt.Printf("[任务%d] 窗口绑定成功!\n", task.ID)
	} else {
		fmt.Printf("[任务%d] 窗口绑定失败!\n", task.ID)
		return
	}

	// 5. 查找Edit控件
	fmt.Printf("[任务%d] 查找Edit控件...\n", task.ID)
	editHwnd := op.FindWindowEx(hwnd, "Edit", "")
	if editHwnd == 0 {
		fmt.Printf("[任务%d] 未找到Edit控件!\n", task.ID)
		op.UnBindWindow()
		return
	}
	fmt.Printf("[任务%d] 找到Edit控件, 句柄: %d\n", task.ID, editHwnd)

	// 6. 输入文本
	fmt.Printf("[任务%d] 输入文本...\n", task.ID)
	op.SendString(editHwnd, task.Text)
	fmt.Printf("[任务%d] 已输入: %s\n", task.ID, task.Text)

	// 7. 延时3秒
	fmt.Printf("[任务%d] 延时3秒...\n", task.ID)
	op.Delay(3000)

	// 8. 解绑窗口
	fmt.Printf("[任务%d] 解绑窗口...\n", task.ID)
	op.UnBindWindow()

	// 9. 结束进程
	fmt.Printf("[任务%d] 结束记事本进程...\n", task.ID)
	pid := op.GetWindowProcessId(hwnd)
	if pid > 0 {
		op.TerminateProcess(pid)
		fmt.Printf("[任务%d] 进程已终止!\n", task.ID)
	}

	fmt.Printf("[任务%d] 任务完成!\n", task.ID)
}

func main() {
	fmt.Println("=== GOP 多线程示例 ===")

	// 定义3个任务
	tasks := []TaskInfo{
		{
			ID:     1,
			Notepad: "无标题 - 记事本",
			Text:   "Hello GOP! 这是任务1输入的文本。",
		},
		{
			ID:     2,
			Notepad: "无标题 - 记事本",
			Text:   "Hello GOP! 这是任务2输入的文本。",
		},
		{
			ID:     3,
			Notepad: "无标题 - 记事本",
			Text:   "Hello GOP! 这是任务3输入的文本。",
		},
	}

	// 使用WaitGroup等待所有任务完成
	var wg sync.WaitGroup

	// 启动3个并发任务
	for _, task := range tasks {
		wg.Add(1)
		go NotepadTask(&wg, task)
		// 每个任务之间间隔500毫秒启动
		time.Sleep(500 * time.Millisecond)
	}

	// 等待所有任务完成
	wg.Wait()

	fmt.Println("\n=== 所有任务完成 ===")
}
