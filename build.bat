@echo off
REM ============================================================
REM OP插件DLL构建脚本
REM 支持编译为x86和x64架构的DLL
REM ============================================================

echo ========================================
echo OP插件DLL构建脚本
echo ========================================
echo.

REM 检查Go是否安装
where go >nul 2>nul
if %errorlevel% neq 0 (
    echo [错误] 未找到Go，请先安装Go环境
    pause
    exit /b 1
)

echo [信息] Go版本:
go version
echo.

REM 创建输出目录
if not exist "build" mkdir build

REM 下载依赖
echo [信息] 下载依赖...
go mod tidy
echo.

REM 编译x64 DLL
echo ========================================
echo 开始编译 x64 DLL...
echo ========================================
set GOARCH=amd64
set CGO_ENABLED=1
go build -buildmode=c-shared -o build/gop_amd64.dll ./dll
if %errorlevel% equ 0 (
    echo [成功] x64 DLL编译完成: build/gop_amd64.dll
) else (
    echo [失败] x64 DLL编译失败
    pause
    exit /b 1
)
echo.

REM 编译x86 DLL
echo ========================================
echo 开始编译 x86 DLL...
echo ========================================
set GOARCH=386
go build -buildmode=c-shared -o build/gop_386.dll ./dll
if %errorlevel% equ 0 (
    echo [成功] x86 DLL编译完成: build/gop_386.dll
) else (
    echo [失败] x86 DLL编译失败
    pause
    exit /b 1
)
echo.

REM 显示编译结果
echo ========================================
echo 编译完成！
echo ========================================
echo 输出文件:
dir /b build\*.dll
echo.
echo 文件大小:
for %%f in (build\*.dll) do (
    echo   %%~nf: %%~zf 字节
)
echo.

pause
