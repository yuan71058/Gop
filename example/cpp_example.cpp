/**
 * GOP C++ 使用示例
 * 
 * 本示例演示如何在 C++ 中使用 GOP DLL
 * 功能: 注册插件 -> 输出版本号 -> 打开记事本 -> 后台绑定窗口 -> 输入文本 -> 延时5秒 -> 关闭窗口
 * 编译环境: Visual Studio 2019+ 或 MinGW
 * 架构: x64
 */

#include <iostream>
#include <string>
#include <windows.h>

// 定义函数指针类型
typedef int (*CreateOpFunc)();
typedef const char* (*VerFunc)();
typedef int (*SetPathFunc)(const char*);
typedef int (*FindWindowFunc)(const char*, const char*);
typedef int (*GetWindowRectFunc)(int, int*, int*, int*, int*);
typedef int (*GetClientSizeFunc)(int, int*, int*);
typedef int (*BindWindowFunc)(int, const char*, const char*, const char*, int);
typedef int (*UnBindWindowFunc)();
typedef int (*KeyPressCharFunc)(const char*);
typedef int (*DelayFunc)(int);
typedef int (*CloseWindowFunc)(int);

// 全局变量
HMODULE g_hModule = nullptr;
CreateOpFunc CreateOp = nullptr;
VerFunc Ver = nullptr;
SetPathFunc SetPath = nullptr;
FindWindowFunc FindWindow = nullptr;
GetWindowRectFunc GetWindowRect = nullptr;
GetClientSizeFunc GetClientSize = nullptr;
BindWindowFunc BindWindow = nullptr;
UnBindWindowFunc UnBindWindow = nullptr;
KeyPressCharFunc KeyPressChar = nullptr;
DelayFunc Delay = nullptr;
CloseWindowFunc CloseWindow = nullptr;

/**
 * 加载 GOP DLL
 * @param dllPath DLL 文件路径
 * @return true 成功, false 失败
 */
bool LoadGOPDll(const char* dllPath) {
    g_hModule = LoadLibraryA(dllPath);
    if (!g_hModule) {
        std::cerr << "无法加载 GOP DLL: " << dllPath << std::endl;
        return false;
    }

    // 获取函数地址
    CreateOp = (CreateOpFunc)GetProcAddress(g_hModule, "CreateOp");
    Ver = (VerFunc)GetProcAddress(g_hModule, "Ver");
    SetPath = (SetPathFunc)GetProcAddress(g_hModule, "SetPath");
    FindWindow = (FindWindowFunc)GetProcAddress(g_hModule, "FindWindow");
    GetWindowRect = (GetWindowRectFunc)GetProcAddress(g_hModule, "GetWindowRect");
    GetClientSize = (GetClientSizeFunc)GetProcAddress(g_hModule, "GetClientSize");
    BindWindow = (BindWindowFunc)GetProcAddress(g_hModule, "BindWindow");
    UnBindWindow = (UnBindWindowFunc)GetProcAddress(g_hModule, "UnBindWindow");
    KeyPressChar = (KeyPressCharFunc)GetProcAddress(g_hModule, "KeyPressChar");
    Delay = (DelayFunc)GetProcAddress(g_hModule, "Delay");
    CloseWindow = (CloseWindowFunc)GetProcAddress(g_hModule, "CloseWindow");

    // 检查所有函数是否加载成功
    if (!CreateOp || !Ver || !FindWindow || !BindWindow || !KeyPressChar || !Delay || !CloseWindow) {
        std::cerr << "无法获取 GOP 函数地址" << std::endl;
        FreeLibrary(g_hModule);
        g_hModule = nullptr;
        return false;
    }

    return true;
}

/**
 * 释放 GOP DLL
 */
void FreeGOPDll() {
    if (g_hModule) {
        FreeLibrary(g_hModule);
        g_hModule = nullptr;
    }
}

/**
 * 主函数 - 演示 GOP 基本使用流程
 */
int main() {
    std::cout << "=== GOP C++ 示例 ===" << std::endl;

    // 1. 注册插件(加载DLL)
    std::cout << "\n1. 注册插件..." << std::endl;
    const char* dllPath = "..\\dll\\GOP_x64.dll";
    if (!LoadGOPDll(dllPath)) {
        std::cerr << "插件注册失败!" << std::endl;
        return -1;
    }
    std::cout << "插件注册成功!" << std::endl;

    // 创建OP实例
    int opId = CreateOp();
    std::cout << "OP实例ID: " << opId << std::endl;

    // 2. 输出版本号
    std::cout << "\n2. 输出版本号..." << std::endl;
    const char* version = Ver();
    std::cout << "GOP版本: " << version << std::endl;

    // 设置工作路径
    SetPath("C:\\");

    // 3. 打开记事本
    std::cout << "\n3. 打开记事本..." << std::endl;
    ShellExecuteA(NULL, "open", "notepad.exe", NULL, NULL, SW_SHOW);
    std::cout << "正在启动记事本..." << std::endl;
    
    // 等待记事本启动
    Delay(2000);

    // 4. 查找记事本窗口
    std::cout << "\n4. 查找记事本窗口..." << std::endl;
    int hwnd = FindWindow("", "无标题 - 记事本");
    if (hwnd == 0) {
        std::cerr << "未找到记事本窗口!" << std::endl;
        FreeGOPDll();
        return -1;
    }
    std::cout << "找到记事本窗口, 句柄: " << hwnd << std::endl;

    // 获取窗口信息
    int x1, y1, x2, y2;
    GetWindowRect(hwnd, &x1, &y1, &x2, &y2);
    std::cout << "窗口位置: (" << x1 << ", " << y1 << ") - (" << x2 << ", " << y2 << ")" << std::endl;

    int width, height;
    GetClientSize(hwnd, &width, &height);
    std::cout << "客户区大小: " << width << " x " << height << std::endl;

    // 5. 后台绑定窗口
    std::cout << "\n5. 后台绑定窗口..." << std::endl;
    int bindRet = BindWindow(hwnd, "normal", "normal", "normal", 0);
    if (bindRet == 1) {
        std::cout << "窗口绑定成功!" << std::endl;
    } else {
        std::cerr << "窗口绑定失败!" << std::endl;
        CloseWindow(hwnd);
        FreeGOPDll();
        return -1;
    }

    // 6. 向记事本窗口输入一句话
    std::cout << "\n6. 向记事本输入文本..." << std::endl;
    const char* text = "Hello GOP! 这是从C++示例输入的文本。";
    KeyPressChar(text);
    std::cout << "已输入: " << text << std::endl;

    // 7. 延时5秒
    std::cout << "\n7. 延时5秒..." << std::endl;
    std::cout << "等待中..." << std::endl;
    Delay(5000);
    std::cout << "延时结束!" << std::endl;

    // 8. 解绑窗口
    std::cout << "\n8. 解绑窗口..." << std::endl;
    UnBindWindow();
    std::cout << "窗口已解绑!" << std::endl;

    // 9. 关闭记事本窗口
    std::cout << "\n9. 关闭记事本窗口..." << std::endl;
    CloseWindow(hwnd);
    std::cout << "窗口已关闭!" << std::endl;

    // 10. 释放DLL
    std::cout << "\n10. 释放插件..." << std::endl;
    FreeGOPDll();
    std::cout << "插件已释放!" << std::endl;

    std::cout << "\n=== 示例完成 ===" << std::endl;
    return 0;
}
