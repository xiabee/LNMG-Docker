FROM golang:alpine
LABEL maintainer="xiabe.cn <xiabee@foxmail.com>"

ENV GO111MODULE on
ENV GOPROXY https://mirrors.aliyun.com/goproxy/
# # 修改GO的镜像地址

WORKDIR /home/xiabee


ADD go.mod .
ADD go.sum .
RUN go mod download

ADD *.go .
ADD ./public ./public
RUN GOOS=linux CGO_ENABLED=0 go build -ldflags="-s -w" -installsuffix cgo -o main main.go

# RUN go build -o main

CMD [ "./main" ]