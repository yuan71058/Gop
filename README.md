# Go-OP: Windows自动化插件 (Go版本)

[![Go](https://img.shields.io/badge/Go-1.21+-blue.svg)](https://golang.org/)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
[![Platform](https://img.shields.io/badge/platform-Windows-blue.svg)](https://microsoft.com/windows)
[![Go Reference](https://pkg.go.dev/badge/github.com/yuan71058/Gop.svg)](https://pkg.go.dev/github.com/yuan71058/Gop)
[![GitHub stars](https://img.shields.io/github/stars/yuan71058/Gop.svg)](https://github.com/yuan71058/Gop/stargazers)
[![GitHub forks](https://img.shields.io/github/forks/yuan71058/Gop.svg)](https://github.com/yuan71058/Gop/network)

## 📖 简介

**Go-OP** 是一款专为Windows设计的开源自动化插件的Go语言实现，提供屏幕读取、输入模拟和图像处理等功能。

### 两种使用方式

1. **Go第三方库** - 在Go项目中直接调用
2. **DLL动态库** - 供C/C++/Python/C#/易语言等语言调用

## 📦 安装

### 方式一：Go第三方库

```bash
go get github.com/yuan71058/Gop
```

### 方式二：下载DLL

从 [Releases](https://github.com/yuan71058/Gop/releases) 页面下载预编译的DLL文件。

## 🚀 快速开始

### Go库使用示例

```go
package main

import (
    "fmt"
    "github.com/yuan71058/Gop/libop"
    "syscall"
)

func main() {
    // 1. 创建OP实例
    op := libop.NewLibOP()
    
    // 2. 打印版本信息
    fmt.Printf("OP Version: %s\n", op.Ver())
    
    // 3. 设置工作目录
    op.SetPath("C:\\Your\\Path")
    
    // 4. 查找窗口
    hwnd := op.FindWindow("Notepad", "")
    if hwnd != 0 {
        fmt.Printf("Found window: %d\n", hwnd)
        
        // 5. 获取窗口位置
        x1, y1, x2, y2, ret := op.GetWindowRect(syscall.Handle(hwnd))
        if ret == 1 {
            fmt.Printf("Window rect: (%d, %d, %d, %d)\n", x1, y1, x2, y2)
            
            // 6. 移动鼠标到窗口中心
            centerX := (x1 + x2) / 2
            centerY := (y1 + y2) / 2
            op.MoveTo(centerX, centerY)
        }
    }
}
```

### DLL使用示例（C语言）

```c
#include <stdio.h>
#include "gop.h"

int main() {
    // 1. 创建OP实例
    int id = CreateOP();
    
    // 2. 获取版本
    char* version = Ver(id);
    printf("Version: %s\n", version);
    FreeMemory(version);
    
    // 3. 查找窗口
    int hwnd = FindWindow(id, "Notepad", "");
    if (hwnd != 0) {
        printf("Found window: %d\n", hwnd);
    }
    
    // 4. 释放实例
    ReleaseOP(id);
    return 0;
}
```

### DLL使用示例（Python）

```python
import ctypes

# 加载DLL
op = ctypes.WinDLL('./gop_amd64.dll')

# 创建实例
id = op.CreateOP()

# 获取版本
op.Ver.restype = ctypes.c_char_p
version = op.Ver(id)
print(f"Version: {version.decode()}")

# 查找窗口
hwnd = op.FindWindow(id, b"Notepad", b"")
print(f"Window: {hwnd}")

# 释放实例
op.ReleaseOP(id)
```

### DLL使用示例（C#）

```csharp
using System;
using System.Runtime.InteropServices;

class Program
{
    [DllImport("gop_amd64.dll", CallingConvention = CallingConvention.Cdecl)]
    public static extern int CreateOP();
    
    [DllImport("gop_amd64.dll", CallingConvention = CallingConvention.Cdecl)]
    public static extern void ReleaseOP(int id);
    
    [DllImport("gop_amd64.dll", CallingConvention = CallingConvention.Cdecl)]
    public static extern IntPtr Ver(int id);
    
    [DllImport("gop_amd64.dll", CallingConvention = CallingConvention.Cdecl)]
    public static extern int FindWindow(int id, string className, string title);
    
    [DllImport("gop_amd64.dll", CallingConvention = CallingConvention.Cdecl)]
    public static extern void FreeMemory(IntPtr ptr);
    
    static void Main()
    {
        int id = CreateOP();
        
        IntPtr verPtr = Ver(id);
        string version = Marshal.PtrToStringAnsi(verPtr);
        Console.WriteLine($"Version: {version}");
        FreeMemory(verPtr);
        
        int hwnd = FindWindow(id, "Notepad", "");
        Console.WriteLine($"Window: {hwnd}");
        
        ReleaseOP(id);
    }
}
```

### DLL使用示例（易语言）

```
.版本 2
.支持库 spec

.程序集 窗口程序集_启动窗口

.子程序 __启动窗口_创建完毕
.局部变量 op_id, 整数型
.局部变量 version, 文本型

op_id ＝ CreateOP ()
version ＝ Ver (op_id)
调试输出 (“版本: ” ＋ version)
FreeMemory (version)
```

## 📚 API文档

### 核心接口

| 函数名 | 说明 | 返回值 |
|--------|------|--------|
| `NewLibOP()` | 创建OP实例 | `*LibOP` |
| `Ver()` | 获取版本号 | `string` |
| `SetPath(path)` | 设置工作目录 | `int` |
| `GetPath()` | 获取工作目录 | `string` |

### 窗口管理

| 函数名 | 说明 | 参数 | 返回值 |
|--------|------|------|--------|
| `FindWindow(className, title)` | 查找窗口 | 类名, 标题 | `int` (句柄) |
| `GetWindowRect(hwnd)` | 获取窗口位置 | 句柄 | `x1,y1,x2,y2,ret` |
| `GetWindowTitle(hwnd)` | 获取窗口标题 | 句柄 | `string` |

### 输入模拟

| 函数名 | 说明 | 参数 | 返回值 |
|--------|------|------|--------|
| `MoveTo(x, y)` | 移动鼠标 | 坐标 | `int` |
| `LeftClick()` | 左键单击 | 无 | `int` |
| `RightClick()` | 右键单击 | 无 | `int` |
| `KeyPress(key)` | 按键 | 键码 | `int` |
| `SendString(str)` | 输入字符串 | 字符串 | `int` |

### 图像处理

| 函数名 | 说明 | 参数 | 返回值 |
|--------|------|------|--------|
| `Capture(x1,y1,x2,y2,file)` | 截图保存 | 区域, 文件名 | `int` |
| `FindPic(...)` | 找图 | 区域, 图片, 相似度 | `x,y,ret` |
| `FindColor(...)` | 找色 | 区域, 颜色, 相似度 | `x,y,ret` |

### OCR识别

| 函数名 | 说明 | 参数 | 返回值 |
|--------|------|------|--------|
| `Ocr(x1,y1,x2,y2,colorFormat,sim)` | 识别文字 | 区域, 格式, 相似度 | `string` |
| `SetOcrEngine(path,dllName,argv)` | 设置OCR引擎 | 路径, 参数 | `int` |

### 算法

| 函数名 | 说明 | 参数 | 返回值 |
|--------|------|------|--------|
| `AStarFindPath(...)` | A*寻路 | 地图, 障碍物, 起点, 终点 | `string` (路径) |
| `FindNearestPos(allPos,type,x,y)` | 找最近点 | 坐标列表, 类型, 目标 | `string` |

### 工具函数

| 函数名 | 说明 | 参数 | 返回值 |
|--------|------|------|--------|
| `Sleep(ms)` | 休眠 | 毫秒 | `int` |
| `Delay(ms)` | 延时 | 毫秒 | `int` |
| `Delays(min,max)` | 随机延时 | 最小, 最大 | `int` |

## 🛠️ 编译DLL

### 环境要求

- Go 1.21+
- GCC/MinGW (用于CGO)
- Windows系统

### 使用构建脚本

**Windows批处理:**
```bash
build.bat
```

**PowerShell:**
```powershell
.\build.ps1
```

**Make (Git Bash/WSL):**
```bash
make all
```

### 手动编译

```bash
# 编译x64 DLL
set GOARCH=amd64
set CGO_ENABLED=1
go build -buildmode=c-shared -o gop_amd64.dll ./dll

# 编译x86 DLL
set GOARCH=386
go build -buildmode=c-shared -o gop_386.dll ./dll
```

### 编译输出

```
build/
├── gop_amd64.dll    # x64 DLL
├── gop_amd64.h      # x64 C头文件
├── gop_386.dll      # x86 DLL
└── gop_386.h        # x86 C头文件
```

## 📁 项目结构

```
Gop/
├── core/              # 核心类型和工具函数
├── winapi/            # Windows API封装
├── background/        # 后台操作模块
├── imageproc/         # 图像处理模块
├── ocr/               # OCR模块
├── algorithm/         # 算法模块
├── libop/             # 主接口 (Go库使用)
├── dll/               # DLL导出接口 (DLL使用)
├── example/           # 示例程序
├── build/             # 编译输出
├── build.bat          # Windows构建脚本
├── build.ps1          # PowerShell构建脚本
├── Makefile           # Make构建脚本
├── go.mod             # Go模块文件
└── README.md          # 项目说明
```

## ⚙️ 配置说明

### go.mod配置

```go
module github.com/yuan71058/Gop

go 1.21
```

### 版本兼容性

| Go版本 | 状态 | 备注 |
|--------|------|------|
| 1.21+  | ✅ 支持 | 推荐使用 |
| 1.20   | ⚠️ 可能支持 | 需测试 |
| <1.20  | ❌ 不支持 | 使用泛型特性 |

## 📝 开发指南

### 添加新功能

1. 在对应模块中实现功能
2. 在 `libop/libop.go` 中添加接口
3. 在 `dll/main.go` 中添加DLL导出函数
4. 更新文档

### 代码规范

- 所有函数必须添加注释
- 遵循Go官方编码规范
- 使用驼峰命名法

## 🤝 贡献

欢迎提交Issue和Pull Request！

## 📜 许可证

MIT License

## ⚠️ 注意事项

1. **Windows平台**: 仅支持Windows系统
2. **管理员权限**: 某些功能需要管理员权限
3. **CGO依赖**: 编译DLL需要GCC/MinGW
4. **内存管理**: DLL返回的字符串需要调用 `FreeMemory` 释放

## 📞 联系方式

- **GitHub**: [@yuan71058](https://github.com/yuan71058)
- **项目地址**: https://github.com/yuan71058/Gop
- **问题反馈**: [Issues](https://github.com/yuan71058/Gop/issues)
