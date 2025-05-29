#!/bin/bash

# 确保脚本在错误时退出
set -e

# 显示彩色输出的函数
function echo_color() {
  local color=$1
  local message=$2
  case $color in
    "green") echo -e "\033[0;32m$message\033[0m" ;;
    "red") echo -e "\033[0;31m$message\033[0m" ;;
    "yellow") echo -e "\033[0;33m$message\033[0m" ;;
    "blue") echo -e "\033[0;34m$message\033[0m" ;;
    *) echo "$message" ;;
  esac
}

# 获取本地 IP 地址的函数
function get_local_ip() {
  if command -v ip &> /dev/null; then
    # Linux 系统
    ip addr show | grep "inet " | grep -v 127.0.0.1 | awk '{print $2}' | cut -d/ -f1 | head -n 1
  elif command -v ifconfig &> /dev/null; then
    # macOS 系统
    ifconfig | grep "inet " | grep -v 127.0.0.1 | awk '{print $2}' | head -n 1
  else
    echo "无法获取本地 IP 地址"
    return 1
  fi
}

# 检查 Docker 和 Docker Compose 是否已安装
if ! command -v docker &> /dev/null; then
  echo_color "red" "错误: Docker 未安装，请先安装 Docker"
  exit 1
fi

if ! command -v docker-compose &> /dev/null; then
  echo_color "red" "错误: Docker Compose 未安装，请先安装 Docker Compose"
  exit 1
fi

# 备份数据库（如果容器正在运行）
if docker ps | grep -q blog-mysql; then
  echo_color "blue" "正在备份数据库..."
  mkdir -p ./backups
  docker exec blog-mysql mysqldump -u root -proot mysql_blog > ./backups/blog_backup_$(date +%Y%m%d_%H%M%S).sql
  echo_color "green" "数据库备份完成"
fi

# 停止并删除旧容器（如果存在）
echo_color "blue" "正在停止旧容器..."
docker-compose down

# 构建并启动新容器
echo_color "blue" "正在构建并启动新容器..."
docker-compose up -d --build

# 等待服务启动
echo_color "blue" "等待服务启动..."
sleep 10

# 检查服务是否正常运行
if docker ps | grep -q blog-app && docker ps | grep -q blog-mysql; then
  echo_color "green" "✅ 部署成功完成！"
  echo_color "green" "应用现在可以通过以下地址访问:"
  local_ip=$(get_local_ip)
  echo_color "green" "http://$local_ip"
else
  echo_color "red" "❌ 部署可能存在问题，请检查容器日志:"
  echo_color "yellow" "docker-compose logs"
fi