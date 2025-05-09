# 使用官方 Go 镜像作为基础镜像
FROM golang:1.20

# 设置工作目录
WORKDIR /app

# 将 go.mod 和 go.sum 复制到工作目录
COPY go.mod go.sum ./

# 下载依赖
RUN go mod download

# 将项目的所有文件复制到工作目录
COPY . .

# 构建 Go 应用
RUN go build -o main .

# 暴露应用运行的端口
EXPOSE 3000

# 设置环境变量（可选）
ENV PORT=3000

# 运行应用
CMD ["./main"]