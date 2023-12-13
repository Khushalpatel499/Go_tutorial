package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {

	fmt.Println("Welcome to webreqverbs")
	PerformGetRequest()
	PerformPostJsonRequest()
}

func PerformGetRequest() {
	const myurl = "http://localhost:8000/get"
	response, err := http.Get(myurl)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	fmt.Println("Status Code:", response.StatusCode)
	fmt.Println("Content Length:", response.ContentLength)

	var responseString strings.Builder
	content, err := ioutil.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)
	if err != nil {
		panic(err)
	}
	fmt.Println("ByteCount is:", byteCount)
	fmt.Println(responseString.String())
	//fmt.Println(content)
	//fmt.Println(string(content))
}
func PerformPostJsonRequest() {
	const myurl = "http://localhost:8000/post"
	//fake json payload

	requestBody := strings.NewReader(`
	    {
			"coursename":"lets go lang",
			 "price":0,
			 "platform":"learncodeonline.in"
		}
	`)
	response, _ := http.Post(myurl, "application/json", requestBody)
	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}
func PerformPostFormRequest() {
	const myurl = "http://localhost:8000/postform"
	//form data

	data := url.Values{}
	data.Add("FirstName", "Khushal")
	data.Add("LastName", "Patel")
	data.Add("Email", "Khushal@go.dev")

	response, _ := http.PostForm(myurl, data)
	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))

}
