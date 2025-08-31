package main

import "fmt"

func main() {
	age := 18

	if age >= 18 {
		fmt.Println("You are an adult.")
	} else if age >= 13 && age < 18 {
		fmt.Println("You are a Teenager.")
	} else {
		fmt.Println("You are a Child.")
	}

	if gender := "male"; gender == "male" {
		fmt.Println("You are a boy.")
	} else {
		fmt.Println("You are a girl.")
	}
}
