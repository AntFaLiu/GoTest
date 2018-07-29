package main

import (
	"bufio"
	"fmt"
	"net"
	"time"
	"crypto/tls"
	"log"
	"flag"
)

var (
	quitSemaphore chan bool
	message       string
)

const (
	TRUE = "1"
)

func main() {
	cAddress := flag.String("address", "请输入访问地址：例：127.0.0.1:1234", "")
	log.Println(*cAddress)
	cTls := flag.String("tls", "请选择是否加密：0：不加密，1：加密", "")
	flag.Parse()
	if *cTls == TRUE {
		conf := &tls.Config{
			InsecureSkipVerify: true,
		}
		conn, err := tls.Dial("tcp", *cAddress, conf)
		if err != nil {
			log.Println(err)
			return
		}
		defer conn.Close()
		fmt.Scanln(&message)
		b := []byte(message)
		log.Println("sd", message)
		conn.Write(b)
		go messageHandeler(conn)
		<-quitSemaphore
	} else {
		log.Println(*cAddress)
		var tcpAddr *net.TCPAddr
		var err error
		tcpAddr, err = net.ResolveTCPAddr("tcp", *cAddress)
		if err != nil {
			log.Println("ggjhghjh", err)
		}

		conn, _ := net.DialTCP("tcp", nil, tcpAddr)
		defer conn.Close()
		fmt.Println("connected!")
		fmt.Scanln(&message)
		b := []byte(message)
		conn.Write(b) //发送消息
		go onMessageRecived(conn)
		<-quitSemaphore
	}
}

//receive message
func onMessageRecived(conn *net.TCPConn) {
	reader := bufio.NewReader(conn)
	for {
		data := make([]byte, 128)
		total, err := reader.Read(data)
		if err != nil {
			quitSemaphore <- true
			break
		}
		fmt.Println("receive from server:  ", string(data[:total]))
		time.Sleep(time.Second)
		conn.Write(data[:total])
	}
}

//receive message
func messageHandeler(conn net.Conn) {
	reader := bufio.NewReader(conn)
	for {
		data := make([]byte, 128)
		total, err := reader.Read(data)
		if err != nil {
			quitSemaphore <- true
			break
		}
		fmt.Println("receive from server:  ", string(data[:total]))
		time.Sleep(time.Second)
		conn.Write(data[:total])
	}
}
