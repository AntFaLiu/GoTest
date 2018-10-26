package main

import (
	"log"
	"net"
	"bufio"
	"crypto/tls"
)

func main() {
	cert, err := tls.LoadX509KeyPair("/Users/ant_oliu/go/1.8/src/LypTest/server.pem", "/Users/ant_oliu/go/1.8/src/LypTest/server.key")
	//要创建一个server.pem 文件 存放秘钥
	if err != nil {
		log.Println(err)
		return
	}
	config := &tls.Config{Certificates: []tls.Certificate{cert}}
	ln, err := tls.Listen("tcp", ":2345", config)
	if err != nil {
		log.Println(err)
		return
	}
	defer ln.Close()
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConn(conn)
	}
}
func handleConn(conn net.Conn) {
	defer conn.Close()
	r := bufio.NewReader(conn)
	for {
		msg, err := r.ReadString('\n')
		if err != nil {
			log.Println(err)
			return
		}
		println(msg)
		n, err := conn.Write([]byte("world\n"))
		if err != nil {
			log.Println(n, err)
			return
		}
	}
}