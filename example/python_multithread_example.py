# GOP多线程示例 - Python版本
# 功能: 打开3个记事本,分别后台绑定并输入50行文本

import ctypes
import os
import subprocess
import time
from threading import Thread
import re


class GOPLib:
    """GOP库封装类"""

    def __init__(self, dll_path):
        """
        初始化GOP库

        参数:
            dll_path: GOP DLL文件路径
        """
        self.dll = ctypes.WinDLL(dll_path)
        self._setup_functions()

    def _setup_functions(self):
        """设置DLL函数参数和返回类型"""
        # 基础函数
        self.dll.CreateOp.restype = ctypes.c_int
        self.dll.Ver.restype = ctypes.c_char_p

        # 路径设置
        self.dll.SetPath.argtypes = [ctypes.c_char_p]
        self.dll.SetPath.restype = ctypes.c_int

        # 窗口操作
        self.dll.FindWindow.argtypes = [ctypes.c_char_p, ctypes.c_char_p]
        self.dll.FindWindow.restype = ctypes.c_int
        self.dll.FindWindowEx.argtypes = [ctypes.c_int, ctypes.c_char_p, ctypes.c_char_p]
        self.dll.FindWindowEx.restype = ctypes.c_int
        self.dll.MoveWindow.argtypes = [ctypes.c_int, ctypes.c_int, ctypes.c_int]
        self.dll.MoveWindow.restype = ctypes.c_int

        # 绑定窗口
        self.dll.BindWindow.argtypes = [ctypes.c_int, ctypes.c_char_p, ctypes.c_char_p,
                                        ctypes.c_char_p, ctypes.c_int]
        self.dll.BindWindow.restype = ctypes.c_int
        self.dll.UnBindWindow.restype = ctypes.c_int

        # 文本输入
        self.dll.SendString.argtypes = [ctypes.c_int, ctypes.c_char_p]
        self.dll.SendString.restype = ctypes.c_int
        self.dll.KeyPress.argtypes = [ctypes.c_int]
        self.dll.KeyPress.restype = ctypes.c_int

        # 进程操作
        self.dll.GetWindowProcessId.argtypes = [ctypes.c_int]
        self.dll.GetWindowProcessId.restype = ctypes.c_int
        self.dll.TerminateProcess.argtypes = [ctypes.c_int]
        self.dll.TerminateProcess.restype = ctypes.c_int
        self.dll.EnumProcess.argtypes = [ctypes.c_char_p]
        self.dll.EnumProcess.restype = ctypes.c_char_p
        self.dll.FindWindowByProcessId.argtypes = [ctypes.c_int, ctypes.c_char_p, ctypes.c_char_p]
        self.dll.FindWindowByProcessId.restype = ctypes.c_int

        # 延迟
        self.dll.Delay.argtypes = [ctypes.c_int]
        self.dll.Delay.restype = ctypes.c_int

        # 创建OP实例
        self.dll.CreateOp()

    def ver(self):
        """获取版本号"""
        return self.dll.Ver().decode('gbk')

    def set_path(self, path):
        """设置工作路径"""
        return self.dll.SetPath(path.encode('gbk'))

    def find_window(self, class_name, title):
        """查找窗口"""
        return self.dll.FindWindow(class_name.encode('gbk'), title.encode('gbk'))

    def find_window_ex(self, parent, class_name, title):
        """查找子窗口"""
        return self.dll.FindWindowEx(parent, class_name.encode(), title.encode())

    def move_window(self, hwnd, x, y):
        """移动窗口"""
        return self.dll.MoveWindow(hwnd, x, y)

    def bind_window(self, hwnd, display="gdi", mouse="normal", keypad="windows", mode=0):
        """绑定窗口"""
        return self.dll.BindWindow(hwnd, display.encode(), mouse.encode(), keypad.encode(), mode)

    def unbind_window(self):
        """解绑窗口"""
        return self.dll.UnBindWindow()

    def send_string(self, hwnd, text):
        """发送文本到窗口"""
        return self.dll.SendString(hwnd, text.encode('gbk'))

    def key_press(self, key):
        """按键"""
        return self.dll.KeyPress(key)

    def get_window_process_id(self, hwnd):
        """获取窗口进程ID"""
        return self.dll.GetWindowProcessId(hwnd)

    def terminate_process(self, pid):
        """结束进程"""
        return self.dll.TerminateProcess(pid)

    def enum_process(self, name):
        """枚举进程"""
        return self.dll.EnumProcess(name.encode()).decode('gbk')

    def find_window_by_process_id(self, pid, class_name, title):
        """通过进程ID查找窗口"""
        return self.dll.FindWindowByProcessId(pid, class_name.encode(), title.encode())

    def delay(self, ms):
        """延迟"""
        return self.dll.Delay(ms)


def generate_text_content(thread_id, line_count):
    """生成多行文本内容"""
    lines = []
    for i in range(1, line_count + 1):
        lines.append(f"线程{thread_id} - 第{i}行: GOP多线程文本输入测试内容。")
    return "\n".join(lines)


def parse_pid_list(pid_list):
    """解析进程ID列表(竖线分隔)"""
    if not pid_list:
        return []
    return [int(pid) for pid in pid_list.split('|') if pid.strip().isdigit()]


def notepad_task(task_id, hwnd, edit_hwnd, content):
    """记事本任务函数"""
    print(f"[线程{task_id}] 开始执行...")

    # 1. 创建独立的GOP对象
    dll_path = os.path.join(os.path.dirname(__file__), "..", "dll", "GOP_x64.dll")
    gop = GOPLib(dll_path)
    print(f"[线程{task_id}] 创建GOP对象成功")

    # 2. 后台绑定窗口(GDI模式 + Windows键盘模式)
    print(f"[线程{task_id}] 后台绑定窗口(Windows键盘模式)...")
    bind_ret = gop.bind_window(edit_hwnd, "gdi", "normal", "windows", 0)
    if bind_ret == 1:
        print(f"[线程{task_id}] 窗口绑定成功!")
    else:
        print(f"[线程{task_id}] 窗口绑定失败!")
        return

    # 3. 输入文本(逐行输入,每行延时200毫秒)
    print(f"[线程{task_id}] 开始写入文字...")
    
    lines = content.split('\n')
    total_chars = 0
    line_count = 0
    
    for line in lines:
        if not line:
            continue
        
        # 发送当前行文本
        gop.send_string(edit_hwnd, line)
        total_chars += len(line)
        
        # 发送回车键
        gop.key_press(13)  # VK_RETURN = 13
        total_chars += 1
        line_count += 1
        
        # 每行延时200毫秒
        gop.delay(200)
    
    print(f"[线程{task_id}] 写入完成, 共{total_chars}个字符, {line_count}行")

    # 4. 解绑窗口
    print(f"[线程{task_id}] 解绑窗口...")
    gop.unbind_window()

    print(f"[线程{task_id}] 任务完成!")


def main():
    """主函数"""
    print("=== GOP Python 多线程示例 ===")

    # 加载GOP DLL获取函数
    dll_path = os.path.join(os.path.dirname(__file__), "..", "dll", "GOP_x64.dll")
    gop = GOPLib(dll_path)
    print("插件注册成功!")
    print(f"GOP版本: {gop.ver()}")

    # 设置工作路径
    gop.set_path("C:\\")

    # 第一步: 打开3个记事本
    print("\n========== 第一步: 打开记事本 ==========")
    processes = []
    for i in range(3):
        proc = subprocess.Popen(["notepad.exe"])
        processes.append(proc)
        print(f"打开记事本 {i+1} (PID: {proc.pid})")
    
    # 等待窗口创建
    print("\n等待窗口打开...")
    time.sleep(2)

    # 第二步: 查找记事本窗口
    print("\n========== 第二步: 查找记事本窗口 ==========")
    
    process_list = gop.enum_process("notepad.exe")
    print(f"进程列表: {process_list}")
    
    pids = parse_pid_list(process_list)
    print(f"找到 {len(pids)} 个notepad进程")
    
    hwnds = []
    for pid in pids:
        if len(hwnds) >= 3:
            break
        hwnd = gop.find_window_by_process_id(pid, "Notepad", "")
        if hwnd > 0:
            hwnds.append(hwnd)
            print(f"找到记事本窗口 (PID={pid}): {hwnd}")
    
    if len(hwnds) < 3:
        print("记事本窗口数量不足!")
        return

    # 第三步: 智能排列窗口
    print("\n========== 第三步: 智能排列窗口 ==========")
    window_width = 800
    window_height = 600
    margin = 20
    
    for i in range(3):
        row = i // 2
        col = i % 2
        x = col * (window_width + margin)
        y = row * (window_height + margin)
        
        ret = gop.move_window(hwnds[i], x, y)
        if ret == 1:
            print(f"窗口{i+1} 移动到 ({x}, {y})")
    
    gop.delay(500)

    # 第四步: 查找Edit控件
    print("\n========== 第四步: 查找Edit编辑框控件 ==========")
    edit_hwnds = []
    for i in range(3):
        edit_hwnd = gop.find_window_ex(hwnds[i], "Edit", "")
        if edit_hwnd == 0:
            print(f"找不到记事本 {i+1} 的Edit编辑框")
            return
        edit_hwnds.append(edit_hwnd)
        print(f"记事本 {i+1}: 主窗口={hwnds[i]}, Edit控件={edit_hwnd}")

    # 第五步: 创建任务并启动线程
    print("\n========== 第五步: 开始多线程写入 ==========")
    
    start_time = time.time()
    
    threads = []
    for i in range(3):
        content = generate_text_content(i + 1, 50)
        thread = Thread(
            target=notepad_task,
            args=(i + 1, hwnds[i], edit_hwnds[i], content)
        )
        threads.append(thread)
        thread.start()

    # 等待所有线程完成
    for thread in threads:
        thread.join()
    
    elapsed = time.time() - start_time
    print("\n========== 所有线程完成 ==========")
    print(f"总耗时: {elapsed*1000:.1f}ms")

    # 等待5秒查看效果
    print("\n等待5秒,查看效果...")
    time.sleep(5)

    # 第六步: 结束进程
    print("\n========== 第六步: 结束进程 ==========")
    for i in range(3):
        pid = gop.get_window_process_id(hwnds[i])
        if pid > 0:
            gop.terminate_process(pid)
            print(f"结束进程: 记事本 {i+1} (PID: {pid})")

    # 第七步: 释放资源
    print("\n========== 第七步: 释放资源 ==========")
    print("GOP多线程示例完成")

    print("\n========== 示例完成 ==========")


if __name__ == "__main__":
    main()
