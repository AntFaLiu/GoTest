package main

import (
	"net"
	"fmt"
	"bufio"
	"crypto/tls"
	"log"
	"flag"
)

var (
	ConMap map[string]net.Conn
)

const (
	STRUE = "1"
)

func main() {
	Port := flag.String("port", "请输入端口号：例：1234", "")
	cSTls := flag.String("tls", "请选择是否加密：0：不加密，1：加密", "")
	flag.Parse()
	//log.Println(cSPort)
	if *cSTls == STRUE {
		cSPort := ":" + *Port
		cert, err := tls.LoadX509KeyPair("/Users/ant_oliu/go/1.8/src/GoTest/server.pem",
			"/Users/ant_oliu/go/1.8/src/GoTest/server.key")
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
		cSPort := ":" + *Port
		log.Println("dfdsfdf",cSPort)
		var tcpAddr *net.TCPAddr
		var err error
		tcpAddr, err = net.ResolveTCPAddr("tcp", "127.0.0.1"+cSPort)
		if err != nil{
			log.Println(err)
		}
		tcpListener, err := net.ListenTCP("tcp", tcpAddr)
		if err != nil{
			log.Println("rerweew",err)
		}
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

func echo(conn net.Conn) {
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

func tcpPipe(conn *net.TCPConn) {
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
