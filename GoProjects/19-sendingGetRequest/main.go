package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

const url string = "https://jsonplaceholder.typicode.com/posts/1"

func main() {
	sendGetRequest(url)
}

func sendGetRequest(url string) {
	// Implementation for sending a GET request

	response, err := http.Get(url)
	if err != nil {
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
