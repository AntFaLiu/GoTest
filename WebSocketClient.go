package main

import (
	"golang.org/x/net/websocket"
	"net/url"
	"log"
	"fmt"
	"crypto/tls"
	"os"
)

const ISWEBTLS = "1"

func main() {
	Host := os.Args[1]
	isTls := os.Args[2]
	if isTls == ISWEBTLS {
		webSocketTls(Host)
	} else {
		webSocketNoTls(Host)
	}
}
func webSocketNoTls(Host string) {
	u := url.URL{Scheme: "ws", Host: Host, Path: ""}
	log.Println(u.String())
	ws, err := websocket.Dial(u.String(), "", "http://"+Host)
	log.Println("http://" + Host)
	defer ws.Close() //关闭连接
	if err != nil {
		log.Println("ws.Dial", err)
	}
	for {
		var str string
		fmt.Scanln(&str)
		_, err = ws.Write([]byte(str))
		if err != nil {
			log.Println(err)
			return
		}
		b := make([]byte, 128)
		total, err := ws.Read(b)
		if err != nil {
			log.Println(err)
			return
		}
		log.Println(string(b[:total]))
	}
}
func webSocketTls(Host string) {
	log.Println("加密")
	url := "wss://" + Host + "/"
	origin := "https://" + "Host" + "/"
	var err error
	var cfg *websocket.Config
	if cfg, err = websocket.NewConfig(url, origin); err != nil {
		panic(err)
	}
	cfg.Protocol = []string{""}
	cfg.TlsConfig = &tls.Config{InsecureSkipVerify: true}
	wss, err := websocket.DialConfig(cfg)
	defer wss.Close() //关闭连接
	if err != nil {
		log.Println("wss.Dial", err)
	}
	for {
		var str string
		fmt.Scanln(&str)
		_, err = wss.Write([]byte(str))
		if err != nil {
			log.Println(err)
			return
		}
		b := make([]byte, 128)
		total, err := wss.Read(b)
		if err != nil {
			log.Println(err)
			return
		}
		log.Println(string(b[:total]))
	}
}
