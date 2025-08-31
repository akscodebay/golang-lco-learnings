package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	jsonFromWeb := []byte(`{
                "id": 2,
                "name": "John Doe",
                "email": "john.doe@example.com",
                "address": "456 Main St, Anytown, USA",
                "mobile": "555-1234"
        }`)
	var employee Employee
	checkvalid := json.Valid(jsonFromWeb)
	if checkvalid {
		fmt.Println("JSON is valid")
		json.Unmarshal(jsonFromWeb, &employee)
		fmt.Printf("%#v\n", employee)
	} else {
		fmt.Println("JSON is invalid")
	}

	//cases where struct is not present or not required
	var genericMap map[string]interface{}
	json.Unmarshal(jsonFromWeb, &genericMap)
	fmt.Printf("%#v\n", genericMap)

	for key, value := range genericMap {
		fmt.Printf("Key: %s, Value: %v\n", key, value)
	}
}

type Employee struct {
	Id      int      `json:"id"`
	Name    string   `json:"name"`
	Email   string   `json:"email"`
	Salary  int      `json:"-"`
	Address string   `json:"address"`
	Mobile  string   `json:"mobile,omitempty"`
	Project []string `json:"project,omitempty"`
}
