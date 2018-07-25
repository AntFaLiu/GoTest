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

const HTTPTRUE = "1"

func main() {
	//clientTls()
	address := os.Args[1]
	userName := os.Args[2]
	password := os.Args[3]
	isTls := os.Args[4]
	httpGet(address, userName, password, isTls)
	httpPost(address, userName, password, isTls)

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
		}
		defer response.Body.Close()
		body, _ := ioutil.ReadAll(response.Body)
		log.Println(string(body))
	} else {
		response, err := http.Get("http://" + address + "/hello?userName=" + userName + "&password=" + password)
		if err != nil {
			log.Println(err)
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
		}

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
			// handle error
		}

		log.Println(string(body))
	} else {
		resp, err := http.Post("http://"+address+"/index",
			"application/x-www-form-urlencoded",
			strings.NewReader(string(jsonStr)))
		if err != nil {
			log.Println(err)
		}

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
			// handle error
		}

		log.Println(string(body))
	}

}

//func clientTls()  {
//	creds, err := credentials.NewClientTLSFromFile("/Users/ant_oliu/go/1.8/src/LypTest/server.pem",
//		"lyp")
//
//	//caCert, err := ioutil.ReadFile("/Users/ant_oliu/go/1.8/src/LypTest/server.key") //添加服务端的证书
//	//if err != nil {
//	//	log.Fatal(err)
//	//}
//	caCertPool := x509.NewCertPool()
//	caCertPool.AppendCertsFromPEM([]byte(creds))
//
//	client := &http.Client{
//		Transport: &http.Transport{
//			TLSClientConfig: &tls.Config{
//				RootCAs:      caCertPool, //添加认证
//			},
//		},
//	}
//
//	resp, err := client.Get("https://localhost:8080/hello?userName=zhangsan&password=123456")
//	if err != nil {
//		log.Println(err)
//		return
//	}
//
//	htmlData, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		log.Println(err)
//		return
//	}
//	defer resp.Body.Close()
//	log.Printf("%v\n", resp.Status)
//	log.Printf(string(htmlData))
//}
