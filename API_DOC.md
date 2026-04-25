# GOP API 完整文档

## 概述

GOP (Go OP) 是一个用 Go 语言编写的自动化操作库，完全兼容原版 OP 插件的 API 接口。本项目提供 DLL 导出接口，支持 x64 架构，可被 C++、Python、C# 等多种语言调用。

**项目地址:** <https://github.com/yuan71058/GOP>

**版本:** v1.0.0

***

## 目录

- [基础设置](#基础设置)
- [窗口操作](#窗口操作)
- [后台绑定](#后台绑定)
- [鼠标操作](#鼠标操作)
- [键盘操作](#键盘操作)
- [图像识别](#图像识别)
- [OCR 识别](#ocr-识别)
- [A* 路径查找](#a-路径查找)
- [内存操作](#内存操作)
- [工具函数](#工具函数)

***

## 基础设置

### CreateOp

创建 OP 实例。

**函数签名:**

```c
int CreateOp();
```

**参数:** 无

**返回值:**

- `1`: 创建成功

**示例:**

```c
int ret = CreateOp();
```

***

### Ver

获取库版本号。

**函数签名:**

```c
char* Ver();
```

**参数:** 无

**返回值:**

- `char*`: 版本字符串

**示例:**

```c
char* version = Ver();
// 返回: "1.0.0"
```

***

### SetPath

设置工作路径。

**函数签名:**

```c
int SetPath(char* path);
```

**参数:**

| 参数   | 类型     | 说明     |
| ------ | -------- | -------- |
| path   | char*    | 工作目录路径 |

**返回值:**

- `1`: 成功

**示例:**

```c
SetPath("C:\\game\\resource");
```

***

### GetPath

获取工作路径。

**函数签名:**

```c
char* GetPath();
```

**参数:** 无

**返回值:**

- `char*`: 当前工作路径

***

### GetBasePath

获取基础路径。

**函数签名:**

```c
char* GetBasePath();
```

**参数:** 无

**返回值:**

- `char*`: 基础路径

***

### GetID

获取实例 ID。

**函数签名:**

```c
int GetID();
```

**参数:** 无

**返回值:**

- `int`: 实例唯一标识

***

### GetLastError

获取最后错误代码。

**函数签名:**

```c
int GetLastError();
```

**参数:** 无

**返回值:**

- `int`: 错误代码

***

### OpGetLastError

获取OP库最后错误代码。

**函数签名:**

```c
int OpGetLastError();
```

**参数:** 无

**返回值:**

- `int`: 错误代码

***

### SetShowErrorMsg

设置是否显示错误消息。

**函数签名:**

```c
int SetShowErrorMsg(int show);
```

**参数:**

| 参数   | 类型  | 说明         |
| ------ | ----- | ------------ |
| show   | int   | 0=关闭, 1=消息框, 2=保存到文件, 3=输出到标准输出 |

**返回值:**

- `1`: 成功

***

### Sleep

线程休眠（毫秒）。

**函数签名:**

```c
int Sleep(int ms);
```

**参数:**

| 参数 | 类型  | 说明       |
| ---- | ----- | ---------- |
| ms   | int   | 休眠时间（毫秒） |

**返回值:**

- `1`: 成功

***

### Delay

精确延时（毫秒）。

**函数签名:**

```c
int Delay(int ms);
```

**参数:**

| 参数 | 类型  | 说明       |
| ---- | ----- | ---------- |
| ms   | int   | 延时时间（毫秒） |

**返回值:**

- `1`: 成功

***

### Delays

随机延时。

**函数签名:**

```c
int Delays(int msMin, int msMax);
```

**参数:**

| 参数    | 类型  | 说明   |
| ------- | ----- | ------ |
| msMin   | int   | 最小延时 |
| msMax   | int   | 最大延时 |

**返回值:**

- `1`: 成功

***

### EnablePicCache

启用图片缓存。

**函数签名:**

```c
int EnablePicCache(int enable);
```

**参数:**

| 参数   | 类型  | 说明         |
| ------ | ----- | ------------ |
| enable | int   | 1=启用, 0=禁用 |

**返回值:**

- `1`: 成功

***

### CapturePre

捕获上一张屏幕。

**函数签名:**

```c
int CapturePre(char* fileName);
```

**参数:**

| 参数       | 类型     | 说明    |
| ---------- | -------- | ------- |
| fileName   | char*    | 保存文件名 |

**返回值:**

- `1`: 成功

***

### SetScreenDataMode

设置屏幕数据模式。

**函数签名:**

```c
int SetScreenDataMode(int mode);
```

**参数:**

| 参数   | 类型  | 说明   |
| ------ | ----- | ------ |
| mode   | int   | 0=从上到下(默认), 1=从下到上 |

**返回值:**

- `1`: 成功

***

## 窗口操作

### EnumWindow

枚举窗口。

**函数签名:**

```c
char* EnumWindow(int parent, char* title, char* className, int filter);
```

**参数:**

| 参数        | 类型     | 说明    |
| ----------- | -------- | ------- |
| parent      | int      | 父窗口句柄 |
| title       | char*    | 窗口标题  |
| className   | char*    | 窗口类名  |
| filter      | int      | 过滤标志  |

**返回值:**

- `char*`: 窗口句柄列表，格式: "hwnd1,hwnd2,..."

***

### EnumWindowByProcess

通过进程枚举窗口。

**函数签名:**

```c
char* EnumWindowByProcess(char* processName, char* title, char* className, int filter);
```

**参数:**

| 参数          | 类型     | 说明   |
| ------------- | -------- | ------ |
| processName   | char*    | 进程名称 |
| title         | char*    | 窗口标题 |
| className     | char*    | 窗口类名 |
| filter        | int      | 过滤标志 |

**返回值:**

- `char*`: 窗口句柄列表

***

### EnumProcess

枚举进程。

**函数签名:**

```c
char* EnumProcess(char* name);
```

**参数:**

| 参数   | 类型     | 说明   |
| ------ | -------- | ------ |
| name   | char*    | 进程名称 |

**返回值:**

- `char*`: 进程 ID 列表，格式: "pid1|pid2|..."

***

### FindWindow

查找窗口。

**函数签名:**

```c
int FindWindow(char* className, char* title);
```

**参数:**

| 参数        | 类型     | 说明   |
| ----------- | -------- | ------ |
| className   | char*    | 窗口类名 |
| title       | char*    | 窗口标题 |

**返回值:**

- `int`: 窗口句柄，0 表示未找到

***

### FindWindowByProcess

通过进程查找窗口。

**函数签名:**

```c
int FindWindowByProcess(char* processName, char* className, char* title);
```

**参数:**

| 参数          | 类型     | 说明   |
| ------------- | -------- | ------ |
| processName   | char*    | 进程名称 |
| className     | char*    | 窗口类名 |
| title         | char*    | 窗口标题 |

**返回值:**

- `int`: 窗口句柄

***

### FindWindowByProcessId

通过进程 ID 查找窗口。

**函数签名:**

```c
int FindWindowByProcessId(int processId, char* className, char* title);
```

**参数:**

| 参数        | 类型     | 说明    |
| ----------- | -------- | ------- |
| processId   | int      | 进程 ID |
| className   | char*    | 窗口类名  |
| title       | char*    | 窗口标题  |

**返回值:**

- `int`: 窗口句柄

***

### FindWindowEx

查找子窗口。

**函数签名:**

```c
int FindWindowEx(int parent, char* className, char* title);
```

**参数:**

| 参数        | 类型     | 说明    |
| ----------- | -------- | ------- |
| parent      | int      | 父窗口句柄 |
| className   | char*    | 窗口类名  |
| title       | char*    | 窗口标题  |

**返回值:**

- `int`: 窗口句柄

***

### GetClientRect

获取客户区矩形。

**函数签名:**

```c
int GetClientRect(int hwnd, int* x1, int* y1, int* x2, int* y2);
```

**参数:**

| 参数   | 类型   | 说明   |
| ------ | ------ | ------ |
| hwnd   | int    | 窗口句柄 |
| x1     | int*   | 左坐标(输出) |
| y1     | int*   | 上坐标(输出) |
| x2     | int*   | 右坐标(输出) |
| y2     | int*   | 下坐标(输出) |

**返回值:**

- `1`: 成功

***

### GetClientSize

获取客户区大小。

**函数签名:**

```c
int GetClientSize(int hwnd, int* width, int* height);
```

**参数:**

| 参数   | 类型   | 说明   |
| ------ | ------ | ------ |
| hwnd   | int    | 窗口句柄 |
| width  | int*   | 宽度(输出) |
| height | int*   | 高度(输出) |

**返回值:**

- `1`: 成功

***

### GetForegroundFocus

获取前台焦点窗口。

**函数签名:**

```c
int GetForegroundFocus();
```

**参数:** 无

**返回值:**

- `int`: 窗口句柄

***

### GetForegroundWindow

获取前台窗口。

**函数签名:**

```c
int GetForegroundWindow();
```

**参数:** 无

**返回值:**

- `int`: 窗口句柄

***

### GetMousePointWindow

获取鼠标指向的窗口。

**函数签名:**

```c
int GetMousePointWindow();
```

**参数:** 无

**返回值:**

- `int`: 窗口句柄

***

### GetPointWindow

获取指定点的窗口。

**函数签名:**

```c
int GetPointWindow(int x, int y);
```

**参数:**

| 参数 | 类型  | 说明   |
| ---- | ----- | ------ |
| x    | int   | X 坐标 |
| y    | int   | Y 坐标 |

**返回值:**

- `int`: 窗口句柄

***

### GetProcessInfo

获取进程信息。

**函数签名:**

```c
char* GetProcessInfo(int pid);
```

**参数:**

| 参数  | 类型  | 说明    |
| ----- | ----- | ------- |
| pid   | int   | 进程 ID |

**返回值:**

- `char*`: 进程信息

***

### GetSpecialWindow

获取特殊窗口。

**函数签名:**

```c
int GetSpecialWindow(int flag);
```

**参数:**

| 参数   | 类型  | 说明 |
| ------ | ----- | ---- |
| flag   | int   | 0=桌面, 1=任务栏 |

**返回值:**

- `int`: 窗口句柄

***

### GetWindow

获取窗口。

**函数签名:**

```c
int GetWindow(int hwnd, int flag);
```

**参数:**

| 参数   | 类型  | 说明   |
| ------ | ----- | ------ |
| hwnd   | int   | 窗口句柄 |
| flag   | int   | 0=父窗口, 1=第一个子窗口, 2=下一个兄弟窗口, 3=上一个兄弟窗口 |

**返回值:**

- `int`: 窗口句柄

***

### GetWindowClass

获取窗口类名。

**函数签名:**

```c
char* GetWindowClass(int hwnd);
```

**参数:**

| 参数   | 类型  | 说明   |
| ------ | ----- | ------ |
| hwnd   | int   | 窗口句柄 |

**返回值:**

- `char*`: 类名

***

### GetWindowProcessId

获取窗口进程 ID。

**函数签名:**

```c
int GetWindowProcessId(int hwnd);
```

**参数:**

| 参数   | 类型  | 说明   |
| ------ | ----- | ------ |
| hwnd   | int   | 窗口句柄 |

**返回值:**

- `int`: 进程 ID

***

### GetWindowProcessPath

获取窗口进程路径。

**函数签名:**

```c
char* GetWindowProcessPath(int hwnd);
```

**参数:**

| 参数   | 类型  | 说明   |
| ------ | ----- | ------ |
| hwnd   | int   | 窗口句柄 |

**返回值:**

- `char*`: 进程路径

***

### GetWindowRect

获取窗口矩形。

**函数签名:**

```c
int GetWindowRect(int hwnd, int* x1, int* y1, int* x2, int* y2);
```

**参数:**

| 参数   | 类型   | 说明   |
| ------ | ------ | ------ |
| hwnd   | int    | 窗口句柄 |
| x1     | int*   | 左坐标(输出) |
| y1     | int*   | 上坐标(输出) |
| x2     | int*   | 右坐标(输出) |
| y2     | int*   | 下坐标(输出) |

**返回值:**

- `1`: 成功

***

### GetWindowState

获取窗口状态。

**函数签名:**

```c
int GetWindowState(int hwnd, int flag);
```

**参数:**

| 参数   | 类型  | 说明   |
| ------ | ----- | ------ |
| hwnd   | int   | 窗口句柄 |
| flag   | int   | 0=可见, 1=活动, 2=最小化, 3=最大化, 4=启用 |

**返回值:**

- `int`: 状态值

***

### GetWindowTitle

获取窗口标题。

**函数签名:**

```c
char* GetWindowTitle(int hwnd);
```

**参数:**

| 参数   | 类型  | 说明   |
| ------ | ----- | ------ |
| hwnd   | int   | 窗口句柄 |

**返回值:**

- `char*`: 窗口标题

***

### MoveWindow

移动窗口。

**函数签名:**

```c
int MoveWindow(int hwnd, int x, int y);
```

**参数:**

| 参数   | 类型  | 说明   |
| ------ | ----- | ------ |
| hwnd   | int   | 窗口句柄 |
| x      | int   | X 坐标 |
| y      | int   | Y 坐标 |

**返回值:**

- `1`: 成功

***

### ClientToScreen

客户区坐标转屏幕坐标。

**函数签名:**

```c
int ClientToScreen(int hwnd, int* x, int* y);
```

**参数:**

| 参数   | 类型   | 说明   |
| ------ | ------ | ------ |
| hwnd   | int    | 窗口句柄 |
| x      | int*   | X坐标(输入/输出) |
| y      | int*   | Y坐标(输入/输出) |

**返回值:**

- `1`: 成功

***

### ScreenToClient

屏幕坐标转客户区坐标。

**函数签名:**

```c
int ScreenToClient(int hwnd, int* x, int* y);
```

**参数:**

| 参数   | 类型   | 说明   |
| ------ | ------ | ------ |
| hwnd   | int    | 窗口句柄 |
| x      | int*   | X坐标(输入/输出) |
| y      | int*   | Y坐标(输入/输出) |

**返回值:**

- `1`: 成功

***

### SendPaste

发送粘贴。

**函数签名:**

```c
int SendPaste(int hwnd);
```

**参数:**

| 参数   | 类型  | 说明   |
| ------ | ----- | ------ |
| hwnd   | int   | 窗口句柄 |

**返回值:**

- `1`: 成功

***

### SetClientSize

设置客户区大小。

**函数签名:**

```c
int SetClientSize(int hwnd, int width, int height);
```

**参数:**

| 参数     | 类型  | 说明   |
| -------- | ----- | ------ |
| hwnd     | int   | 窗口句柄 |
| width    | int   | 宽度   |
| height   | int   | 高度   |

**返回值:**

- `1`: 成功

***

### SetWindowState

设置窗口状态。

**函数签名:**

```c
int SetWindowState(int hwnd, int flag);
```

**参数:**

| 参数   | 类型  | 说明                                    |
| ------ | ----- | --------------------------------------- |
| hwnd   | int   | 窗口句柄                                  |
| flag   | int   | 0=显示, 1=隐藏, 2=最小化, 3=最大化, 4=还原, 5=激活, 6=关闭 |

**返回值:**

- `1`: 成功

***

### SetWindowSize

设置窗口大小。

**函数签名:**

```c
int SetWindowSize(int hwnd, int width, int height);
```

**参数:**

| 参数     | 类型  | 说明   |
| -------- | ----- | ------ |
| hwnd     | int   | 窗口句柄 |
| width    | int   | 宽度   |
| height   | int   | 高度   |

**返回值:**

- `1`: 成功

***

### SetWindowText

设置窗口标题。

**函数签名:**

```c
int SetWindowText(int hwnd, char* title);
```

**参数:**

| 参数    | 类型     | 说明   |
| ------- | -------- | ------ |
| hwnd    | int      | 窗口句柄 |
| title   | char*    | 窗口标题 |

**返回值:**

- `1`: 成功

***

### SetWindowTransparent

设置窗口透明度。

**函数签名:**

```c
int SetWindowTransparent(int hwnd, int trans);
```

**参数:**

| 参数    | 类型  | 说明          |
| ------- | ----- | ------------- |
| hwnd    | int   | 窗口句柄        |
| trans   | int   | 透明度 (0-255)  |

**返回值:**

- `1`: 成功

***

### SendString

发送字符串。

**函数签名:**

```c
int SendString(int hwnd, char* str);
```

**参数:**

| 参数   | 类型     | 说明   |
| ------ | -------- | ------ |
| hwnd   | int      | 窗口句柄 |
| str    | char*    | 字符串  |

**返回值:**

- `1`: 成功

***

### SendStringIme

使用 IME 发送字符串。

**函数签名:**

```c
int SendStringIme(int hwnd, char* str);
```

**参数:**

| 参数   | 类型     | 说明   |
| ------ | -------- | ------ |
| hwnd   | int      | 窗口句柄 |
| str    | char*    | 字符串  |

**返回值:**

- `1`: 成功

***

### RunApp

运行应用。

**函数签名:**

```c
int RunApp(char* cmdline, int mode);
```

**参数:**

| 参数     | 类型     | 说明   |
| -------- | -------- | ------ |
| cmdline  | char*    | 命令行  |
| mode     | int      | 0=普通, 1=显示, 2=隐藏 |

**返回值:**

- `1`: 成功

***

### WinExec

执行 Windows 命令。

**函数签名:**

```c
int WinExec(char* cmdline, int cmdShow);
```

**参数:**

| 参数      | 类型     | 说明   |
| --------- | -------- | ------ |
| cmdline   | char*    | 命令行  |
| cmdShow   | int      | 显示模式 |

**返回值:**

- `1`: 成功

***

### GetCmdStr

获取命令输出字符串。

**函数签名:**

```c
char* GetCmdStr(char* cmd, int timeOut);
```

**参数:**

| 参数      | 类型     | 说明       |
| --------- | -------- | ---------- |
| cmd       | char*    | 命令       |
| timeOut   | int      | 超时时间（毫秒） |

**返回值:**

- `char*`: 命令输出

***

### SetClipboard

设置剪贴板内容。

**函数签名:**

```c
int SetClipboard(char* str);
```

**参数:**

| 参数  | 类型     | 说明    |
| ----- | -------- | ------- |
| str   | char*    | 剪贴板内容 |

**返回值:**

- `1`: 成功

***

### GetClipboard

获取剪贴板内容。

**函数签名:**

```c
char* GetClipboard();
```

**参数:** 无

**返回值:**

- `char*`: 剪贴板内容

***

### InjectDll

注入 DLL。

**函数签名:**

```c
int InjectDll(char* processName, char* dllName);
```

**参数:**

| 参数          | 类型     | 说明     |
| ------------- | -------- | -------- |
| processName   | char*    | 进程名称   |
| dllName       | char*    | DLL 路径   |

**返回值:**

- `1`: 成功

***

### TerminateProcess

结束进程。

**函数签名:**

```c
int TerminateProcess(int pid);
```

**参数:**

| 参数  | 类型  | 说明    |
| ----- | ----- | ------- |
| pid   | int   | 进程 ID |

**返回值:**

- `1`: 成功

***

## 后台绑定

### BindWindow

绑定窗口。

**函数签名:**

```c
int BindWindow(int hwnd, char* display, char* mouse, char* keypad, int mode);
```

**参数:**

| 参数      | 类型     | 说明                                            |
| --------- | -------- | ----------------------------------------------- |
| hwnd      | int      | 窗口句柄                                          |
| display   | char*    | 显示模式: "normal", "gdi", "dx", "dx2", "opengl"  |
| mouse     | char*    | 鼠标模式: "normal", "windows", "windows3", "dx"   |
| keypad    | char*    | 键盘模式: "normal", "windows", "dx"               |
| mode      | int      | 绑定模式: 0=普通, 1=后台                          |

**返回值:**

- `1`: 成功
- `0`: 失败

**示例:**

```c
// GDI模式绑定，使用Windows键盘模式
BindWindow(hwnd, "gdi", "normal", "windows", 0);
```

***

### UnBindWindow

解绑窗口。

**函数签名:**

```c
int UnBindWindow();
```

**参数:** 无

**返回值:**

- `1`: 成功

***

### GetBindWindow

获取绑定窗口。

**函数签名:**

```c
int GetBindWindow();
```

**参数:** 无

**返回值:**

- `int`: 窗口句柄，未绑定返回0

***

### IsBind

是否已绑定。

**函数签名:**

```c
int IsBind();
```

**参数:** 无

**返回值:**

- `1`: 已绑定
- `0`: 未绑定

***

## 鼠标操作

### GetCursorPos

获取鼠标位置。

**函数签名:**

```c
int GetCursorPos(int* x, int* y);
```

**参数:**

| 参数 | 类型   | 说明   |
| ---- | ------ | ------ |
| x    | int*   | X坐标(输出) |
| y    | int*   | Y坐标(输出) |

**返回值:**

- `1`: 成功

***

### MoveTo

移动鼠标。

**函数签名:**

```c
int MoveTo(int x, int y);
```

**参数:**

| 参数 | 类型  | 说明   |
| ---- | ----- | ------ |
| x    | int   | X 坐标 |
| y    | int   | Y 坐标 |

**返回值:**

- `1`: 成功

***

### MoveR

相对移动鼠标。

**函数签名:**

```c
int MoveR(int x, int y);
```

**参数:**

| 参数 | 类型  | 说明      |
| ---- | ----- | --------- |
| x    | int   | 相对 X 偏移 |
| y    | int   | 相对 Y 偏移 |

**返回值:**

- `1`: 成功

***

### MoveToEx

移动鼠标到范围内。

**函数签名:**

```c
int MoveToEx(int x, int y, int w, int h);
```

**参数:**

| 参数 | 类型  | 说明   |
| ---- | ----- | ------ |
| x    | int   | X 坐标 |
| y    | int   | Y 坐标 |
| w    | int   | 宽度范围 |
| h    | int   | 高度范围 |

**返回值:**

- `1`: 成功

***

### LeftClick

左键单击。

**函数签名:**

```c
int LeftClick();
```

**参数:** 无

**返回值:**

- `1`: 成功

***

### LeftDoubleClick

左键双击。

**函数签名:**

```c
int LeftDoubleClick();
```

**参数:** 无

**返回值:**

- `1`: 成功

***

### LeftDown

左键按下。

**函数签名:**

```c
int LeftDown();
```

**参数:** 无

**返回值:**

- `1`: 成功

***

### LeftUp

左键释放。

**函数签名:**

```c
int LeftUp();
```

**参数:** 无

**返回值:**

- `1`: 成功

***

### MiddleClick

中键单击。

**函数签名:**

```c
int MiddleClick();
```

**参数:** 无

**返回值:**

- `1`: 成功

***

### MiddleDown

中键按下。

**函数签名:**

```c
int MiddleDown();
```

**参数:** 无

**返回值:**

- `1`: 成功

***

### MiddleUp

中键释放。

**函数签名:**

```c
int MiddleUp();
```

**参数:** 无

**返回值:**

- `1`: 成功

***

### RightClick

右键单击。

**函数签名:**

```c
int RightClick();
```

**参数:** 无

**返回值:**

- `1`: 成功

***

### RightDown

右键按下。

**函数签名:**

```c
int RightDown();
```

**参数:** 无

**返回值:**

- `1`: 成功

***

### RightUp

右键释放。

**函数签名:**

```c
int RightUp();
```

**参数:** 无

**返回值:**

- `1`: 成功

***

### WheelDown

滚轮向下。

**函数签名:**

```c
int WheelDown();
```

**参数:** 无

**返回值:**

- `1`: 成功

***

### WheelUp

滚轮向上。

**函数签名:**

```c
int WheelUp();
```

**参数:** 无

**返回值:**

- `1`: 成功

***

### SetMouseDelay

设置鼠标延迟。

**函数签名:**

```c
int SetMouseDelay(char* type, int delay);
```

**参数:**

| 参数    | 类型     | 说明            |
| ------- | -------- | --------------- |
| type    | char*    | 类型 ("normal")  |
| delay   | int      | 延迟时间（毫秒）   |

**返回值:**

- `1`: 成功

***

## 键盘操作

### GetKeyState

获取按键状态。

**函数签名:**

```c
int GetKeyState(int vkCode);
```

**参数:**

| 参数     | 类型  | 说明   |
| -------- | ----- | ------ |
| vkCode   | int   | 虚拟键码 |

**返回值:**

- `1`: 按下
- `0`: 释放

***

### KeyDown

按键按下。

**函数签名:**

```c
int KeyDown(int vkCode);
```

**参数:**

| 参数     | 类型  | 说明   |
| -------- | ----- | ------ |
| vkCode   | int   | 虚拟键码 |

**返回值:**

- `1`: 成功

***

### KeyDownChar

字符按键按下。

**函数签名:**

```c
int KeyDownChar(char* key);
```

**参数:**

| 参数  | 类型     | 说明                   |
| ----- | -------- | ---------------------- |
| key   | char*    | 字符（如 "a", "1", "F1"） |

**返回值:**

- `1`: 成功

***

### KeyUp

按键释放。

**函数签名:**

```c
int KeyUp(int vkCode);
```

**参数:**

| 参数     | 类型  | 说明   |
| -------- | ----- | ------ |
| vkCode   | int   | 虚拟键码 |

**返回值:**

- `1`: 成功

***

### KeyUpChar

字符按键释放。

**函数签名:**

```c
int KeyUpChar(char* key);
```

**参数:**

| 参数  | 类型     | 说明 |
| ----- | -------- | ---- |
| key   | char*    | 字符 |

**返回值:**

- `1`: 成功

***

### WaitKey

等待按键。

**函数签名:**

```c
int WaitKey(int vkCode, int timeOut);
```

**参数:**

| 参数      | 类型  | 说明       |
| --------- | ----- | ---------- |
| vkCode    | int   | 虚拟键码     |
| timeOut   | int   | 超时时间（毫秒） |

**返回值:**

- `1`: 按键被按下
- `0`: 超时

***

### KeyPress

按键。

**函数签名:**

```c
int KeyPress(int vkCode);
```

**参数:**

| 参数     | 类型  | 说明   |
| -------- | ----- | ------ |
| vkCode   | int   | 虚拟键码 |

**返回值:**

- `1`: 成功

***

### KeyPressChar

字符按键。

**函数签名:**

```c
int KeyPressChar(char* key);
```

**参数:**

| 参数  | 类型     | 说明 |
| ----- | -------- | ---- |
| key   | char*    | 字符 |

**返回值:**

- `1`: 成功

***

### SetKeypadDelay

设置键盘延迟。

**函数签名:**

```c
int SetKeypadDelay(char* type, int delay);
```

**参数:**

| 参数    | 类型     | 说明            |
| ------- | -------- | --------------- |
| type    | char*    | 类型 ("normal")  |
| delay   | int      | 延迟时间（毫秒）   |

**返回值:**

- `1`: 成功

***

### KeyPressStr

按键字符串。

**函数签名:**

```c
int KeyPressStr(char* keyStr, int delay);
```

**参数:**

| 参数     | 类型     | 说明       |
| -------- | -------- | ---------- |
| keyStr   | char*    | 按键序列     |
| delay    | int      | 按键间隔（毫秒） |

**返回值:**

- `1`: 成功

***

## 图像识别

### Capture

截图。

**函数签名:**

```c
int Capture(int x1, int y1, int x2, int y2, char* fileName);
```

**参数:**

| 参数       | 类型     | 说明    |
| ---------- | -------- | ------- |
| x1         | int      | 左上角 X |
| y1         | int      | 左上角 Y |
| x2         | int      | 右下角 X |
| y2         | int      | 右下角 Y |
| fileName   | char*    | 保存文件名 |

**返回值:**

- `1`: 成功

***

### GetColor

获取颜色。

**函数签名:**

```c
char* GetColor(int x, int y);
```

**参数:**

| 参数 | 类型  | 说明   |
| ---- | ----- | ------ |
| x    | int   | X 坐标 |
| y    | int   | Y 坐标 |

**返回值:**

- `char*`: 颜色值，格式: "RRGGBB"

***

### CmpColor

比较颜色。

**函数签名:**

```c
int CmpColor(int x, int y, char* color, double sim);
```

**参数:**

| 参数    | 类型     | 说明            |
| ------- | -------- | --------------- |
| x       | int      | X 坐标          |
| y       | int      | Y 坐标          |
| color   | char*    | 颜色值           |
| sim     | double   | 相似度 (0.1-1.0) |

**返回值:**

- `1`: 匹配
- `0`: 不匹配

***

### FindColor

查找颜色。

**函数签名:**

```c
int FindColor(int x1, int y1, int x2, int y2, char* color, double sim, int dir, int* x, int* y);
```

**参数:**

| 参数    | 类型     | 说明      |
| ------- | -------- | --------- |
| x1      | int      | 左上角 X   |
| y1      | int      | 左上角 Y   |
| x2      | int      | 右下角 X   |
| y2      | int      | 右下角 Y   |
| color   | char*    | 颜色值     |
| sim     | double   | 相似度     |
| dir     | int      | 搜索方向    |
| x       | int*     | 输出 X 坐标 |
| y       | int*     | 输出 Y 坐标 |

**返回值:**

- `1`: 找到
- `0`: 未找到

***

### FindColorEx

查找颜色（扩展）。

**函数签名:**

```c
char* FindColorEx(int x1, int y1, int x2, int y2, char* color, double sim, int dir);
```

**参数:**

| 参数    | 类型     | 说明    |
| ------- | -------- | ------- |
| x1      | int      | 左上角 X |
| y1      | int      | 左上角 Y |
| x2      | int      | 右下角 X |
| y2      | int      | 右下角 Y |
| color   | char*    | 颜色值   |
| sim     | double   | 相似度   |
| dir     | int      | 搜索方向  |

**返回值:**

- `char*`: 所有找到的位置，格式: "x1,y1|x2,y2|..."

***

### GetColorNum

获取颜色数量。

**函数签名:**

```c
int GetColorNum(int x1, int y1, int x2, int y2, char* color, double sim);
```

**参数:**

| 参数    | 类型     | 说明    |
| ------- | -------- | ------- |
| x1      | int      | 左上角 X |
| y1      | int      | 左上角 Y |
| x2      | int      | 右下角 X |
| y2      | int      | 右下角 Y |
| color   | char*    | 颜色值   |
| sim     | double   | 相似度   |

**返回值:**

- `int`: 颜色点数量

***

### FindMultiColor

查找多点颜色。

**函数签名:**

```c
int FindMultiColor(int x1, int y1, int x2, int y2, char* firstColor, char* offsetColor, double sim, int dir, int* x, int* y);
```

**参数:**

| 参数          | 类型     | 说明      |
| ------------- | -------- | --------- |
| x1            | int      | 左上角 X   |
| y1            | int      | 左上角 Y   |
| x2            | int      | 右下角 X   |
| y2            | int      | 右下角 Y   |
| firstColor    | char*    | 第一个颜色   |
| offsetColor   | char*    | 偏移颜色    |
| sim           | double   | 相似度     |
| dir           | int      | 搜索方向    |
| x             | int*     | 输出 X 坐标 |
| y             | int*     | 输出 Y 坐标 |

**返回值:**

- `1`: 找到
- `0`: 未找到

***

### FindMultiColorEx

查找多点颜色（扩展）。

**函数签名:**

```c
char* FindMultiColorEx(int x1, int y1, int x2, int y2, char* firstColor, char* offsetColor, double sim, int dir);
```

**参数:** 同上

**返回值:**

- `char*`: 所有找到的位置

***

### FindPic

查找图片。

**函数签名:**

```c
int FindPic(int x1, int y1, int x2, int y2, char* files, char* deltaColor, double sim, int dir, int* x, int* y);
```

**参数:**

| 参数         | 类型     | 说明      |
| ------------ | -------- | --------- |
| x1           | int      | 左上角 X   |
| y1           | int      | 左上角 Y   |
| x2           | int      | 右下角 X   |
| y2           | int      | 右下角 Y   |
| files        | char*    | 图片名称    |
| deltaColor   | char*    | 偏色      |
| sim          | double   | 相似度     |
| dir          | int      | 搜索方向    |
| x            | int*     | 输出 X 坐标 |
| y            | int*     | 输出 Y 坐标 |

**返回值:**

- `1`: 找到
- `0`: 未找到

***

### FindPicEx

查找图片（扩展）。

**函数签名:**

```c
char* FindPicEx(int x1, int y1, int x2, int y2, char* files, char* deltaColor, double sim, int dir);
```

**参数:** 同上

**返回值:**

- `char*`: 所有找到的位置

***

### FindPicExS

查找图片（扩展字符串）。

**函数签名:**

```c
char* FindPicExS(int x1, int y1, int x2, int y2, char* files, char* deltaColor, double sim, int dir);
```

**参数:** 同上

**返回值:**

- `char*`: 所有找到的位置，格式: "file1,x,y|file2,x,y|..."

***

### FindColorBlock

查找色块。

**函数签名:**

```c
int FindColorBlock(int x1, int y1, int x2, int y2, char* color, double sim, int count, int height, int width, int* x, int* y);
```

**参数:**

| 参数     | 类型     | 说明      |
| -------- | -------- | --------- |
| x1       | int      | 左上角 X   |
| y1       | int      | 左上角 Y   |
| x2       | int      | 右下角 X   |
| y2       | int      | 右下角 Y   |
| color    | char*    | 颜色值     |
| sim      | double   | 相似度     |
| count    | int      | 最小点数    |
| height   | int      | 最小高度    |
| width    | int      | 最小宽度    |
| x        | int*     | 输出 X 坐标 |
| y        | int*     | 输出 Y 坐标 |

**返回值:**

- `1`: 找到
- `0`: 未找到

***

### FindColorBlockEx

查找色块（扩展）。

**函数签名:**

```c
char* FindColorBlockEx(int x1, int y1, int x2, int y2, char* color, double sim, int count, int height, int width);
```

**参数:** 同上（无输出坐标）

**返回值:**

- `char*`: 所有找到的位置

***

### SetDisplayInput

设置显示输入。

**函数签名:**

```c
int SetDisplayInput(char* mode);
```

**参数:**

| 参数   | 类型     | 说明   |
| ------ | -------- | ------ |
| mode   | char*    | 输入模式 |

**返回值:**

- `1`: 成功

***

### LoadPic

加载图像。

**函数签名:**

```c
int LoadPic(char* fileName);
```

**参数:**

| 参数       | 类型     | 说明    |
| ---------- | -------- | ------- |
| fileName   | char*    | 图像文件名 |

**返回值:**

- `1`: 成功

***

### FreePic

释放图像。

**函数签名:**

```c
int FreePic(char* fileName);
```

**参数:**

| 参数       | 类型     | 说明    |
| ---------- | -------- | ------- |
| fileName   | char*    | 图像文件名 |

**返回值:**

- `1`: 成功

***

### LoadMemPic

从内存加载图像。

**函数签名:**

```c
int LoadMemPic(char* fileName, void* data, int size);
```

**参数:**

| 参数       | 类型     | 说明    |
| ---------- | -------- | ------- |
| fileName   | char*    | 图像名   |
| data       | void*    | 图像数据  |
| size       | int      | 数据大小  |

**返回值:**

- `1`: 成功

***

### GetPicSize

获取图像大小。

**函数签名:**

```c
int GetPicSize(char* picName, int* width, int* height);
```

**参数:**

| 参数       | 类型     | 说明    |
| ---------- | -------- | ------- |
| picName    | char*    | 图像名   |
| width      | int*     | 宽度(输出) |
| height     | int*     | 高度(输出) |

**返回值:**

- `1`: 成功

***

### GetScreenData

获取屏幕数据。

**函数签名:**

```c
int GetScreenData(int x1, int y1, int x2, int y2, void** data);
```

**参数:**

| 参数   | 类型     | 说明    |
| ------ | -------- | ------- |
| x1     | int      | 左上角 X |
| y1     | int      | 左上角 Y |
| x2     | int      | 右下角 X |
| y2     | int      | 右下角 Y |
| data   | void**   | 数据指针(输出) |

**返回值:**

- `1`: 成功

***

### GetScreenDataBmp

获取BMP格式屏幕数据。

**函数签名:**

```c
int GetScreenDataBmp(int x1, int y1, int x2, int y2, void** data, int* size);
```

**参数:**

| 参数   | 类型     | 说明    |
| ------ | -------- | ------- |
| x1     | int      | 左上角 X |
| y1     | int      | 左上角 Y |
| x2     | int      | 右下角 X |
| y2     | int      | 右下角 Y |
| data   | void**   | 数据指针(输出) |
| size   | int*     | 数据大小(输出) |

**返回值:**

- `1`: 成功

***

### GetScreenFrameInfo

获取屏幕帧信息。

**函数签名:**

```c
int GetScreenFrameInfo(int* frameID, int* frameTime);
```

**参数:**

| 参数       | 类型     | 说明    |
| ---------- | -------- | ------- |
| frameID    | int*     | 帧ID(输出) |
| frameTime  | int*     | 帧时间(输出) |

**返回值:**

- `1`: 成功

***

### MatchPicName

匹配图像名。

**函数签名:**

```c
char* MatchPicName(char* picName);
```

**参数:**

| 参数       | 类型     | 说明    |
| ---------- | -------- | ------- |
| picName    | char*    | 图像名模式 |

**返回值:**

- `char*`: 匹配的图像名

***

## OCR 识别

### SetOcrEngine

设置 OCR 引擎。

**函数签名:**

```c
int SetOcrEngine(char* pathOfEngine, char* dllName, char* argv);
```

**参数:**

| 参数            | 类型     | 说明                     |
| --------------- | -------- | ------------------------ |
| pathOfEngine    | char*    | OCR引擎路径               |
| dllName         | char*    | DLL名称                  |
| argv            | char*    | 参数                     |

**返回值:**

- `1`: 成功

***

### SetDict

设置字库。

**函数签名:**

```c
int SetDict(int idx, char* fileName);
```

**参数:**

| 参数       | 类型     | 说明    |
| ---------- | -------- | ------- |
| idx        | int      | 字库索引 (0-9) |
| fileName   | char*    | 字库文件名 |

**返回值:**

- `1`: 成功

***

### GetDict

获取字库信息。

**函数签名:**

```c
char* GetDict(int idx, int fontIndex);
```

**参数:**

| 参数        | 类型  | 说明    |
| ----------- | ----- | ------- |
| idx         | int   | 字库索引 |
| fontIndex   | int   | 字体条目索引 |

**返回值:**

- `char*`: 字库信息字符串

***

### SetMemDict

设置内存字库。

**函数签名:**

```c
int SetMemDict(int idx, void* data, int size);
```

**参数:**

| 参数   | 类型   | 说明    |
| ------ | ------ | ------- |
| idx    | int    | 字库索引 (0-9) |
| data   | void*  | 字库数据  |
| size   | int    | 数据大小  |

**返回值:**

- `1`: 成功

***

### UseDict

使用字库。

**函数签名:**

```c
int UseDict(int idx);
```

**参数:**

| 参数   | 类型  | 说明    |
| ------ | ----- | ------- |
| idx    | int   | 字库索引 |

**返回值:**

- `1`: 成功

***

### AddDict

添加字库。

**函数签名:**

```c
int AddDict(int idx, char* dictInfo);
```

**参数:**

| 参数       | 类型     | 说明    |
| ---------- | -------- | ------- |
| idx        | int      | 字库索引 |
| dictInfo   | char*    | 字库信息 |

**返回值:**

- `1`: 成功

***

### SaveDict

保存字库。

**函数签名:**

```c
int SaveDict(int idx, char* fileName);
```

**参数:**

| 参数       | 类型     | 说明    |
| ---------- | -------- | ------- |
| idx        | int      | 字库索引 |
| fileName   | char*    | 输出文件名 |

**返回值:**

- `1`: 成功

***

### ClearDict

清空字库。

**函数签名:**

```c
int ClearDict(int idx);
```

**参数:**

| 参数   | 类型  | 说明    |
| ------ | ----- | ------- |
| idx    | int   | 字库索引 |

**返回值:**

- `1`: 成功

***

### GetDictCount

获取字库字符数量。

**函数签名:**

```c
int GetDictCount(int idx);
```

**参数:**

| 参数   | 类型  | 说明    |
| ------ | ----- | ------- |
| idx    | int   | 字库索引 |

**返回值:**

- `int`: 字符数量

***

### GetNowDict

获取当前字库索引。

**函数签名:**

```c
int GetNowDict();
```

**参数:** 无

**返回值:**

- `int`: 当前字库索引

***

### FetchWord

提取字符点阵。

**函数签名:**

```c
char* FetchWord(int x1, int y1, int x2, int y2, char* color, char* word);
```

**参数:**

| 参数   | 类型     | 说明    |
| ------ | -------- | ------- |
| x1     | int      | 左上角 X |
| y1     | int      | 左上角 Y |
| x2     | int      | 右下角 X |
| y2     | int      | 右下角 Y |
| color  | char*    | 颜色格式 |
| word   | char*    | 要提取的字符串 |

**返回值:**

- `char*`: 点阵信息

***

### GetWordsNoDict

不使用字库识别词组。

**函数签名:**

```c
char* GetWordsNoDict(int x1, int y1, int x2, int y2, char* color);
```

**参数:**

| 参数   | 类型     | 说明    |
| ------ | -------- | ------- |
| x1     | int      | 左上角 X |
| y1     | int      | 左上角 Y |
| x2     | int      | 右下角 X |
| y2     | int      | 右下角 Y |
| color  | char*    | 颜色格式 |

**返回值:**

- `char*`: 识别的词组信息

***

### GetWordResultCount

获取词组结果数量。

**函数签名:**

```c
int GetWordResultCount(char* result);
```

**参数:**

| 参数     | 类型     | 说明    |
| -------- | -------- | ------- |
| result   | char*    | 识别结果字符串 |

**返回值:**

- `int`: 词组数量

***

### GetWordResultPos

获取词组坐标。

**函数签名:**

```c
int GetWordResultPos(char* result, int index, int* x, int* y);
```

**参数:**

| 参数     | 类型     | 说明    |
| -------- | -------- | ------- |
| result   | char*    | 识别结果字符串 |
| index    | int      | 词组索引 |
| x        | int*     | X坐标(输出) |
| y        | int*     | Y坐标(输出) |

**返回值:**

- `1`: 成功

***

### GetWordResultStr

获取词组内容。

**函数签名:**

```c
char* GetWordResultStr(char* result, int index);
```

**参数:**

| 参数     | 类型     | 说明    |
| -------- | -------- | ------- |
| result   | char*    | 识别结果字符串 |
| index    | int      | 词组索引 |

**返回值:**

- `char*`: 词组内容

***

### Ocr

OCR 识别。

**函数签名:**

```c
char* Ocr(int x1, int y1, int x2, int y2, char* color, double sim);
```

**参数:**

| 参数    | 类型     | 说明            |
| ------- | -------- | --------------- |
| x1      | int      | 左上角 X         |
| y1      | int      | 左上角 Y         |
| x2      | int      | 右下角 X         |
| y2      | int      | 右下角 Y         |
| color   | char*    | 颜色格式          |
| sim     | double   | 相似度 (0.1-1.0) |

**返回值:**

- `char*`: 识别字符串

***

### OcrEx

OCR 识别（扩展）。

**函数签名:**

```c
char* OcrEx(int x1, int y1, int x2, int y2, char* color, double sim);
```

**参数:** 同上

**返回值:**

- `char*`: 识别结果，格式: "x,y|text|..."

***

### FindStr

查找字符串。

**函数签名:**

```c
int FindStr(int x1, int y1, int x2, int y2, char* strs, char* color, double sim, int* retX, int* retY);
```

**参数:**

| 参数    | 类型     | 说明      |
| ------- | -------- | --------- |
| x1      | int      | 左上角 X   |
| y1      | int      | 左上角 Y   |
| x2      | int      | 右下角 X   |
| y2      | int      | 右下角 Y   |
| strs    | char*    | 要查找的字符串 |
| color   | char*    | 颜色格式    |
| sim     | double   | 相似度     |
| retX    | int*     | 输出 X 坐标 |
| retY    | int*     | 输出 Y 坐标 |

**返回值:**

- `1`: 找到
- `0`: 未找到

***

### FindStrEx

查找字符串（扩展）。

**函数签名:**

```c
char* FindStrEx(int x1, int y1, int x2, int y2, char* strs, char* color, double sim);
```

**参数:** 同上

**返回值:**

- `char*`: 所有找到的位置

***

### OcrAuto

自动二值化OCR识别。

**函数签名:**

```c
char* OcrAuto(int x1, int y1, int x2, int y2, double sim);
```

**参数:**

| 参数    | 类型     | 说明            |
| ------- | -------- | --------------- |
| x1      | int      | 左上角 X         |
| y1      | int      | 左上角 Y         |
| x2      | int      | 右下角 X         |
| y2      | int      | 右下角 Y         |
| sim     | double   | 相似度 (0.1-1.0) |

**返回值:**

- `char*`: 识别字符串

***

### OcrFromFile

从图像文件识别文本。

**函数签名:**

```c
char* OcrFromFile(char* fileName, char* colorFormat, double sim);
```

**参数:**

| 参数          | 类型     | 说明            |
| ------------- | -------- | --------------- |
| fileName      | char*    | 图像文件名        |
| colorFormat   | char*    | 颜色格式          |
| sim           | double   | 相似度 (0.1-1.0) |

**返回值:**

- `char*`: 识别文本

***

### OcrAutoFromFile

从图像文件自动识别文本。

**函数签名:**

```c
char* OcrAutoFromFile(char* fileName, double sim);
```

**参数:**

| 参数       | 类型     | 说明            |
| ---------- | -------- | --------------- |
| fileName   | char*    | 图像文件名        |
| sim        | double   | 相似度 (0.1-1.0) |

**返回值:**

- `char*`: 识别文本

***

### FindLine

查找线条。

**函数签名:**

```c
char* FindLine(int x1, int y1, int x2, int y2, char* color, double sim);
```

**参数:**

| 参数    | 类型     | 说明            |
| ------- | -------- | --------------- |
| x1      | int      | 左上角 X         |
| y1      | int      | 左上角 Y         |
| x2      | int      | 右下角 X         |
| y2      | int      | 右下角 Y         |
| color   | char*    | 颜色格式          |
| sim     | double   | 相似度 (0.1-1.0) |

**返回值:**

- `char*`: 找到的线条信息

***

## A* 路径查找

### AStarFindPath

A* 算法查找路径。

**函数签名:**

```c
char* AStarFindPath(int mapWidth, int mapHeight, char* disablePoints, int beginX, int beginY, int endX, int endY);
```

**参数:**

| 参数            | 类型     | 说明             |
| --------------- | -------- | ---------------- |
| mapWidth        | int      | 地图宽度           |
| mapHeight       | int      | 地图高度           |
| disablePoints   | char*    | 障碍点，格式: "x1,y1|x2,y2|..." |
| beginX          | int      | 起点 X           |
| beginY          | int      | 起点 Y           |
| endX            | int      | 终点 X           |
| endY            | int      | 终点 Y           |

**返回值:**

- `char*`: 路径，格式: "x1,y1|x2,y2|..."

***

### FindNearestPos

查找最近位置。

**函数签名:**

```c
char* FindNearestPos(char* allPos, int posType, int x, int y);
```

**参数:**

| 参数      | 类型     | 说明                   |
| --------- | -------- | ---------------------- |
| allPos    | char*    | 所有位置，格式: "x1,y1|x2,y2|..." |
| posType   | int      | 距离类型 (0=欧几里得, 1=曼哈顿) |
| x         | int      | 目标 X                 |
| y         | int      | 目标 Y                 |

**返回值:**

- `char*`: 最近位置，格式: "x,y"

***

## 内存操作

### WriteData

写入进程内存。

**函数签名:**

```c
int WriteData(int hwnd, char* address, char* data, int size);
```

**参数:**

| 参数      | 类型     | 说明    |
| --------- | -------- | ------- |
| hwnd      | int      | 进程句柄 |
| address   | char*    | 内存地址 |
| data      | char*    | 要写入的数据 |
| size      | int      | 数据大小 |

**返回值:**

- `1`: 成功
- `0`: 失败（待实现）

***

### ReadData

读取进程内存。

**函数签名:**

```c
char* ReadData(int hwnd, char* address, int size);
```

**参数:**

| 参数      | 类型     | 说明    |
| --------- | -------- | ------- |
| hwnd      | int      | 进程句柄 |
| address   | char*    | 内存地址 |
| size      | int      | 数据大小 |

**返回值:**

- `char*`: 读取的数据
- `""`: 失败（待实现）

***

## 工具函数

### GetKeycode

获取虚拟键码。

**函数签名:**

```c
int GetKeycode(char* keyName);
```

**参数:**

| 参数       | 类型     | 说明    |
| ---------- | -------- | ------- |
| keyName    | char*    | 键名    |

**返回值:**

- `int`: 虚拟键码

***
