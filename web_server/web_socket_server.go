package main

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/websocket"
	"flag"
)

const WEBTRUE = "0"

func Echo(ws *websocket.Conn) {
	var err error

	for {
		var reply string

		if err = websocket.Message.Receive(ws, &reply); err != nil {
			fmt.Println("Can't receive")
			break
		}

		fmt.Println("Received back from client: " + reply)

		msg := "Received:  " + reply
		fmt.Println("Sending to client: " + msg)

		if err = websocket.Message.Send(ws, msg); err != nil {
			fmt.Println("Can't send")
			break
		}
	}
}

func main() {
	port := flag.String("port", "请输入端口号：例：1234", "")
	isTls := flag.String("tls", "请选择是否加密：0：不加密，1：加密", "")
	flag.Parse()
	http.Handle("/", websocket.Handler(Echo))
	if *isTls == WEBTRUE {
		wPort := ":" + *port
		if err := http.ListenAndServe(wPort, nil); err != nil {
			log.Fatal("ListenAndServe:", err)
			return
		}
	} else {
		wPort := ":" + *port
		if err := http.ListenAndServeTLS(wPort, "/Users/ant_oliu/go/1.8/src/GoTest/server.pem",
			"/Users/ant_oliu/go/1.8/src/GoTest/server.key", nil); err != nil {
			log.Fatal("ListenAndServe:", err)
			return
		}
	}
}
