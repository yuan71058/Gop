/**
 * GOP多线程示例 - C++版本
 * 
 * 功能: 打开3个记事本,分别后台绑定并输入文本
 * 编译环境: Visual Studio 2019+ 或 MinGW
 * 架构: x64
 */

#include <iostream>
#include <string>
#include <windows.h>
#include <thread>
#include <chrono>

// 定义函数指针类型
typedef int (*CreateOpFunc)();
typedef const char* (*VerFunc)();
typedef int (*SetPathFunc)(const char*);
typedef int (*FindWindowFunc)(const char*, const char*);
typedef int (*FindWindowExFunc)(int, const char*, const char*);
typedef int (*BindWindowFunc)(int, const char*, const char*, const char*, int);
typedef int (*UnBindWindowFunc)();
typedef int (*SendStringFunc)(int, const char*);
typedef int (*GetWindowProcessIdFunc)(int);
typedef int (*TerminateProcessFunc)(int);
typedef int (*DelayFunc)(int);

// 全局变量
HMODULE g_hModule = nullptr;
CreateOpFunc CreateOp = nullptr;
VerFunc Ver = nullptr;
SetPathFunc SetPath = nullptr;
FindWindowFunc FindWindow = nullptr;
FindWindowExFunc FindWindowEx = nullptr;
BindWindowFunc BindWindow = nullptr;
UnBindWindowFunc UnBindWindow = nullptr;
SendStringFunc SendString = nullptr;
GetWindowProcessIdFunc GetWindowProcessId = nullptr;
TerminateProcessFunc TerminateProcess = nullptr;
DelayFunc Delay = nullptr;

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
    FindWindowEx = (FindWindowExFunc)GetProcAddress(g_hModule, "FindWindowEx");
    BindWindow = (BindWindowFunc)GetProcAddress(g_hModule, "BindWindow");
    UnBindWindow = (UnBindWindowFunc)GetProcAddress(g_hModule, "UnBindWindow");
    SendString = (SendStringFunc)GetProcAddress(g_hModule, "SendString");
    GetWindowProcessId = (GetWindowProcessIdFunc)GetProcAddress(g_hModule, "GetWindowProcessId");
    TerminateProcess = (TerminateProcessFunc)GetProcAddress(g_hModule, "TerminateProcess");
    Delay = (DelayFunc)GetProcAddress(g_hModule, "Delay");

    // 检查所有函数是否加载成功
    if (!CreateOp || !Ver || !FindWindow || !BindWindow || !FindWindowEx || 
        !SendString || !Delay || !GetWindowProcessId || !TerminateProcess) {
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
 * 记事本任务结构
 */
struct NotepadTask {
    int id;
    const char* title;
    const char* text;
};

/**
 * 记事本任务函数
 * @param task 任务信息
 */
void NotepadTask(NotepadTask task) {
    std::cout << "[任务" << task.id << "] 开始执行..." << std::endl;

    // 1. 创建独立的OP对象
    int opId = CreateOp();
    std::cout << "[任务" << task.id << "] 创建OP对象成功, ID: " << opId << std::endl;

    // 2. 打开记事本
    std::cout << "[任务" << task.id << "] 打开记事本..." << std::endl;
    ShellExecuteA(NULL, "open", "notepad.exe", NULL, NULL, SW_SHOW);
    Delay(2000);

    // 3. 查找记事本窗口
    std::cout << "[任务" << task.id << "] 查找记事本窗口..." << std::endl;
    int hwnd = FindWindow("", task.title);
    if (hwnd == 0) {
        std::cerr << "[任务" << task.id << "] 未找到记事本窗口!" << std::endl;
        return;
    }
    std::cout << "[任务" << task.id << "] 找到记事本窗口, 句柄: " << hwnd << std::endl;

    // 4. 后台绑定窗口(GDI模式)
    std::cout << "[任务" << task.id << "] 后台绑定窗口(GDI模式)..." << std::endl;
    int bindRet = BindWindow(hwnd, "gdi", "normal", "normal", 0);
    if (bindRet == 1) {
        std::cout << "[任务" << task.id << "] 窗口绑定成功!" << std::endl;
    } else {
        std::cerr << "[任务" << task.id << "] 窗口绑定失败!" << std::endl;
        return;
    }

    // 5. 查找Edit控件
    std::cout << "[任务" << task.id << "] 查找Edit控件..." << std::endl;
    int editHwnd = FindWindowEx(hwnd, "Edit", "");
    if (editHwnd == 0) {
        std::cerr << "[任务" << task.id << "] 未找到Edit控件!" << std::endl;
        UnBindWindow();
        return;
    }
    std::cout << "[任务" << task.id << "] 找到Edit控件, 句柄: " << editHwnd << std::endl;

    // 6. 输入文本
    std::cout << "[任务" << task.id << "] 输入文本..." << std::endl;
    SendString(editHwnd, task.text);
    std::cout << "[任务" << task.id << "] 已输入: " << task.text << std::endl;

    // 7. 延时3秒
    std::cout << "[任务" << task.id << "] 延时3秒..." << std::endl;
    Delay(3000);

    // 8. 解绑窗口
    std::cout << "[任务" << task.id << "] 解绑窗口..." << std::endl;
    UnBindWindow();

    // 9. 结束进程
    std::cout << "[任务" << task.id << "] 结束记事本进程..." << std::endl;
    int pid = GetWindowProcessId(hwnd);
    if (pid > 0) {
        TerminateProcess(pid);
        std::cout << "[任务" << task.id << "] 进程已终止!" << std::endl;
    }

    std::cout << "[任务" << task.id << "] 任务完成!" << std::endl;
}

/**
 * 主函数 - 演示 GOP 多线程使用
 */
int main() {
    std::cout << "=== GOP C++ 多线程示例 ===" << std::endl;

    // 加载DLL
    const char* dllPath = "..\\dll\\GOP_x64.dll";
    if (!LoadGOPDll(dllPath)) {
        std::cerr << "插件注册失败!" << std::endl;
        return -1;
    }
    std::cout << "插件注册成功!" << std::endl;

    // 输出版本号
    std::cout << "GOP版本: " << Ver() << std::endl;

    // 设置工作路径
    SetPath("C:\\");

    // 定义3个任务
    NotepadTask tasks[3] = {
        {1, "无标题 - 记事本", "Hello GOP! 这是任务1输入的文本。"},
        {2, "无标题 - 记事本", "Hello GOP! 这是任务2输入的文本。"},
        {3, "无标题 - 记事本", "Hello GOP! 这是任务3输入的文本。"}
    };

    // 创建并启动3个线程
    std::thread threads[3];
    for (int i = 0; i < 3; i++) {
        threads[i] = std::thread(NotepadTask, tasks[i]);
        // 每个任务之间间隔500毫秒启动
        std::this_thread::sleep_for(std::chrono::milliseconds(500));
    }

    // 等待所有线程完成
    for (int i = 0; i < 3; i++) {
        threads[i].join();
    }

    // 释放DLL
    std::cout << "\n释放插件..." << std::endl;
    FreeGOPDll();

    std::cout << "\n=== 所有任务完成 ===" << std::endl;
    return 0;
}
