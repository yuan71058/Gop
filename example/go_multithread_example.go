// GOP多线程示例 - Go语言版本
// 功能: 打开3个记事本,排列窗口,使用独立GOP对象并发输入文本

package main

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/yuan71058/GOP/libop"
)

const (
	WorkerCount = 3 // 工作线程数
)

// TextWorker 文本写入工作线程
type TextWorker struct {
	ID       int
	OP       *libop.LibOP
	MainHwnd int
	EditHwnd int
	Content  string
}

// NewTextWorker 创建新的工作线程
func NewTextWorker(id int, mainHwnd, editHwnd int, content string) *TextWorker {
	return &TextWorker{
		ID:       id,
		MainHwnd: mainHwnd,
		EditHwnd: editHwnd,
		Content:  content,
	}
}

// Init 初始化GOP子对象
func (w *TextWorker) Init() bool {
	w.OP = libop.NewLibOP()
	if w.OP == nil {
		fmt.Printf("[线程%d] 创建子对象失败\n", w.ID)
		return false
	}
	fmt.Printf("[线程%d] 子对象初始化完成, 地址: %p\n", w.ID, w.OP)
	return true
}

// BindWindow 绑定窗口
func (w *TextWorker) BindWindow() bool {
	fmt.Printf("[线程%d] 准备绑定Edit控件: %d\n", w.ID, w.EditHwnd)

	// GDI模式 + Windows键盘模式绑定Edit编辑框
	ret := w.OP.BindWindow(w.EditHwnd, "gdi", "windows", "windows", 0)
	fmt.Printf("[线程%d] BindWindow返回值: %d\n", w.ID, ret)

	if ret != 1 {
		fmt.Printf("[线程%d] 绑定窗口失败: %d\n", w.ID, ret)
		return false
	}

	fmt.Printf("[线程%d] 绑定Edit控件成功: %d\n", w.ID, w.EditHwnd)
	return true
}

// WriteText 写入文字
func (w *TextWorker) WriteText() {
	fmt.Printf("[线程%d] 开始写入文字...\n", w.ID)

	// 使用SendString发送文字(绑定窗口后使用后台模式)
	ret := w.OP.SendString(w.EditHwnd, w.Content)
	if ret != 1 {
		fmt.Printf("[线程%d] SendString失败: %d\n", w.ID, ret)
		return
	}

	// 延时,让文字显示更清晰
	time.Sleep(500 * time.Millisecond)

	fmt.Printf("[线程%d] 写入完成, 共%d个字符\n", w.ID, len([]rune(w.Content)))
}

// UnbindWindow 解绑窗口
func (w *TextWorker) UnbindWindow() {
	if w.EditHwnd != 0 {
		w.OP.UnBindWindow()
		fmt.Printf("[线程%d] 解绑窗口\n", w.ID)
	}
}

// Run 运行工作线程
func (w *TextWorker) Run(wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("[线程%d] 子对象地址: %p\n", w.ID, w.OP)

	// 绑定窗口
	if !w.BindWindow() {
		return
	}

	// 写入文字
	w.WriteText()

	// 解绑窗口
	w.UnbindWindow()

	fmt.Printf("[线程%d] 完成\n", w.ID)
}

// ParseHwndList 解析窗口句柄列表字符串(逗号分隔)
func ParseHwndList(hwndList string) []int {
	if hwndList == "" {
		return nil
	}

	parts := strings.Split(hwndList, ",")
	hwnds := make([]int, 0, len(parts))

	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}
		hwnd, err := strconv.Atoi(part)
		if err != nil {
			continue
		}
		hwnds = append(hwnds, hwnd)
	}

	return hwnds
}

// ParsePidList 解析进程ID列表字符串(竖线分隔)
func ParsePidList(pidList string) []int {
	if pidList == "" {
		return nil
	}

	parts := strings.Split(pidList, "|")
	pids := make([]int, 0, len(parts))

	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}
		pid, err := strconv.Atoi(part)
		if err != nil {
			continue
		}
		pids = append(pids, pid)
	}

	return pids
}

func main() {
	fmt.Println("========== GOP 多线程写入文本示例 ==========")
	fmt.Println()

	// 第一步: 打开3个记事本
	fmt.Println("========== 第一步: 打开记事本 ==========")
	processes := make([]*exec.Cmd, WorkerCount)
	for i := 0; i < WorkerCount; i++ {
		cmd := exec.Command("notepad.exe")
		cmd.Start()
		processes[i] = cmd
		fmt.Printf("打开记事本 %d (PID: %d)\n", i+1, cmd.Process.Pid)
	}

	// 等待窗口打开
	fmt.Println("\n等待窗口打开...")
	time.Sleep(2 * time.Second)
	fmt.Println()

	// 第二步: 查找记事本窗口
	fmt.Println("========== 第二步: 查找记事本窗口 ==========")
	mainOP := libop.NewLibOP()

	// 等待窗口完全创建
	time.Sleep(1 * time.Second)

	// 使用进程枚举方式查找记事本窗口
	hwnds := make([]int, 0, WorkerCount)

	// 枚举所有notepad.exe进程
	processList := mainOP.EnumProcess("notepad.exe")
	fmt.Printf("进程列表: %s\n", processList)

	if processList == "" {
		fmt.Println("未找到notepad.exe进程!")
		return
	}

	// 解析进程ID
	pids := ParsePidList(processList)
	fmt.Printf("找到 %d 个notepad进程\n", len(pids))

	// 对每个进程,查找其窗口
	for _, pid := range pids {
		if len(hwnds) >= WorkerCount {
			break
		}

		// 通过进程ID查找窗口
		hwnd := mainOP.FindWindowByProcessId(pid, "Notepad", "")
		if hwnd > 0 {
			hwnds = append(hwnds, hwnd)
			fmt.Printf("找到记事本窗口 (PID=%d): %d\n", pid, hwnd)
		}
	}

	fmt.Printf("找到 %d 个记事本窗口\n", len(hwnds))

	if len(hwnds) < WorkerCount {
		fmt.Printf("记事本窗口数量不足, 需要 %d 个, 找到 %d 个\n", WorkerCount, len(hwnds))
		return
	}

	// 第三步: 智能排列窗口
	fmt.Println("\n========== 第三步: 智能排列窗口 ==========")

	windowWidth := 800
	windowHeight := 600
	margin := 20

	for i := 0; i < WorkerCount; i++ {
		// 计算窗口位置(网格排列)
		row := i / 2
		col := i % 2
		x := col * (windowWidth + margin)
		y := row * (windowHeight + margin)

		// 移动窗口
		ret := mainOP.MoveWindow(hwnds[i], x, y)
		if ret == 1 {
			fmt.Printf("窗口%d 移动到 (%d, %d)\n", i+1, x, y)
		} else {
			fmt.Printf("窗口%d 移动失败\n", i+1)
		}
	}

	// 等待窗口移动完成
	time.Sleep(500 * time.Millisecond)
	fmt.Println()

	// 第四步: 查找每个记事本的Edit编辑框控件
	fmt.Println("========== 第四步: 查找Edit编辑框控件 ==========")
	editHwnds := make([]int, WorkerCount)
	for i := 0; i < WorkerCount; i++ {
		editHwnd := mainOP.FindWindowEx(hwnds[i], "Edit", "")
		if editHwnd == 0 {
			fmt.Printf("找不到记事本 %d 的Edit编辑框\n", i+1)
			return
		}
		editHwnds[i] = editHwnd
		fmt.Printf("记事本 %d: 主窗口=%d, Edit控件=%d\n", i+1, hwnds[i], editHwnd)
	}

	// 第五步: 创建GOP子对象并绑定编辑框控件
	fmt.Println("\n========== 第五步: 创建GOP子对象并绑定编辑框控件 ==========")
	var wg sync.WaitGroup

	// 准备不同的文字内容
	contents := []string{
		"Hello GOP! 这是线程1输入的文本内容。",
		"Hello GOP! 这是线程2输入的文本内容。",
		"Hello GOP! 这是线程3输入的文本内容。",
	}

	// 创建工作线程(GOP子对象)
	workers := make([]*TextWorker, WorkerCount)
	for i := 0; i < WorkerCount; i++ {
		workers[i] = NewTextWorker(i+1, hwnds[i], editHwnds[i], contents[i])
		if !workers[i].Init() {
			fmt.Printf("线程%d初始化失败\n", i+1)
			return
		}
		fmt.Printf("创建GOP子对象 %d, 准备绑定Edit控件: %d\n", i+1, editHwnds[i])
	}

	// 第六步: 并发执行写入
	fmt.Println("\n========== 第六步: 开始多线程写入 ==========")
	startTime := time.Now()

	for i := 0; i < WorkerCount; i++ {
		wg.Add(1)
		go workers[i].Run(&wg)
	}

	// 等待所有线程完成
	wg.Wait()

	elapsed := time.Since(startTime)
	fmt.Println()
	fmt.Printf("========== 所有线程完成 ==========\n")
	fmt.Printf("总耗时: %v\n", elapsed)
	fmt.Println()

	// 延时5秒,让用户看到效果
	fmt.Println("等待5秒,查看效果...")
	time.Sleep(5 * time.Second)

	// 第七步: 关闭记事本窗口
	fmt.Println("\n========== 第七步: 结束进程 ==========")
	for i := 0; i < WorkerCount; i++ {
		pid := mainOP.GetWindowProcessId(hwnds[i])
		if pid > 0 {
			mainOP.TerminateProcess(pid)
			fmt.Printf("结束进程: 记事本 %d (PID: %d)\n", i+1, pid)
		}
	}

	// 第八步: 释放资源
	fmt.Println("\n========== 第八步: 释放资源 ==========")
	fmt.Println("GOP多线程示例完成")
	fmt.Println("\n========== 示例完成 ==========")
}
