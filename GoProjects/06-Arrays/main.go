package main

import "fmt"

func main() {
	var list [4]string
	list[0] = "Hello"
	list[1] = "World"
	list[3] = "Language"
	fmt.Println("List values are", list)
	fmt.Println("List length is", len(list))

	var anotherList = [4]string{"Hello", "World", "Go", "Language"}
	fmt.Println("Another List values are", anotherList)
	fmt.Println("Another List length is", len(anotherList))
}
