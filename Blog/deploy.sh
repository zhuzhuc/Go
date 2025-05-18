#!/bin/bash

# 确保脚本在错误时退出
set -e

# 构建前端
echo "Building frontend..."
cd client
npm run build
cd ..

# 停止并删除旧容器（如果存在）
echo "Stopping old containers..."
docker-compose down

# 构建并启动新容器
echo "Building and starting new containers..."
docker-compose up -d --build

echo "Deployment completed successfully!" 