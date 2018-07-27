package main

import (
	"fmt"
	"os"
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

var port string
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
	port = ":" + os.Args[1]
	isTls := os.Args[2]
	http.Handle("/", websocket.Handler(Echo))
	if isTls == WEBTRUE {
		if err := http.ListenAndServe(port, nil); err != nil {
			log.Fatal("ListenAndServe:", err)
			return
		}
	} else {
		if err := http.ListenAndServeTLS(port, "/Users/ant_oliu/go/1.8/src/GoTest/server.pem",
			"/Users/ant_oliu/go/1.8/src/GoTest/server.key", nil); err != nil {
			log.Fatal("ListenAndServe:", err)
			return
		}
	}
}


