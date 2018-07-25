package main

import (
	"net/http"
	"io/ioutil"
	"strings"
	"log"
	"crypto/tls"
	"encoding/json"
	"flag"
)

const (
	HTTPTRUE = "1"
	ISGET    = "1"
	ISPOST   = "0"
)

func main() {
	address := flag.String("address", "请输入访问地址：例：127.0.0.1:1234", "")
	userName := flag.String("userName", "请输入用户名：例：zhangsan", "")
	password := flag.String("password", "请输入密码：例：123456", "")
	isTls := flag.String("tls","请选择是否加密：0：不加密，1：加密","")
	getOrPost := flag.String("way","请选择请求方式：0：post，1：get","")
	flag.Parse()
	if *getOrPost == ISGET {
		httpGet(*address, *userName, *password, *isTls)
	} else if *getOrPost == ISPOST {
		httpPost(*address, *userName, *password, *isTls)
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
		log.Println("get tls：",string(body))
	} else {
		response, err := http.Get("http://" + address + "/hello?userName=" + userName + "&password=" + password)
		if err != nil {
			log.Println(err)
			return
		}
		defer response.Body.Close()
		body, _ := ioutil.ReadAll(response.Body)
		log.Println("get : ",string(body))
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
		log.Println("post tls :",string(body))
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
		log.Println("post  :"+string(body))
	}
}

