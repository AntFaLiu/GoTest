package main

import (
	"golang.org/x/net/websocket"
	"net/url"
	"log"
	"fmt"
	"crypto/tls"
)

func main() {
	webSocketNoTls()
	//webSocketTls()
}
func webSocketNoTls()  {
	Host := "localhost:1234"
	Path := ""
	u := url.URL{Scheme: "ws", Host: Host, Path: Path}
	log.Println(u.String())
	ws, err := websocket.Dial(u.String(), "", "http://"+Host)
	log.Println("http://"+Host)
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
func webSocketTls() {
	//tr := &http.Transport{
	//	TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	//}
	//client := &http.Client{Transport: tr}
	//websocket.DialConfig()

	webTls := &tls.Config{InsecureSkipVerify: true}

	Host := "localhost:12344"
	Path := ""
	u := url.URL{Scheme: "wss", Host: Host, Path: Path}
	origin := url.URL{Scheme:"http",Host: Host}
	log.Println("U:    ",u.String())
	log.Println("origin:    ",origin.String())
	//ws, err := websocket.Dial(u.String(), "", "http://"+Host)
	config := websocket.Config{
		Location : &u,
		Origin :   &origin,
		//,
		TlsConfig: webTls,
	}
	wss, err :=websocket.DialConfig(&config)
	//defer ws.Close() //关闭连接
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
