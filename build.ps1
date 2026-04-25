# ============================================================
# OP插件DLL构建脚本 (PowerShell版本)
# 支持编译为x86和x64架构的DLL
# ============================================================

$ErrorActionPreference = "Stop"

Write-Host "========================================" -ForegroundColor Cyan
Write-Host "OP插件DLL构建脚本" -ForegroundColor Cyan
Write-Host "========================================" -ForegroundColor Cyan
Write-Host ""

# 检查Go是否安装
try {
    $goVersion = go version
    Write-Host "[信息] Go版本: $goVersion" -ForegroundColor Green
} catch {
    Write-Host "[错误] 未找到Go，请先安装Go环境" -ForegroundColor Red
    Read-Host "按回车键退出"
    exit 1
}

# 创建输出目录
if (-not (Test-Path "build")) {
    New-Item -ItemType Directory -Path "build" | Out-Null
}

# 下载依赖
Write-Host "[信息] 下载依赖..." -ForegroundColor Yellow
go mod tidy
Write-Host ""

# 编译x64 DLL
Write-Host "========================================" -ForegroundColor Cyan
Write-Host "开始编译 x64 DLL..." -ForegroundColor Cyan
Write-Host "========================================" -ForegroundColor Cyan
$env:GOARCH = "amd64"
$env:CGO_ENABLED = "1"
go build -buildmode=c-shared -o build/gop_amd64.dll ./dll
if ($LASTEXITCODE -eq 0) {
    Write-Host "[成功] x64 DLL编译完成: build/gop_amd64.dll" -ForegroundColor Green
} else {
    Write-Host "[失败] x64 DLL编译失败" -ForegroundColor Red
    Read-Host "按回车键退出"
    exit 1
}
Write-Host ""

# 编译x86 DLL
Write-Host "========================================" -ForegroundColor Cyan
Write-Host "开始编译 x86 DLL..." -ForegroundColor Cyan
Write-Host "========================================" -ForegroundColor Cyan
$env:GOARCH = "386"
go build -buildmode=c-shared -o build/gop_386.dll ./dll
if ($LASTEXITCODE -eq 0) {
    Write-Host "[成功] x86 DLL编译完成: build/gop_386.dll" -ForegroundColor Green
} else {
    Write-Host "[失败] x86 DLL编译失败" -ForegroundColor Red
    Read-Host "按回车键退出"
    exit 1
}
Write-Host ""

# 显示编译结果
Write-Host "========================================" -ForegroundColor Cyan
Write-Host "编译完成！" -ForegroundColor Green
Write-Host "========================================" -ForegroundColor Cyan
Write-Host "输出文件:" -ForegroundColor Yellow
Get-ChildItem -Path "build\*.dll" | ForEach-Object {
    $size = [math]::Round($_.Length / 1KB, 2)
    Write-Host "  $($_.Name): $size KB" -ForegroundColor White
}
Write-Host ""

Read-Host "按回车键退出"
