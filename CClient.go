package main

import (
	"bufio"
	"fmt"
	"net"
	"time"
	"crypto/tls"
	"log"
	"os"
)

var (
	quitSemaphore           chan bool
	message, cAddress, cTls string
)

const (
	TRUE = "1"
)

func main() {
	cAddress = os.Args[1]

	log.Println(cAddress)
	cTls = os.Args[2]
	if cTls == TRUE {
		conf := &tls.Config{
			InsecureSkipVerify: true, //用来控制服务器主机名是否和证书和服务器主机名  如果设置为true则不会效验证书以及证书中的主机名和服务器主机名是否一致
		}
		conn, err := tls.Dial("tcp", cAddress, conf)
		if err != nil {
			log.Println(err)
			return
		}
		defer conn.Close()
		fmt.Scanln(&message)
		b := []byte(message)
		conn.Write(b) //发送消息
		go messageHandeler(conn)
		<-quitSemaphore
	} else {
		var tcpAddr *net.TCPAddr
		tcpAddr, _ = net.ResolveTCPAddr("tcp", cAddress)
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

func onMessageRecived(conn *net.TCPConn) { //接收消息
	reader := bufio.NewReader(conn) //读取数据
	for {
		data := make([]byte, 128)
		total, err := reader.Read(data)
		if err != nil {
			quitSemaphore <- true //如果读取数据失败就发送信号通知不要再发送了
			break
		}
		fmt.Println("receive from server:  ", string(data[:total]))
		time.Sleep(time.Second)
		conn.Write(data[:total])
	}
}

func messageHandeler(conn net.Conn) { //接收消息
	reader := bufio.NewReader(conn) //读取数据
	for {
		data := make([]byte, 128)
		total, err := reader.Read(data)
		if err != nil {
			quitSemaphore <- true //如果读取数据失败就发送信号通知不要再发送了
			break
		}
		fmt.Println("receive from server:  ", string(data[:total]))
		time.Sleep(time.Second)
		conn.Write(data[:total])
	}
}
