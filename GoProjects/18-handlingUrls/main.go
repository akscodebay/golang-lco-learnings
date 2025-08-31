package main

import (
	"fmt"
	"net/url"
)

const urlString string = "https://jsonplaceholder.typicode.com:2000/todos/1?name=aks&age=20"

func main() {
	//read different parts of the URL
	parsedUrl, err := url.Parse(urlString)
	if err != nil {
		panic(err)
	}
	// Print the different parts of the URL
	fmt.Println("Scheme:", parsedUrl.Scheme)
	fmt.Println("Host:", parsedUrl.Host)
	fmt.Println("Path:", parsedUrl.Path)
	fmt.Println("Port:", parsedUrl.Port())
	fmt.Println("Query:", parsedUrl.RawQuery)

	// Get query parameters
	queryParams := parsedUrl.Query()
	for key, values := range queryParams {
		fmt.Println("Query Parameter:", key, "Values:", values)
	}

	//constructing url
	newUrl := &url.URL{
		Scheme:   "https",
		Host:     "jsonplaceholder.typicode.com",
		Path:     "/todos/1",
		RawQuery: "name=aks&age=20",
	}
	fmt.Println("Constructed URL:", newUrl.String())
}
