# GOP: Windows自动化插件 (Go版本)

[![Go](https://img.shields.io/badge/Go-1.21+-blue.svg)](https://golang.org/)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
[![Platform](https://img.shields.io/badge/platform-Windows-blue.svg)](https://microsoft.com/windows)
[![Go Reference](https://pkg.go.dev/badge/github.com/yuan71058/GOP.svg)](https://pkg.go.dev/github.com/yuan71058/GOP)
[![GitHub stars](https://img.shields.io/github/stars/yuan71058/GOP.svg)](https://github.com/yuan71058/GOP/stargazers)
[![GitHub forks](https://img.shields.io/github/forks/yuan71058/GOP.svg)](https://github.com/yuan71058/GOP/network)

## 📖 简介

**GOP** 是一款专为Windows设计的开源自动化插件的Go语言实现，提供屏幕读取、输入模拟和图像处理等功能。

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
go build -buildmode=c-shared -o bin/gop_x64.dll ./dll

# 编译x86 DLL (需要32位MinGW)
$env:GOARCH="386"
go build -buildmode=c-shared -o bin/gop_x86.dll ./dll
```

## 🚀 快速开始

### Go语言调用

```go
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
    
    // 绑定窗口
    ret := op.BindWindow(hwnd, "gdi", "windows", "windows", 0)
    if ret == 0 {
        fmt.Println("绑定窗口失败")
        return
    }
    
    // 移动鼠标并点击
    op.MoveTo(100, 100)
    op.LeftClick()
    
    // 输入文字
    op.SendString("Hello GOP!")
    
    // 解绑窗口
    op.UnBindWindow()
}
```

### Python调用 (通过DLL)

```python
import ctypes
import os

# 加载DLL
dll = ctypes.WinDLL("bin/gop_x64.dll")

# 创建OP实例
dll.CreateOp()

# 获取版本号
dll.Ver.restype = ctypes.c_char_p
version = dll.Ver().decode()
print(f"GOP Version: {version}")

# 查找窗口
dll.FindWindow.restype = ctypes.c_int
hwnd = dll.FindWindow(b"", b"记事本")
print(f"窗口句柄: {hwnd}")

# 绑定窗口
dll.BindWindow(hwnd, b"gdi", b"windows", b"windows", 0)

# 移动鼠标
dll.MoveTo(100, 100)

# 点击
dll.LeftClick()

# 输入文字
dll.SendString(b"Hello GOP from Python!")

# 解绑窗口
dll.UnBindWindow()
```

### C#调用 (通过DLL)

```csharp
using System;
using System.Runtime.InteropServices;

class Program
{
    [DllImport("gop_x64.dll", CallingConvention = CallingConvention.Cdecl)]
    public static extern int CreateOp();

    [DllImport("gop_x64.dll", CallingConvention = CallingConvention.Cdecl)]
    public static extern IntPtr Ver();

    [DllImport("gop_x64.dll", CallingConvention = CallingConvention.Cdecl)]
    public static extern int FindWindow(string className, string title);

    [DllImport("gop_x64.dll", CallingConvention = CallingConvention.Cdecl)]
    public static extern int BindWindow(int hwnd, string display, string mouse, string keypad, int mode);

    [DllImport("gop_x64.dll", CallingConvention = CallingConvention.Cdecl)]
    public static extern int MoveTo(int x, int y);

    [DllImport("gop_x64.dll", CallingConvention = CallingConvention.Cdecl)]
    public static extern int LeftClick();

    [DllImport("gop_x64.dll", CallingConvention = CallingConvention.Cdecl)]
    public static extern int SendString(string str);

    [DllImport("gop_x64.dll", CallingConvention = CallingConvention.Cdecl)]
    public static extern int UnBindWindow();

    static void Main()
    {
        CreateOp();
        
        string version = Marshal.PtrToStringAnsi(Ver());
        Console.WriteLine($"GOP Version: {version}");
        
        int hwnd = FindWindow("", "记事本");
        Console.WriteLine($"窗口句柄: {hwnd}");
        
        BindWindow(hwnd, "gdi", "windows", "windows", 0);
        MoveTo(100, 100);
        LeftClick();
        SendString("Hello GOP from C#!");
        UnBindWindow();
    }
}
```

### AutoHotkey调用 (通过DLL)

```autohotkey
; 加载DLL
hModule := DllLoad("gop_x64.dll")

; 创建OP实例
DllCall("gop_x64.dll\CreateOp")

; 获取版本号
version := DllCall("gop_x64.dll\Ver", "Str")
MsgBox GOP Version: %version%

; 查找窗口
hwnd := DllCall("gop_x64.dll\FindWindow", "Str", "", "Str", "记事本", "Int")

; 绑定窗口
DllCall("gop_x64.dll\BindWindow", "Int", hwnd, "Str", "gdi", "Str", "windows", "Str", "windows", "Int", 0)

; 移动鼠标
DllCall("gop_x64.dll\MoveTo", "Int", 100, "Int", 100)

; 点击
DllCall("gop_x64.dll\LeftClick")

; 输入文字
DllCall("gop_x64.dll\SendString", "Str", "Hello GOP from AutoHotkey!")

; 解绑窗口
DllCall("gop_x64.dll\UnBindWindow")
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
| `SendString(str)` | `str: string` | `int` | 输入字符串 |
| `FindPic(x1, y1, x2, y2, picName, deltaColor, sim, dir)` | 多个参数 | `int, int, int` | 找图 |
| `FindColor(x1, y1, x2, y2, color, sim, dir)` | 多个参数 | `int, int, int` | 找色 |
| `Ocr(x1, y1, x2, y2, color, sim)` | 多个参数 | `string` | OCR识别 |
| `GetCursorPos()` | 无 | `int, int, int` | 获取鼠标位置 |

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
│   └── main.go     # Go调用示例
├── bin/            # 编译输出
│   ├── gop_x64.dll # 64位DLL
│   └── gop_x86.dll # 32位DLL
├── go.mod          # Go模块配置
├── README.md       # 英文文档
└── 说明文档.md     # 中文文档
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
go build -buildmode=c-shared -o bin/gop_x64.dll ./dll

# 编译x86 DLL (需要32位MinGW)
$env:GOARCH="386"
go build -buildmode=c-shared -o bin/gop_x86.dll ./dll

# 运行示例
go run example/main.go
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
