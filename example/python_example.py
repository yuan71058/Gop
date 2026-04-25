# GOP示例代码 - Python版本
# 功能: 注册插件 -> 输出版本号 -> 打开记事本 -> 后台绑定窗口 -> 输入文本 -> 延时5秒 -> 关闭窗口

import ctypes
import os
import subprocess
import time


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
        self.dll.GetWindowRect.argtypes = [ctypes.c_int, ctypes.POINTER(ctypes.c_int),
                                           ctypes.POINTER(ctypes.c_int), ctypes.POINTER(ctypes.c_int),
                                           ctypes.POINTER(ctypes.c_int)]
        self.dll.GetWindowRect.restype = ctypes.c_int
        self.dll.GetClientSize.argtypes = [ctypes.c_int, ctypes.POINTER(ctypes.c_int),
                                           ctypes.POINTER(ctypes.c_int)]
        self.dll.GetClientSize.restype = ctypes.c_int

        # 绑定窗口
        self.dll.BindWindow.argtypes = [ctypes.c_int, ctypes.c_char_p, ctypes.c_char_p,
                                        ctypes.c_char_p, ctypes.c_int]
        self.dll.BindWindow.restype = ctypes.c_int
        self.dll.UnBindWindow.restype = ctypes.c_int

        # 键盘操作
        self.dll.KeyPressChar.argtypes = [ctypes.c_char_p]
        self.dll.KeyPressChar.restype = ctypes.c_int

        # 延迟
        self.dll.Delay.argtypes = [ctypes.c_int]
        self.dll.Delay.restype = ctypes.c_int

        # 关闭窗口
        self.dll.CloseWindow.argtypes = [ctypes.c_int]
        self.dll.CloseWindow.restype = ctypes.c_int

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

    def get_window_rect(self, hwnd):
        """获取窗口位置"""
        x1, y1, x2, y2 = ctypes.c_int(0), ctypes.c_int(0), ctypes.c_int(0), ctypes.c_int(0)
        ret = self.dll.GetWindowRect(hwnd, ctypes.byref(x1), ctypes.byref(y1),
                                     ctypes.byref(x2), ctypes.byref(y2))
        return ret, x1.value, y1.value, x2.value, y2.value

    def get_client_size(self, hwnd):
        """获取客户区大小"""
        width, height = ctypes.c_int(0), ctypes.c_int(0)
        ret = self.dll.GetClientSize(hwnd, ctypes.byref(width), ctypes.byref(height))
        return ret, width.value, height.value

    def bind_window(self, hwnd, display="normal", mouse="normal", keypad="normal", mode=0):
        """绑定窗口"""
        return self.dll.BindWindow(hwnd, display.encode(), mouse.encode(), keypad.encode(), mode)

    def unbind_window(self):
        """解绑窗口"""
        return self.dll.UnBindWindow()

    def key_press_char(self, char):
        """按键输入"""
        return self.dll.KeyPressChar(char.encode('gbk'))

    def delay(self, ms):
        """延迟"""
        return self.dll.Delay(ms)

    def close_window(self, hwnd):
        """关闭窗口"""
        return self.dll.CloseWindow(hwnd)


def main():
    """主函数"""
    print("=== GOP Python 示例 ===")

    # 获取DLL路径
    dll_path = os.path.join(os.path.dirname(__file__), "..", "dll", "GOP_x64.dll")

    if not os.path.exists(dll_path):
        print(f"错误: 找不到DLL文件 {dll_path}")
        print("请先编译GOP DLL")
        return

    # 1. 注册插件(加载DLL)
    print("\n1. 注册插件...")
    gop = GOPLib(dll_path)
    print("插件注册成功!")

    # 2. 输出版本号
    print("\n2. 输出版本号...")
    version = gop.ver()
    print(f"GOP版本: {version}")

    # 设置工作路径
    gop.set_path("C:\\")

    # 3. 打开记事本
    print("\n3. 打开记事本...")
    subprocess.Popen(["notepad.exe"])
    print("正在启动记事本...")

    # 等待记事本启动
    time.sleep(2)

    # 4. 查找记事本窗口
    print("\n4. 查找记事本窗口...")
    hwnd = gop.find_window("", "无标题 - 记事本")
    if hwnd == 0:
        print("未找到记事本窗口!")
        return
    print(f"找到记事本窗口, 句柄: {hwnd}")

    # 获取窗口信息
    ret, x1, y1, x2, y2 = gop.get_window_rect(hwnd)
    print(f"窗口位置: ({x1}, {y1}) - ({x2}, {y2})")

    ret, width, height = gop.get_client_size(hwnd)
    print(f"客户区大小: {width} x {height}")

    # 5. 后台绑定窗口
    print("\n5. 后台绑定窗口...")
    bind_ret = gop.bind_window(hwnd, "normal", "normal", "normal", 0)
    if bind_ret == 1:
        print("窗口绑定成功!")
    else:
        print("窗口绑定失败!")
        gop.close_window(hwnd)
        return

    # 6. 向记事本窗口输入一句话
    print("\n6. 向记事本输入文本...")
    text = "Hello GOP! 这是从Python示例输入的文本。"
    gop.key_press_char(text)
    print(f"已输入: {text}")

    # 7. 延时5秒
    print("\n7. 延时5秒...")
    print("等待中...")
    gop.delay(5000)
    print("延时结束!")

    # 8. 解绑窗口
    print("\n8. 解绑窗口...")
    gop.unbind_window()
    print("窗口已解绑!")

    # 9. 关闭记事本窗口
    print("\n9. 关闭记事本窗口...")
    gop.close_window(hwnd)
    print("窗口已关闭!")

    print("\n=== 示例完成 ===")


if __name__ == "__main__":
    main()
