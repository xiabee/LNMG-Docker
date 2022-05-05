# LNMG-Docker

Go web server with docker

## 项目描述

- `LNMG`：`Linux`+`nignx`+`mysql`+`Golang` 服务器
- 主要用于一键式部署`Golang-WEB`环境

## 使用方式

```bash
docker-compose up -d
```

## 当前进度

* 实现`Golang`服务器的编译运行
* 实现`Golang`与`Nginx`的连接

## 

## 进行中的工作

* 连接数据库



## 注意事项

* 主程序`./www/main.go`默认开放端口为`8000`，`./nginx/nginx.conf`监听内网服务端口为`8000`，并将`8000`端口的流量转发至`80`端口供外网访问。若需修改内网监听端口，需要同时修改此两处：
  
  ```go
  // main.go
  err := http.ListenAndServe(":8000", nil)
  	// 与nginx.conf中的upstream一致
  ```
  
  ```nginx
  # nginx.conf
  upstream golang-handler {
      server golang:8000;
      # 与main.go中设置的端口一致
  }
  ```

* 外网端口映射通过`docker-compose`解决：
  
  ```yml
  nginx:
     ...
      ports:
        - "80:80"
  
  ```
