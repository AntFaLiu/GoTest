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
			InsecureSkipVerify: true,
		}
		conn, err := tls.Dial("tcp", cAddress, conf)
		if err != nil {
			log.Println(err)
			return
		}
		defer conn.Close()
		fmt.Scanln(&message)
		b := []byte(message)
		conn.Write(b)
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
