# 多阶段构建

# 前端构建阶段
FROM node:18-alpine AS frontend-builder
WORKDIR /app
COPY client/package*.json ./
RUN npm install
COPY client/ ./
RUN npm run build

# 后端构建阶段
FROM golang:1.24-alpine AS backend-builder
WORKDIR /app
# 复制 go.mod 和 go.sum 文件
COPY server/go.mod server/go.sum ./
# 下载依赖
RUN go mod download
# 复制源代码
COPY server/ ./
# 构建应用
RUN CGO_ENABLED=0 GOOS=linux go build -o server .

# 最终镜像
FROM nginx:alpine
WORKDIR /app

# 复制 Nginx 配置
COPY nginx/default.conf /etc/nginx/conf.d/default.conf

# 复制前端构建产物
COPY --from=frontend-builder /app/build /usr/share/nginx/html

# 创建后端目录
RUN mkdir -p /app/server

# 复制后端构建产物
COPY --from=backend-builder /app/server /app/server

# 复制环境配置文件
COPY server/.env /app/

# 复制启动脚本
COPY docker-entrypoint.sh /app/
RUN chmod +x /app/docker-entrypoint.sh

# 创建上传目录
RUN mkdir -p /app/uploads && chmod 777 /app/uploads

# 暴露端口
EXPOSE 80 8000

# 设置入口点
ENTRYPOINT ["/app/docker-entrypoint.sh"]
