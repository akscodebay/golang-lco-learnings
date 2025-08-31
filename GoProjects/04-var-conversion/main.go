package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Enter a number:")
	reader := bufio.NewReader(os.Stdin)
	numStr, _ := reader.ReadString('\n')
	num, err := strconv.ParseFloat(strings.TrimSpace(numStr), 64)
	if err != nil {
		fmt.Println("Error in reader input" + err.Error())
	} else {
		fmt.Println("Increamented num:", num+1)
	}

}
