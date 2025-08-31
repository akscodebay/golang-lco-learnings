package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()
	fmt.Println("Current Time:", currentTime)
	formattedCurrentTime := currentTime.Format("01-02-2006 15:04:05 Monday")
	fmt.Println("Formatted Current Time:", formattedCurrentTime)
	date := time.Date(2025, time.August, 28, 0, 0, 0, 0, time.UTC)
	fmt.Println("Specific Date:", date)
}
