@echo off
REM ============================================================
REM GOP Plugin DLL Build Script
REM Supports building x86 and x64 DLLs
REM ============================================================

echo ========================================
echo GOP Plugin DLL Build Script
echo ========================================
echo.

REM Check if Go is installed
where go >nul 2>nul
if %errorlevel% neq 0 (
    echo [ERROR] Go not found. Please install Go first.
    pause
    exit /b 1
)

echo [INFO] Go Version:
go version
echo.

REM Create output directory
if not exist "build" mkdir build

REM Download dependencies
echo [INFO] Downloading dependencies...
go mod tidy
echo.

REM Build x64 DLL
echo ========================================
echo Building x64 DLL...
echo ========================================
set GOARCH=amd64
set CGO_ENABLED=1
go build -buildmode=c-shared -o build/gop_amd64.dll ./dll
if %errorlevel% equ 0 (
    echo [SUCCESS] x64 DLL built: build/gop_amd64.dll
) else (
    echo [FAILED] x64 DLL build failed
    pause
    exit /b 1
)
echo.

REM Build x86 DLL
echo ========================================
echo Building x86 DLL...
echo ========================================
set GOARCH=386
go build -buildmode=c-shared -o build/gop_386.dll ./dll
if %errorlevel% equ 0 (
    echo [SUCCESS] x86 DLL built: build/gop_386.dll
) else (
    echo [FAILED] x86 DLL build failed
    pause
    exit /b 1
)
echo.

REM Show results
echo ========================================
echo Build Complete!
echo ========================================
echo Output files:
dir /b build\*.dll
echo.
echo File sizes:
for %%f in (build\*.dll) do (
    echo   %%~nf: %%~zf bytes
)
echo.

pause
