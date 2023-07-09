From golang:latest

# 创建文件夹
RUN mkdir /app

# 当前工作目录
WORKDIR /app

# 将当前文件夹的所有文件复制到容器中
COPY . /app

# 设置go代理
ENV  GO111MODULE=on GOPROXY=https://goproxy.cn/,direct
