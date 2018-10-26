package main

import (
	"fmt"
	"net/http"
)

func main() {
	f()
	fmt.Println("djdkfjls")

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("fjdskfjdklsfj")
		}
	}()
	fmt.Println("jfoksdjf")
}

func f() (int, error) {
	err := http.ListenAndServeTLS(":", "/Users/ant_oliu/go/1.8/src/GoTest/server.pem",
		"/Users/ant_oliu/go/1.8/src/GoTest/server.ke", nil) //tls
	if err != nil{
		fmt.Println("ndnmsb")
	}
	return 1, err
}
