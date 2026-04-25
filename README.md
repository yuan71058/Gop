# GOP: Windows自动化插件 (Go版本)

[![Go](https://img.shields.io/badge/Go-1.21+-blue.svg)](https://golang.org/)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
[![Platform](https://img.shields.io/badge/platform-Windows-blue.svg)](https://microsoft.com/windows)
[![Go Reference](https://pkg.go.dev/badge/github.com/yuan71058/GOP.svg)](https://pkg.go.dev/github.com/yuan71058/GOP)
[![GitHub stars](https://img.shields.io/github/stars/yuan71058/GOP.svg)](https://github.com/yuan71058/GOP/stargazers)
[![GitHub forks](https://img.shields.io/github/forks/yuan71058/GOP.svg)](https://github.com/yuan71058/GOP/network)

## 📖 简介

**GOP** 是一款专为Windows设计的开源自动化插件的Go语言实现，提供屏幕读取、输入模拟和图像处理等功能。完整兼容原OP插件API，支持多种语言调用。

### ✨ 核心特性

- 🖥️ **窗口操作**: 查找窗口、绑定窗口、后台输入模拟
- 🖱️ **鼠标控制**: 移动、点击、拖拽等鼠标操作
- ⌨️ **键盘控制**: 按键、组合键、字符串输入
- 🎨 **图像处理**: 找图、找色、多点找色
- 🔍 **OCR识别**: 文字识别（支持本地和HTTP接口）
- 🗺️ **路径规划**: A*寻路算法
- 📦 **DLL导出**: 支持x86/x64架构，可被多种语言调用

## 📦 安装

### 方式一：Go Module

```bash
go get github.com/yuan71058/GOP
```

### 方式二：编译DLL

```bash
# 克隆仓库
git clone https://github.com/yuan71058/GOP.git
cd GOP

# 编译x64 DLL
go build -buildmode=c-shared -o dll/GOP_x64.dll ./dll

# 编译x86 DLL (需要32位MinGW)
$env:GOARCH="386"
go build -buildmode=c-shared -o dll/GOP_x86.dll ./dll
```

## 🚀 快速开始

### 完整示例流程

以下示例演示完整的使用流程：**注册插件 → 输出版本号 → 打开记事本 → 后台绑定窗口 → 输入文本 → 延时5秒 → 关闭窗口**

### Go语言调用

```go
package main

import (
    "fmt"
    "os/exec"
    "time"

    "github.com/yuan71058/GOP/libop"
)

func main() {
    // 1. 注册插件(创建OP实例)
    fmt.Println("1. 注册插件...")
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
    cmd.Start()
    time.Sleep(2 * time.Second)

    // 4. 查找记事本窗口
    fmt.Println("\n4. 查找记事本窗口...")
    hwnd := op.FindWindow("", "无标题 - 记事本")
    if hwnd == 0 {
        fmt.Println("未找到记事本窗口!")
        return
    }
    fmt.Printf("找到记事本窗口, 句柄: %d\n", hwnd)

    // 5. 后台绑定窗口
    fmt.Println("\n5. 后台绑定窗口...")
    bindRet := op.BindWindow(hwnd, "normal", "normal", "normal", 0)
    if bindRet == 1 {
        fmt.Println("窗口绑定成功!")
    } else {
        fmt.Println("窗口绑定失败!")
        return
    }

    // 6. 向记事本窗口输入一句话
    fmt.Println("\n6. 向记事本输入文本...")
    text := "Hello GOP! 这是从Go示例输入的文本。"
    op.KeyPressStr(text, 50)
    fmt.Printf("已输入: %s\n", text)

    // 7. 延时5秒
    fmt.Println("\n7. 延时5秒...")
    op.Delay(5000)
    fmt.Println("延时结束!")

    // 8. 解绑窗口
    fmt.Println("\n8. 解绑窗口...")
    op.UnBindWindow()
    fmt.Println("窗口已解绑!")

    // 9. 关闭记事本窗口
    fmt.Println("\n9. 关闭记事本窗口...")
    op.CloseWindow(hwnd)
    fmt.Println("窗口已关闭!")

    fmt.Println("\n=== 示例完成 ===")
}
```

### Python调用 (通过DLL)

```python
import ctypes
import os
import subprocess
import time

# 加载DLL
dll_path = "dll/GOP_x64.dll"
dll = ctypes.WinDLL(dll_path)

# 1. 注册插件
print("1. 注册插件...")
dll.CreateOp()
print("插件注册成功!")

# 2. 输出版本号
print("\n2. 输出版本号...")
dll.Ver.restype = ctypes.c_char_p
version = dll.Ver().decode('gbk')
print(f"GOP版本: {version}")

# 3. 打开记事本
print("\n3. 打开记事本...")
subprocess.Popen(["notepad.exe"])
time.sleep(2)

# 4. 查找记事本窗口
print("\n4. 查找记事本窗口...")
dll.FindWindow.argtypes = [ctypes.c_char_p, ctypes.c_char_p]
dll.FindWindow.restype = ctypes.c_int
hwnd = dll.FindWindow(b"", b"无标题 - 记事本")
if hwnd == 0:
    print("未找到记事本窗口!")
    exit()
print(f"找到记事本窗口, 句柄: {hwnd}")

# 5. 后台绑定窗口
print("\n5. 后台绑定窗口...")
dll.BindWindow.argtypes = [ctypes.c_int, ctypes.c_char_p, ctypes.c_char_p, ctypes.c_char_p, ctypes.c_int]
dll.BindWindow.restype = ctypes.c_int
bind_ret = dll.BindWindow(hwnd, b"normal", b"normal", b"normal", 0)
if bind_ret == 1:
    print("窗口绑定成功!")
else:
    print("窗口绑定失败!")
    exit()

# 6. 向记事本Edit控件输入文本
print("\n6. 向记事本Edit控件输入文本...")
dll.FindWindowEx.argtypes = [ctypes.c_int, ctypes.c_char_p, ctypes.c_char_p]
dll.FindWindowEx.restype = ctypes.c_int
dll.SendString.argtypes = [ctypes.c_int, ctypes.c_char_p]
dll.SendString.restype = ctypes.c_int
edit_hwnd = dll.FindWindowEx(hwnd, b"Edit", b"")
if edit_hwnd == 0:
    print("未找到Edit控件!")
    exit()
print(f"找到Edit控件, 句柄: {edit_hwnd}")
text = "Hello GOP! 这是从Python示例输入的文本。"
dll.SendString(edit_hwnd, text.encode('gbk'))
print(f"已输入: {text}")

# 7. 延时5秒
print("\n7. 延时5秒...")
dll.Delay.argtypes = [ctypes.c_int]
dll.Delay(5000)
print("延时结束!")

# 8. 解绑窗口
print("\n8. 解绑窗口...")
dll.UnBindWindow()
print("窗口已解绑!")

# 9. 关闭记事本窗口
print("\n9. 关闭记事本窗口...")
dll.CloseWindow.argtypes = [ctypes.c_int]
dll.CloseWindow(hwnd)
print("窗口已关闭!")

print("\n=== 示例完成 ===")
```

### C++调用 (通过DLL)

```cpp
#include <iostream>
#include <windows.h>

typedef int (*CreateOpFunc)();
typedef const char* (*VerFunc)();
typedef int (*FindWindowFunc)(const char*, const char*);
typedef int (*BindWindowFunc)(int, const char*, const char*, const char*, int);
typedef int (*FindWindowExFunc)(int, const char*, const char*);
typedef int (*SendStringFunc)(int, const char*);
typedef int (*DelayFunc)(int);
typedef int (*CloseWindowFunc)(int);

int main() {
    // 加载DLL
    HMODULE hModule = LoadLibraryA("dll/GOP_x64.dll");
    if (!hModule) {
        std::cerr << "无法加载GOP DLL" << std::endl;
        return -1;
    }

    // 获取函数地址
    CreateOpFunc CreateOp = (CreateOpFunc)GetProcAddress(hModule, "CreateOp");
    VerFunc Ver = (VerFunc)GetProcAddress(hModule, "Ver");
    FindWindowFunc FindWindow = (FindWindowFunc)GetProcAddress(hModule, "FindWindow");
    BindWindowFunc BindWindow = (BindWindowFunc)GetProcAddress(hModule, "BindWindow");
    FindWindowExFunc FindWindowEx = (FindWindowExFunc)GetProcAddress(hModule, "FindWindowEx");
    SendStringFunc SendString = (SendStringFunc)GetProcAddress(hModule, "SendString");
    DelayFunc Delay = (DelayFunc)GetProcAddress(hModule, "Delay");
    CloseWindowFunc CloseWindow = (CloseWindowFunc)GetProcAddress(hModule, "CloseWindow");

    // 1. 注册插件
    std::cout << "1. 注册插件..." << std::endl;
    CreateOp();
    std::cout << "插件注册成功!" << std::endl;

    // 2. 输出版本号
    std::cout << "\n2. 输出版本号..." << std::endl;
    std::cout << "GOP版本: " << Ver() << std::endl;

    // 3. 打开记事本
    std::cout << "\n3. 打开记事本..." << std::endl;
    ShellExecuteA(NULL, "open", "notepad.exe", NULL, NULL, SW_SHOW);
    Delay(2000);

    // 4. 查找记事本窗口
    std::cout << "\n4. 查找记事本窗口..." << std::endl;
    int hwnd = FindWindow("", "无标题 - 记事本");
    if (hwnd == 0) {
        std::cerr << "未找到记事本窗口!" << std::endl;
        return -1;
    }
    std::cout << "找到记事本窗口, 句柄: " << hwnd << std::endl;

    // 5. 后台绑定窗口
    std::cout << "\n5. 后台绑定窗口..." << std::endl;
    int bindRet = BindWindow(hwnd, "normal", "normal", "normal", 0);
    if (bindRet == 1) {
        std::cout << "窗口绑定成功!" << std::endl;
    } else {
        std::cerr << "窗口绑定失败!" << std::endl;
        return -1;
    }

    // 6. 向记事本Edit控件输入文本
    std::cout << "\n6. 向记事本Edit控件输入文本..." << std::endl;
    int editHwnd = FindWindowEx(hwnd, "Edit", "");
    if (editHwnd == 0) {
        std::cerr << "未找到Edit控件!" << std::endl;
        return -1;
    }
    std::cout << "找到Edit控件, 句柄: " << editHwnd << std::endl;
    const char* text = "Hello GOP! 这是从C++示例输入的文本。";
    SendString(editHwnd, text);
    std::cout << "已输入: " << text << std::endl;

    // 7. 延时5秒
    std::cout << "\n7. 延时5秒..." << std::endl;
    Delay(5000);
    std::cout << "延时结束!" << std::endl;

    // 8. 解绑窗口
    std::cout << "\n8. 解绑窗口..." << std::endl;
    // UnBindWindow();

    // 9. 关闭记事本窗口
    std::cout << "\n9. 关闭记事本窗口..." << std::endl;
    CloseWindow(hwnd);
    std::cout << "窗口已关闭!" << std::endl;

    // 释放DLL
    FreeLibrary(hModule);

    std::cout << "\n=== 示例完成 ===" << std::endl;
    return 0;
}
```

## 📚 API文档

### 核心API

| 函数 | 参数 | 返回值 | 说明 |
|------|------|--------|------|
| `Ver()` | 无 | `string` | 获取版本号 |
| `SetPath(path)` | `path: string` | `int` | 设置工作路径 |
| `GetPath()` | 无 | `string` | 获取工作路径 |
| `FindWindow(className, title)` | `className: string`, `title: string` | `int` | 查找窗口 |
| `BindWindow(hwnd, display, mouse, keypad, mode)` | `hwnd: int`, `display: string`, `mouse: string`, `keypad: string`, `mode: int` | `int` | 绑定窗口 |
| `UnBindWindow()` | 无 | `int` | 解绑窗口 |
| `MoveTo(x, y)` | `x: int`, `y: int` | `int` | 移动鼠标 |
| `LeftClick()` | 无 | `int` | 鼠标左键单击 |
| `RightClick()` | 无 | `int` | 鼠标右键单击 |
| `KeyPress(key)` | `key: int` | `int` | 按键 |
| `KeyPressStr(str, delay)` | `str: string`, `delay: int` | `int` | 输入字符串 |
| `FindPic(x1, y1, x2, y2, picName, deltaColor, sim, dir)` | 多个参数 | `int, int, int` | 找图 |
| `FindColor(x1, y1, x2, y2, color, sim, dir)` | 多个参数 | `int, int, int` | 找色 |
| `Ocr(x1, y1, x2, y2, color, sim)` | 多个参数 | `string` | OCR识别 |
| `GetCursorPos()` | 无 | `int, int, int` | 获取鼠标位置 |
| `Delay(ms)` | `ms: int` | `int` | 延迟(毫秒) |
| `CloseWindow(hwnd)` | `hwnd: int` | `int` | 关闭窗口 |

### 绑定模式说明

| 参数 | 可选值 | 说明 |
|------|--------|------|
| `display` | `normal`, `gdi`, `dx`, `opengl` | 显示捕获方式 |
| `mouse` | `normal`, `windows`, `dx` | 鼠标输入模拟方式 |
| `keypad` | `normal`, `windows`, `dx` | 键盘输入模拟方式 |
| `mode` | `0`, `1`, `2` | 绑定模式 |

## 🏗️ 项目结构

```
GOP/
├── core/           # 核心数据类型和工具函数
│   ├── types.go    # 数据类型定义 (Point, Rect, Color等)
│   ├── utils.go    # 工具函数 (键码映射、坐标解析等)
│   └── env.go      # 环境配置管理
├── winapi/         # Windows API封装
│   └── winapi.go   # 窗口、进程等API封装
├── background/     # 后台操作
│   └── background.go # 窗口绑定、截图、输入模拟
├── imageproc/      # 图像处理
│   └── imageproc.go  # 找图、找色等功能
├── ocr/            # OCR识别
│   └── ocr.go      # OCR引擎接口和管理器
├── algorithm/      # 算法实现
│   └── astar.go    # A*寻路算法
├── libop/          # 主库接口
│   └── libop.go    # 统一API入口
├── dll/            # DLL导出
│   └── main.go     # C兼容导出函数
├── example/        # 示例程序
│   ├── go_example.go     # Go示例
│   ├── python_example.py # Python示例
│   └── cpp_example.cpp   # C++示例
├── bin/            # 编译输出
│   ├── GOP_x64.dll # 64位DLL
│   └── GOP_x86.dll # 32位DLL
├── go.mod          # Go模块配置
├── README.md       # 项目说明文档
└── API_DOC.md      # 详细API文档
```

## 🔧 开发指南

### 环境要求

- Go 1.21+
- GCC (用于CGO/DLL编译)
- Windows操作系统

### 编译命令

```bash
# 编译所有包
go build ./...

# 编译x64 DLL
go build -buildmode=c-shared -o dll/GOP_x64.dll ./dll

# 编译x86 DLL (需要32位MinGW)
$env:GOARCH="386"
go build -buildmode=c-shared -o dll/GOP_x86.dll ./dll

# 运行示例
go run example/go_example.go
```

## 📝 许可证

本项目采用 MIT 许可证 - 详见 [LICENSE](LICENSE) 文件

## 🤝 贡献

欢迎提交Issue和Pull Request！

1. Fork 本仓库
2. 创建特性分支 (`git checkout -b feature/AmazingFeature`)
3. 提交更改 (`git commit -m 'Add some AmazingFeature'`)
4. 推送到分支 (`git push origin feature/AmazingFeature`)
5. 开启Pull Request

## 📮 联系方式

- **作者**: yuan71058
- **GitHub**: [https://github.com/yuan71058/GOP](https://github.com/yuan71058/GOP)
- **问题反馈**: [Issues](https://github.com/yuan71058/GOP/issues)

## ⭐ 支持项目

如果这个项目对你有帮助，请给它一个Star！
