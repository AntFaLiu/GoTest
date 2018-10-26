package main

import (
	"crypto/tls"
	"log"
)

func main() {
	conf := &tls.Config{
		InsecureSkipVerify: true,   //用来控制服务器主机名是否和证书和服务器主机名  如果设置为true则不会效验证书以及证书中的主机名和服务器主机名是否一致
	}
	conn, err := tls.Dial("tcp", "127.0.0.1:2345", conf)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()
	n, err := conn.Write([]byte("hello\n"))
	if err != nil {
		log.Println(n, err)
		return
	}
	buf := make([]byte, 100)
	n, err = conn.Read(buf)
	if err != nil {
		log.Println(n, err)
		return
	}
	println(string(buf[:n]))
}
