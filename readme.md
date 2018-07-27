# 练习题文档

## 一、tcp client/server
   ### 1.1运行server：
       ./c_server +端口号（例：1234） +是否需要tls加密（0:不加密，1:加密）
   ### 1.2运行client: 
       ./c_client +地址（例：127.0.0.1:1234）+是否需要tls加密（0:不加密，1:加密）

## 二、websocket client/server，实现echo服务
### 2.1运行server:
        运行server：./ web_socket_server +端口号（例：1234）+是否需要tls加密（0:不加密，1:加密）
### 2.2运行client:
        运行client：./web_socket_Client +请求地址（例：127.0.0.1:1234）+是否需要tls加密（0:不加密，1:加密）

## 三、http client/server，实现几个基本的get, post请求处理。
### 3.1运行server： 运行server：
       ./ h_server +host（例：127.0.0.1）+端口号（例：1234）+是否需要tls加密（0:不加密，1:加密）
### 3.2运行client:
       运行client：./h_client +请求地址（例：127.0.0.1:1234）+用户名（例：zhangsan）+密码（例：123456）+是否需要tls加密（0:不加密，1:加密）+ 使用请求的方式（1:get 0:post）

## 四、grpc client/server，实现一个基本的grpc调用请求处理（RpcServer放在server文件下，RpcClient放在client文件下）

### 4.1运行server：
        ./rpc_server +地址（例：127.0.0.1:1234） +是否需要tls加密（0:不加密，1:加密）
### 4.2运行client: 
        ./rpc_client +地址（例：127.0.0.1:1234）+是否需要tls加密（0:不加密，1:加密）

## 注：在运行tls加密时需要生成server.pem 和 server.key
### 1.生成服务器端的私钥
        openssl genrsa -out server.key 2048
### 2.生成服务器端证书
        openssl req -new -x509 -key server.key -out server.pem -days 3650
        or
        go run $GOROOT/src/crypto/tls/generate_cert.go --host localhost
