package main

import (
	"fmt"
	"net/http"
	"sync"
)

var waitGroup sync.WaitGroup
var mutex sync.Mutex
var completionStatus = make(map[string]string)

func main() {
	websites := []string{
		"https://www.google.com",
		"https://www.fb.org",
		"https://www.github.com",
		"https://www.twitter.com",
		"https://www.linkedin.com",
		"https://www.reddit.com",
		"https://www.pinterest.com",
		"https://www.tumblr.com",
		"https://www.quora.com",
		"https://www.irctc.co.in",
	}

	for _, website := range websites {
		go websiteStatus(website)
		waitGroup.Add(1)
	}
	waitGroup.Wait()
	fmt.Println("Completion Status:", completionStatus)
}

func websiteStatus(website string) {
	defer waitGroup.Done()
	resp, err := http.Get(website)
	if err != nil {
		fmt.Println(website, "is down")
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		mutex.Lock()
		completionStatus[website] = "up"
		mutex.Unlock()
		fmt.Println(website, "is up")
	} else {
		fmt.Println(website, "is down")
	}
}
