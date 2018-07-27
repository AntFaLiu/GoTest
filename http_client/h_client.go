package main

import (
	"net/http"
	"io/ioutil"
	"strings"
	"log"
	"crypto/tls"
	"encoding/json"
	"os"
)

const (
	HTTPTRUE = "1"
	ISGET    = "1"
	ISPOST   = "0"
)

func main() {
	address := os.Args[1]
	userName := os.Args[2]
	password := os.Args[3]
	isTls := os.Args[4]
	getOrPost := os.Args[5]
	if getOrPost == ISGET {
		httpGet(address, userName, password, isTls)
	} else if getOrPost == ISPOST {
		httpPost(address, userName, password, isTls)
	}
}

func httpGet(address, userName, password, isTls string) {
	if isTls == HTTPTRUE {
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		client := &http.Client{Transport: tr}
		response, err := client.Get("https://" + address + "/hello?userName=" + userName + "&password=" + password)
		if err != nil {
			log.Println(err)
			return
		}
		defer response.Body.Close()
		body, _ := ioutil.ReadAll(response.Body)
		log.Println(string(body))
	} else {
		response, err := http.Get("http://" + address + "/hello?userName=" + userName + "&password=" + password)
		if err != nil {
			log.Println(err)
			return
		}
		defer response.Body.Close()
		body, _ := ioutil.ReadAll(response.Body)
		log.Println(string(body))
	}
}
func httpPost(address, userName, password, isTls string) {

	var user map[string]string
	user = make(map[string]string)
	user["userName"] = userName
	user["password"] = password
	jsonStr, err := json.Marshal(user)
	if err != nil {
		log.Println(err)
		panic(err)

	}
	if isTls == HTTPTRUE {
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		client := &http.Client{Transport: tr}
		resp, err := client.Post("https://"+address+"/index",
			"application/x-www-form-urlencoded",
			strings.NewReader(string(jsonStr)))
		if err != nil {
			log.Println(err)
			return
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println(err)
			return
		}
		log.Println(string(body))
	} else {
		resp, err := http.Post("http://"+address+"/index",
			"application/x-www-form-urlencoded",
			strings.NewReader(string(jsonStr)))
		if err != nil {
			log.Println(err)
			return
		}

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
			return
		}

		log.Println(string(body))
	}
}

