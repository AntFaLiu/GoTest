package main

import (
	"net/url"
	"log"
	"fmt"
	"crypto/tls"

	"golang.org/x/net/websocket"
	"flag"
)

const ISWEBTLS = "1"

func main() {
	address := flag.String("address", "请输入主机名：例：127.0.0.1:8080", "")
	isTls := flag.String("tls","请选择是否加密：0：不加密，1：加密","")
	flag.Parse()
	if *isTls == ISWEBTLS {
		webSocketTls(*address)
	} else {
		webSocketNoTls(*address)
	}
}
func webSocketNoTls(address string) {
	u := url.URL{Scheme: "ws", Host: address, Path: ""}
	log.Println(u.String())
	ws, err := websocket.Dial(u.String(), "", "http://"+address)
	log.Println("http://" + address)
	defer ws.Close()
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
		log.Println(err)
		return
	}
	cfg.Protocol = []string{""}
	cfg.TlsConfig = &tls.Config{InsecureSkipVerify: true}
	wss, err := websocket.DialConfig(cfg)
	defer wss.Close()
	if err != nil {
		log.Println("wss.Dial", err)
		return
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
