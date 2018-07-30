# 练习题文档

## 一、tcp client/server
   ### 1.1运行server：
       ./c_server -port（例：1234） -tls（0:不加密，1:加密）
   ### 1.2运行client: 
       ./c_client -address（例：127.0.0.1:1234）-tls（0:不加密，1:加密）

## 二、websocket client/server，实现echo服务
### 2.1运行server:
        运行server：./ web_socket_server -port（例：1234）-tls（0:不加密，1:加密）
### 2.2运行client:
        运行client：./web_socket_Client -address（例：127.0.0.1:1234）-tls（0:不加密，1:加密）
        WebSocket Client 在chrome console 使用
        var wsServer = 'ws://localhost:1234';
        var websocket = new WebSocket(wsServer);
        websocket.onopen = function (evt) {
        console.log("Connected to WebSocket server.");
        };
        websocket.onclose = function (evt) {
        console.log("Disconnected");
        };
        websocket.onmessage = function (evt) {
        console.log('Retrieved data from server: ' + evt.data);
        };
        websocket.onerror = function (evt, e) {
        console.log('Error occured: ' + evt.data);
        };

## 三、http client/server，实现几个基本的get, post请求处理。
### 3.1运行server： 运行server：
       ./ h_server +host（例：127.0.0.1）-port（例：1234）-tls（0:不加密，1:加密）
### 3.2运行client:
       运行client：./h_client -address（例：127.0.0.1:1234）-userName（例：zhangsan）-passeord（例：123456）-tls（0:不加密，1:加密）-way（1:get 0:post）

## 四、grpc client/server，实现一个基本的grpc调用请求处理（RpcServer放在server文件下，RpcClient放在client文件下）

### 4.1运行server：
        ./rpc_server -address（例：127.0.0.1:1234） -tls（0:不加密，1:加密）
### 4.2运行client: 
        ./rpc_client -address（例：127.0.0.1:1234）-tls（0:不加密，1:加密）

## 注：在运行tls加密时需要生成server.pem 和 server.key
### 1.生成服务器端的私钥
        openssl genrsa -out server.key 2048
### 2.生成服务器端证书
        openssl req -new -x509 -key server.key -out server.pem -days 3650
        or
        go run $GOROOT/src/crypto/tls/generate_cert.go --host localhost
