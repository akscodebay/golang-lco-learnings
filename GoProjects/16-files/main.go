package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Create("./test.txt")
	checkNilErr(err)

	writeToFile(file, "This is a test insert to the test.txt from go lang")

	readFromFile("./test.txt")
	checkNilErr(err)

	defer file.Close()

}

func writeToFile(file *os.File, content string) {
	io.WriteString(file, content)
}

func readFromFile(filename string) {
	data, err := os.ReadFile(filename)
	checkNilErr(err)
	fmt.Println("File content in bytes:", data)
	fmt.Println("File content:", string(data))
}

func checkNilErr(err error) {
	if err != nil {
		fmt.Println("Error:", err)
		panic(err)
	}
}
