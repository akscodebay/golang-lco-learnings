package main

import "fmt"

func main() {
	mymap := make(map[string]string)
	mymap["name"] = "John"
	mymap["age"] = "30"
	mymap["city"] = "New York"
	mymap["mobile"] = "123-456-7890"
	fmt.Println("Map values are:", mymap)
	fmt.Println("Value for Name key is", mymap["name"])
	delete(mymap, "mobile")

	for key, value := range mymap {
		fmt.Println("Key:", key, ", Value:", value)
	}
}
