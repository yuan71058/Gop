# OP插件DLL构建脚本
# 支持编译为x86和x64架构的DLL

.PHONY: all clean x64 x86 tidy

# 输出目录
BUILD_DIR := build

# 默认目标：编译x64和x86
all: tidy x64 x86

# 下载依赖
tidy:
	@echo [信息] 下载依赖...
	@go mod tidy

# 编译x64 DLL
x64:
	@echo ========================================
	@echo 开始编译 x64 DLL...
	@echo ========================================
	@mkdir -p $(BUILD_DIR)
	@GOARCH=amd64 CGO_ENABLED=1 go build -buildmode=c-shared -o $(BUILD_DIR)/gop_amd64.dll ./dll
	@echo [成功] x64 DLL编译完成: $(BUILD_DIR)/gop_amd64.dll

# 编译x86 DLL
x86:
	@echo ========================================
	@echo 开始编译 x86 DLL...
	@echo ========================================
	@mkdir -p $(BUILD_DIR)
	@GOARCH=386 CGO_ENABLED=1 go build -buildmode=c-shared -o $(BUILD_DIR)/gop_386.dll ./dll
	@echo [成功] x86 DLL编译完成: $(BUILD_DIR)/gop_386.dll

# 清理构建文件
clean:
	@echo [信息] 清理构建文件...
	@rm -rf $(BUILD_DIR)
	@echo [成功] 清理完成

# 显示帮助信息
help:
	@echo OP插件DLL构建脚本
	@echo.
	@echo 可用目标:
	@echo   all     - 编译x64和x86 DLL (默认)
	@echo   x64     - 仅编译x64 DLL
	@echo   x86     - 仅编译x86 DLL
	@echo   tidy    - 下载依赖
	@echo   clean   - 清理构建文件
	@echo   help    - 显示帮助信息
