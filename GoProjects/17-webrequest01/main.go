package main

import (
	"fmt"
	"io"
	"net/http"
)

const URL string = "https://jsonplaceholder.typicode.com/todos/1"

func main() {
	response, err := http.Get(URL)
	if err != nil || response.StatusCode != http.StatusOK {
		panic(err)
	}
	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("Response Body:", string(bodyBytes))
	defer response.Body.Close()
}
