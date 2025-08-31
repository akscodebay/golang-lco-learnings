package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {

	//seeding Data
	employees = []Employee{
		{Id: 1, Name: "John Doe", Email: "john@example.com", Salary: 50000, Address: &Address{
			House:   "123",
			Street:  "Main St",
			City:    "New York",
			State:   "NY",
			Country: "USA",
			ZipCode: "10001",
		}, Mobile: "123-456-7890", Project: []string{"Project A", "Project B"}},
		{Id: 2, Name: "Jane Smith", Email: "jane@example.com", Salary: 60000, Address: &Address{
			House:   "456",
			Street:  "Second St",
			City:    "Los Angeles",
			State:   "CA",
			Country: "USA",
			ZipCode: "90001",
		}, Mobile: "987-654-3210", Project: []string{"Project C"}},
	}

	router := mux.NewRouter()
	router.HandleFunc("/", serveHome).Methods("GET")
	router.HandleFunc("/employees", getAllEmployees).Methods("GET")
	router.HandleFunc("/employees/{id}", getOneEmployee).Methods("GET")
	router.HandleFunc("/employees", createEmployee).Methods("POST")
	router.HandleFunc("/employees", updateEmployee).Methods("PUT")
	router.HandleFunc("/employees/{id}", deleteEmployee).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}

func checkEmpty(emp *Employee) bool {
	return emp == nil || (emp.Name == "" && emp.Email == "" && emp.Address == nil)
}

var employees []Employee

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to Employee Management System</h1>"))
}

func getAllEmployees(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(employees)
}

func getOneEmployee(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	fmt.Println("request params:", params)
	id, error := strconv.Atoi(params["id"])
	if error != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	for _, emp := range employees {
		if emp.Id == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(emp)
			return
		}
	}
	http.NotFound(w, r)
}

func createEmployee(w http.ResponseWriter, r *http.Request) {
	var emp Employee
	if r.Body == nil {
		http.Error(w, "Empty request body", http.StatusBadRequest)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&emp); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	if checkEmpty(&emp) {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	emp.Id = len(employees) + 1
	employees = append(employees, emp)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(emp)
}

func updateEmployee(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		http.Error(w, "Empty request body", http.StatusBadRequest)
		return
	}
	var emp Employee
	if err := json.NewDecoder(r.Body).Decode(&emp); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	if checkEmpty(&emp) {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	for i, e := range employees {
		if e.Id == emp.Id {
			employees[i] = emp
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(emp)
			return
		}
	}
	http.NotFound(w, r)
}

func deleteEmployee(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	convertedId, error := strconv.Atoi(id)
	if error != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	for i, e := range employees {
		if e.Id == convertedId {
			employees = append(employees[:i], employees[i+1:]...)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(e)
			return
		}
	}
	http.NotFound(w, r)
}

type Employee struct {
	Id      int      `json:"id"`
	Name    string   `json:"name"`
	Email   string   `json:"email"`
	Salary  int      `json:"-"`
	Address *Address `json:"address"`
	Mobile  string   `json:"mobile,omitempty"`
	Project []string `json:"project,omitempty"`
}

type Address struct {
	House   string `json:"house"`
	Street  string `json:"street"`
	City    string `json:"city"`
	State   string `json:"state"`
	Country string `json:"country"`
	ZipCode string `json:"zip_code"`
}
