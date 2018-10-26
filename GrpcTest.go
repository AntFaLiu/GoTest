package main

import (
	"net/rpc"
	"net/http"
	"log"
	"net"
	"time"
)

type Arg struct {
	A, B int
}
type Arit int
func (t *Arit) Multiply(args *Arg, reply *([]string)) error {   //乘法
	*reply = append(*reply, "test")
	return nil
}
func main() {
	arith := new(Arit)
	rpc.Register(arith)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	go http.Serve(l, nil)   //这里实现一个rpc服务器


	time.Sleep(5 * time.Second)
	client, err := rpc.DialHTTP("tcp", "127.0.0.1" + ":1234")  //创建一个客户端和服务器端建立连接
	if err != nil {
		log.Fatal("dialing:", err)
	}
	args := &Arg{7,8}
	reply := make([]string, 10)
	err = client.Call("Arit.Multiply", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	log.Println(reply)
}