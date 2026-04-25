# GOP多线程示例 - Python版本
# 功能: 打开3个记事本,分别后台绑定并输入文本

import ctypes
import os
import subprocess
import time
from threading import Thread


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

        # 绑定窗口
        self.dll.BindWindow.argtypes = [ctypes.c_int, ctypes.c_char_p, ctypes.c_char_p,
                                        ctypes.c_char_p, ctypes.c_int]
        self.dll.BindWindow.restype = ctypes.c_int
        self.dll.UnBindWindow.restype = ctypes.c_int

        # 文本输入
        self.dll.SendString.argtypes = [ctypes.c_int, ctypes.c_char_p]
        self.dll.SendString.restype = ctypes.c_int

        # 进程操作
        self.dll.GetWindowProcessId.argtypes = [ctypes.c_int]
        self.dll.GetWindowProcessId.restype = ctypes.c_int
        self.dll.TerminateProcess.argtypes = [ctypes.c_int]
        self.dll.TerminateProcess.restype = ctypes.c_int

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

    def bind_window(self, hwnd, display="gdi", mouse="normal", keypad="normal", mode=0):
        """绑定窗口"""
        return self.dll.BindWindow(hwnd, display.encode(), mouse.encode(), keypad.encode(), mode)

    def unbind_window(self):
        """解绑窗口"""
        return self.dll.UnBindWindow()

    def send_string(self, hwnd, text):
        """发送文本到窗口"""
        return self.dll.SendString(hwnd, text.encode('gbk'))

    def get_window_process_id(self, hwnd):
        """获取窗口进程ID"""
        return self.dll.GetWindowProcessId(hwnd)

    def terminate_process(self, pid):
        """结束进程"""
        return self.dll.TerminateProcess(pid)

    def delay(self, ms):
        """延迟"""
        return self.dll.Delay(ms)


def notepad_task(task_id, notepad_title, text):
    """记事本任务函数"""
    print(f"[任务{task_id}] 开始执行...")

    # 1. 创建独立的GOP对象
    dll_path = os.path.join(os.path.dirname(__file__), "..", "dll", "GOP_x64.dll")
    gop = GOPLib(dll_path)
    print(f"[任务{task_id}] 创建GOP对象成功")

    # 2. 打开记事本
    print(f"[任务{task_id}] 打开记事本...")
    subprocess.Popen(["notepad.exe"])
    time.sleep(2)

    # 3. 查找记事本窗口
    print(f"[任务{task_id}] 查找记事本窗口...")
    hwnd = gop.find_window("", notepad_title)
    if hwnd == 0:
        print(f"[任务{task_id}] 未找到记事本窗口!")
        return
    print(f"[任务{task_id}] 找到记事本窗口, 句柄: {hwnd}")

    # 4. 后台绑定窗口(GDI模式)
    print(f"[任务{task_id}] 后台绑定窗口(GDI模式)...")
    bind_ret = gop.bind_window(hwnd, "gdi", "normal", "normal", 0)
    if bind_ret == 1:
        print(f"[任务{task_id}] 窗口绑定成功!")
    else:
        print(f"[任务{task_id}] 窗口绑定失败!")
        return

    # 5. 查找Edit控件
    print(f"[任务{task_id}] 查找Edit控件...")
    edit_hwnd = gop.find_window_ex(hwnd, "Edit", "")
    if edit_hwnd == 0:
        print(f"[任务{task_id}] 未找到Edit控件!")
        gop.unbind_window()
        return
    print(f"[任务{task_id}] 找到Edit控件, 句柄: {edit_hwnd}")

    # 6. 输入文本
    print(f"[任务{task_id}] 输入文本...")
    gop.send_string(edit_hwnd, text)
    print(f"[任务{task_id}] 已输入: {text}")

    # 7. 延时3秒
    print(f"[任务{task_id}] 延时3秒...")
    gop.delay(3000)

    # 8. 解绑窗口
    print(f"[任务{task_id}] 解绑窗口...")
    gop.unbind_window()

    # 9. 结束进程
    print(f"[任务{task_id}] 结束记事本进程...")
    pid = gop.get_window_process_id(hwnd)
    if pid > 0:
        gop.terminate_process(pid)
        print(f"[任务{task_id}] 进程已终止!")

    print(f"[任务{task_id}] 任务完成!")


def main():
    """主函数"""
    print("=== GOP Python 多线程示例 ===")

    # 定义3个任务
    tasks = [
        {
            "id": 1,
            "title": "无标题 - 记事本",
            "text": "Hello GOP! 这是任务1输入的文本。"
        },
        {
            "id": 2,
            "title": "无标题 - 记事本",
            "text": "Hello GOP! 这是任务2输入的文本。"
        },
        {
            "id": 3,
            "title": "无标题 - 记事本",
            "text": "Hello GOP! 这是任务3输入的文本。"
        }
    ]

    # 创建并启动3个线程
    threads = []
    for task in tasks:
        thread = Thread(
            target=notepad_task,
            args=(task["id"], task["title"], task["text"])
        )
        threads.append(thread)
        thread.start()
        # 每个任务之间间隔500毫秒启动
        time.sleep(0.5)

    # 等待所有线程完成
    for thread in threads:
        thread.join()

    print("\n=== 所有任务完成 ===")


if __name__ == "__main__":
    main()
