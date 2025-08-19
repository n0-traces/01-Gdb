@echo off
echo === 个人博客系统启动脚本 ===

REM 检查Go是否安装
go version >nul 2>&1
if %errorlevel% neq 0 (
    echo 错误: Go未安装，请先安装Go 1.24或更高版本
    pause
    exit /b 1
)

REM 检查MySQL是否安装
mysql --version >nul 2>&1
if %errorlevel% neq 0 (
    echo 警告: MySQL未安装，请先安装MySQL
    echo 参考: DATABASE_SETUP.md
)

REM 显示Go版本
for /f "tokens=3" %%i in ('go version') do set GO_VERSION=%%i
echo 检测到Go版本: %GO_VERSION%

REM 安装依赖
echo 正在安装依赖...
go mod tidy

if %errorlevel% neq 0 (
    echo 错误: 依赖安装失败
    pause
    exit /b 1
)

echo 依赖安装完成

REM 启动服务器
echo 正在启动服务器...
echo 服务器将在 http://localhost:8080 启动
echo 按 Ctrl+C 停止服务器
echo.

go run .
