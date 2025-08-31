package main

import "fmt"

func main() {
	user := User{"John Doe", "john@example.com", 30, 1234567890}
	fmt.Println("User details are", user)
	fmt.Printf("User Details are %+v\n", user)
	fmt.Println("User name is", user.Name)
	fmt.Println("User mobile is", user.mobile) // Accessible within the same package
	email, status := user.GetEmail()
	fmt.Println("User email is", email)
	fmt.Println("Email update status is", status)
	fmt.Println("User details after GetEmail call are", user)
}

type User struct {
	Name   string
	Email  string
	Age    int
	mobile uint64
}

func (u User) GetEmail() (string, string) {
	u.Email = "updated@example.com"
	fmt.Println("Email updated to:", u.Email)
	return u.Email, "Updated"
}
