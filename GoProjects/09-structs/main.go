package main

import "fmt"

func main() {
	user := User{"John Doe", "john@example.com", 30}
	fmt.Println("User details are", user)
	fmt.Printf("User Details are %+v\n", user)
	fmt.Println("User name is", user.Name)
}

type User struct {
	Name  string
	Email string
	Age   int
}
