package main

import "fmt"

func main() {
	sum := add(3, 5)
	fmt.Println("sum is ", sum)

	result, _ := addMoreThanTwo(3, 5, 7, 8)
	fmt.Println("result is ", result)
}

func add(a int, b int) int {
	return a + b
}

func addMoreThanTwo(numbers ...int) (int, string) {
	sum := 0
	for _, n := range numbers {
		sum += n
	}
	return sum, "Success"
}
