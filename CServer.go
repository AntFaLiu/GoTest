package main

import (
	"net"
	"fmt"
	"bufio"
	"crypto/tls"
	"log"
	"os"
)

var (
	ConMap        map[string]net.Conn //存放用户和用户名
	cSPort, cSTls string
)

const (
	STRUE = "1"
)

func main() {
	cSPort = ":" + os.Args[1]
	cSTls = os.Args[2]
	log.Println(cSPort)
	if cSTls == STRUE {
		cert, err := tls.LoadX509KeyPair("/Users/ant_oliu/go/1.8/src/GoTest/server.pem",
			"/Users/ant_oliu/go/1.8/src/GoTest/server.key")          //要创建一个server.pem 文件 存放秘钥
		if err != nil {
			log.Println(err)
			return
		}
		config := &tls.Config{Certificates: []tls.Certificate{cert}}
		ln, err := tls.Listen("tcp", cSPort, config)
		if err != nil {
			log.Println(err)
			return
		}
		ConMap = make(map[string]net.Conn)
		for {
			conn, err := ln.Accept()
			if err != nil {
				log.Println(err)
				continue
			}
			defer conn.Close()
			ConMap[conn.RemoteAddr().String()] = conn
			go echo(conn)
		}
	} else {
		var tcpAddr *net.TCPAddr
		tcpAddr, _ = net.ResolveTCPAddr("tcp", "127.0.0.1"+cSPort)
		tcpListener, _ := net.ListenTCP("tcp", tcpAddr)
		defer tcpListener.Close()
		for {
			tcpConn, err := tcpListener.AcceptTCP()
			if err != nil {
				continue
			}
			fmt.Println("A client connected : " + tcpConn.RemoteAddr().String())
			go tcpPipe(tcpConn)
		}
	}
}

func echo(conn net.Conn) { //接收数据
	ipStr := conn.RemoteAddr().String()
	defer func() {
		fmt.Println("disconnected :" + ipStr)
		conn.Close()
	}()
	reader := bufio.NewReader(conn)

	for {
		data := make([]byte, 128)
		total, err := reader.Read(data)
		if err != nil {
			return
		}
		fmt.Println("receive from client", string(data[:total]))
		conn.Write(data[:total])
	}
}

func tcpPipe(conn *net.TCPConn) { //接收数据
	ipStr := conn.RemoteAddr().String()
	defer func() {
		fmt.Println("disconnected :" + ipStr)
		conn.Close()
	}()
	reader := bufio.NewReader(conn)

	for {
		data := make([]byte, 128)
		total, err := reader.Read(data)
		if err != nil {
			return
		}
		fmt.Println("receive from client：  ", string(data[:total]))
		conn.Write(data[:total])
	}
}
