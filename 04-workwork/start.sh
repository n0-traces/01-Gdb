#!/bin/bash

echo "=== 个人博客系统启动脚本 ==="

# 检查Go是否安装
if ! command -v go &> /dev/null; then
    echo "错误: Go未安装，请先安装Go 1.24或更高版本"
    exit 1
fi

# 检查MySQL是否安装
if ! command -v mysql &> /dev/null; then
    echo "警告: MySQL未安装，请先安装MySQL"
    echo "参考: DATABASE_SETUP.md"
fi

# 检查Go版本
GO_VERSION=$(go version | awk '{print $3}' | sed 's/go//')
echo "检测到Go版本: $GO_VERSION"

# 安装依赖
echo "正在安装依赖..."
go mod tidy

if [ $? -ne 0 ]; then
    echo "错误: 依赖安装失败"
    exit 1
fi

echo "依赖安装完成"

# 启动服务器
echo "正在启动服务器..."
echo "服务器将在 http://localhost:8080 启动"
echo "按 Ctrl+C 停止服务器"
echo ""

go run .
