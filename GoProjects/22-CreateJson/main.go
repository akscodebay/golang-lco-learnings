package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	employees := []Employee{
		{
			ID:      1,
			Name:    "John Doe",
			Email:   "john.doe@example.com",
			Salary:  50000,
			Address: "123 Main St, Anytown, USA",
			Mobile:  "555-1234",
			Project: []string{"Project A", "Project B"},
		},
		{
			ID:      2,
			Name:    "John Doe",
			Email:   "john.doe@example.com",
			Salary:  40000,
			Address: "456 Main St, Anytown, USA",
			Mobile:  "555-1234",
			Project: nil,
		},
	}

	jsonData, err := json.MarshalIndent(employees, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(jsonData))

}

type Employee struct {
	ID      int      `json:"id"`
	Name    string   `json:"name"`
	Email   string   `json:"email"`
	Salary  int      `json:"-"`
	Address string   `json:"address"`
	Mobile  string   `json:"mobile,omitempty"`
	Project []string `json:"project,omitempty"`
}
