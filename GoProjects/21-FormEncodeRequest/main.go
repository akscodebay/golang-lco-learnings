package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const urlString string = "https://jsonplaceholder.typicode.com/posts/1"

func main() {
	sendPostFormEncodedRequest(urlString)
}

func sendPostFormEncodedRequest(urlString string) {
	// Implementation for sending a POST request
	data := url.Values{}
	data.Add("title", "foo")
	data.Add("body", "bar")
	data.Add("userId", "1")
	response, err := http.PostForm(urlString, data)
	if err != nil || response.StatusCode != http.StatusOK {
		panic(err)
	}
	content, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	var responseString strings.Builder
	byteCount, err := responseString.Write(content)
	if err != nil {
		panic(err)
	}
	fmt.Println("Byte count:", byteCount)
	fmt.Println("Response:", responseString.String())
	defer response.Body.Close()
}
