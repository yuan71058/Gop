/**
 * GOP多线程示例 - C++版本
 * 
 * 功能: 打开3个记事本,分别后台绑定并输入50行文本
 * 编译环境: Visual Studio 2019+ 或 MinGW
 * 架构: x64
 */

#include <iostream>
#include <string>
#include <vector>
#include <sstream>
#include <windows.h>
#include <thread>
#include <chrono>
#include <mutex>

// 定义函数指针类型
typedef int (*CreateOpFunc)();
typedef const char* (*VerFunc)();
typedef int (*SetPathFunc)(const char*);
typedef int (*FindWindowFunc)(const char*, const char*);
typedef int (*FindWindowExFunc)(int, const char*, const char*);
typedef int (*BindWindowFunc)(int, const char*, const char*, const char*, int);
typedef int (*UnBindWindowFunc)();
typedef int (*SendStringFunc)(int, const char*);
typedef int (*KeyPressFunc)(int);
typedef int (*MoveWindowFunc)(int, int, int);
typedef int (*GetWindowProcessIdFunc)(int);
typedef int (*TerminateProcessFunc)(int);
typedef int (*DelayFunc)(int);
typedef int (*EnumProcessFunc)(const char*);
typedef int (*FindWindowByProcessIdFunc)(int, const char*, const char*);

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
KeyPressFunc KeyPress = nullptr;
MoveWindowFunc MoveWindow = nullptr;
GetWindowProcessIdFunc GetWindowProcessId = nullptr;
TerminateProcessFunc TerminateProcess = nullptr;
DelayFunc Delay = nullptr;
EnumProcessFunc EnumProcess = nullptr;
FindWindowByProcessIdFunc FindWindowByProcessId = nullptr;

std::mutex g_mtx;

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
    KeyPress = (KeyPressFunc)GetProcAddress(g_hModule, "KeyPress");
    MoveWindow = (MoveWindowFunc)GetProcAddress(g_hModule, "MoveWindow");
    GetWindowProcessId = (GetWindowProcessIdFunc)GetProcAddress(g_hModule, "GetWindowProcessId");
    TerminateProcess = (TerminateProcessFunc)GetProcAddress(g_hModule, "TerminateProcess");
    Delay = (DelayFunc)GetProcAddress(g_hModule, "Delay");
    EnumProcess = (EnumProcessFunc)GetProcAddress(g_hModule, "EnumProcess");
    FindWindowByProcessId = (FindWindowByProcessIdFunc)GetProcAddress(g_hModule, "FindWindowByProcessId");

    // 检查所有函数是否加载成功
    if (!CreateOp || !Ver || !FindWindow || !BindWindow || !FindWindowEx || 
        !SendString || !KeyPress || !MoveWindow || !Delay || !GetWindowProcessId || 
        !TerminateProcess || !EnumProcess || !FindWindowByProcessId) {
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
 * 生成多行文本内容
 * @param threadID 线程ID
 * @param lineCount 行数
 * @return 生成的文本内容
 */
std::string generateTextContent(int threadID, int lineCount) {
    std::ostringstream oss;
    for (int i = 1; i <= lineCount; i++) {
        oss << "线程" << threadID << " - 第" << i << "行: GOP多线程文本输入测试内容。" << std::endl;
    }
    return oss.str();
}

/**
 * 解析进程ID列表(竖线分隔)
 * @param pidList 进程ID列表字符串
 * @return 进程ID向量
 */
std::vector<int> parsePidList(const std::string& pidList) {
    std::vector<int> pids;
    std::istringstream iss(pidList);
    std::string token;
    while (std::getline(iss, token, '|')) {
        if (!token.empty()) {
            try {
                pids.push_back(std::stoi(token));
            } catch (...) {}
        }
    }
    return pids;
}

/**
 * 记事本任务结构
 */
struct NotepadTask {
    int id;
    int hwnd;
    int editHwnd;
    std::string content;
};

/**
 * 记事本任务函数
 * @param task 任务信息
 */
void NotepadTask(NotepadTask task) {
    std::lock_guard<std::mutex> lock(g_mtx);
    std::cout << "[线程" << task.id << "] 开始执行..." << std::endl;

    // 1. 创建独立的OP对象
    int opId = CreateOp();
    std::cout << "[线程" << task.id << "] 创建OP对象成功, ID: " << opId << std::endl;

    // 2. 后台绑定窗口(GDI模式 + Windows键盘模式)
    std::cout << "[线程" << task.id << "] 后台绑定窗口(Windows键盘模式)..." << std::endl;
    int bindRet = BindWindow(task.editHwnd, "gdi", "normal", "windows", 0);
    if (bindRet == 1) {
        std::cout << "[线程" << task.id << "] 窗口绑定成功!" << std::endl;
    } else {
        std::cerr << "[线程" << task.id << "] 窗口绑定失败!" << std::endl;
        return;
    }

    // 3. 输入文本(逐行输入,每行延时200毫秒)
    std::cout << "[线程" << task.id << "] 开始写入文字..." << std::endl;
    
    std::istringstream iss(task.content);
    std::string line;
    int totalChars = 0;
    int lineCount = 0;
    
    while (std::getline(iss, line)) {
        if (line.empty()) continue;
        
        // 发送当前行文本
        SendString(task.editHwnd, line.c_str());
        totalChars += line.length();
        
        // 发送回车键
        KeyPress(13); // VK_RETURN = 13
        totalChars++;
        lineCount++;
        
        // 每行延时200毫秒
        Delay(200);
    }
    
    std::cout << "[线程" << task.id << "] 写入完成, 共" << totalChars << "个字符, " 
              << lineCount << "行" << std::endl;

    // 4. 解绑窗口
    std::cout << "[线程" << task.id << "] 解绑窗口..." << std::endl;
    UnBindWindow();

    std::cout << "[线程" << task.id << "] 任务完成!" << std::endl;
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

    // 第一步: 打开3个记事本
    std::cout << "\n========== 第一步: 打开记事本 ==========" << std::endl;
    PROCESS_INFORMATION processes[3];
    for (int i = 0; i < 3; i++) {
        STARTUPINFOA si = { sizeof(si) };
        si.dwFlags = STARTF_USESHOWWINDOW;
        si.wShowWindow = SW_SHOW;
        
        if (CreateProcessA(NULL, "notepad.exe", NULL, NULL, FALSE, 0, NULL, NULL, &si, &processes[i])) {
            std::cout << "打开记事本 " << (i + 1) << " (PID: " << processes[i].dwProcessId << ")" << std::endl;
        }
    }
    
    // 等待窗口创建
    std::cout << "\n等待窗口打开..." << std::endl;
    Delay(2000);

    // 第二步: 查找记事本窗口
    std::cout << "\n========== 第二步: 查找记事本窗口 ==========" << std::endl;
    int mainOp = CreateOp();
    
    std::string processList = EnumProcess("notepad.exe");
    std::cout << "进程列表: " << processList << std::endl;
    
    std::vector<int> pids = parsePidList(processList);
    std::cout << "找到 " << pids.size() << " 个notepad进程" << std::endl;
    
    std::vector<int> hwnds;
    for (int pid : pids) {
        if (hwnds.size() >= 3) break;
        int hwnd = FindWindowByProcessId(pid, "Notepad", "");
        if (hwnd > 0) {
            hwnds.push_back(hwnd);
            std::cout << "找到记事本窗口 (PID=" << pid << "): " << hwnd << std::endl;
        }
    }
    
    if (hwnds.size() < 3) {
        std::cout << "记事本窗口数量不足!" << std::endl;
        return -1;
    }

    // 第三步: 智能排列窗口
    std::cout << "\n========== 第三步: 智能排列窗口 ==========" << std::endl;
    int windowWidth = 800;
    int windowHeight = 600;
    int margin = 20;
    
    for (int i = 0; i < 3; i++) {
        int row = i / 2;
        int col = i % 2;
        int x = col * (windowWidth + margin);
        int y = row * (windowHeight + margin);
        
        int ret = MoveWindow(hwnds[i], x, y);
        if (ret == 1) {
            std::cout << "窗口" << (i + 1) << " 移动到 (" << x << ", " << y << ")" << std::endl;
        }
    }
    
    Delay(500);

    // 第四步: 查找Edit控件
    std::cout << "\n========== 第四步: 查找Edit编辑框控件 ==========" << std::endl;
    std::vector<int> editHwnds;
    for (int i = 0; i < 3; i++) {
        int editHwnd = FindWindowEx(hwnds[i], "Edit", "");
        if (editHwnd == 0) {
            std::cout << "找不到记事本 " << (i + 1) << " 的Edit编辑框" << std::endl;
            return -1;
        }
        editHwnds.push_back(editHwnd);
        std::cout << "记事本 " << (i + 1) << ": 主窗口=" << hwnds[i] << ", Edit控件=" << editHwnd << std::endl;
    }

    // 第五步: 创建任务并启动线程
    std::cout << "\n========== 第五步: 开始多线程写入 ==========" << std::endl;
    
    NotepadTask tasks[3];
    for (int i = 0; i < 3; i++) {
        tasks[i].id = i + 1;
        tasks[i].hwnd = hwnds[i];
        tasks[i].editHwnd = editHwnds[i];
        tasks[i].content = generateTextContent(i + 1, 50);
    }
    
    auto startTime = std::chrono::high_resolution_clock::now();
    
    std::thread threads[3];
    for (int i = 0; i < 3; i++) {
        threads[i] = std::thread(NotepadTask, tasks[i]);
    }
    
    for (int i = 0; i < 3; i++) {
        threads[i].join();
    }
    
    auto elapsed = std::chrono::duration_cast<std::chrono::milliseconds>(
        std::chrono::high_resolution_clock::now() - startTime);
    std::cout << "\n========== 所有线程完成 ==========" << std::endl;
    std::cout << "总耗时: " << elapsed.count() << "ms" << std::endl;

    // 等待5秒查看效果
    std::cout << "\n等待5秒,查看效果..." << std::endl;
    Delay(5000);

    // 第六步: 结束进程
    std::cout << "\n========== 第六步: 结束进程 ==========" << std::endl;
    for (int i = 0; i < 3; i++) {
        int pid = GetWindowProcessId(hwnds[i]);
        if (pid > 0) {
            TerminateProcess(pid);
            std::cout << "结束进程: 记事本 " << (i + 1) << " (PID: " << pid << ")" << std::endl;
        }
    }

    // 第七步: 释放资源
    std::cout << "\n========== 第七步: 释放资源 ==========" << std::endl;
    FreeGOPDll();
    std::cout << "GOP多线程示例完成" << std::endl;

    std::cout << "\n========== 示例完成 ==========" << std::endl;
    return 0;
}
