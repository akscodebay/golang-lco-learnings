package main

import (
	"fmt"
	"sort"
)

func main() {
	var oneList []string
	oneList = append(oneList, "one")
	oneList = append(oneList, "two")
	oneList = append(oneList, "three")
	oneList = append(oneList, "four")
	fmt.Println("One List values are", oneList)

	oneList = oneList[:3]
	fmt.Println("One List values after removal are", oneList)

	secondList := make([]string, 2)
	secondList[0] = "five"
	secondList[1] = "six"
	secondList = append(secondList, "seven")
	secondList[2] = "eight"
	fmt.Println("Second List values are", secondList)
	fmt.Println("Second List length is", len(secondList))
	sort.Strings(secondList)
	fmt.Println("Second List values after sorting are", secondList)

	secondList = append(secondList[:1], secondList[2:]...)
	fmt.Println("Second List values after removal are", secondList)
}
