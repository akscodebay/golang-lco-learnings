package main

import "fmt"

func main() {
	days := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

	fmt.Println("Print using for loop:")
	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	}

	fmt.Println("Print using for range loop index:")
	for index := range days {
		fmt.Println(days[index])
	}

	fmt.Println("Print using for range loop value:")
	for _, day := range days {
		fmt.Println(day)
	}

	fmt.Println("Print for while loop:")
	i := 0
	for i < len(days) {
		if days[i] == "Wednesday" {
			goto gotoLabel
		}
		if days[i] == "Thursday" {
			i++
			continue
		}
		if days[i] == "Saturday" {
			break
		}
		fmt.Println(days[i])
	gotoLabel:
		i++
	}

}
