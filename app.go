package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/karanrn/go-rest-api/models"
)

var emps []models.Employee

// Homepage of the REST API
func Homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Homepage")
}

// GetEmployees pulls/gets all the employees
func GetEmployees(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(emps)
}

// GetEmployee pulls/gets particular employee basis Id
func GetEmployee(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, emp := range emps {
		if emp.ID == params["id"] {
			json.NewEncoder(w).Encode(emp)
			return
		}
	}
	json.NewEncoder(w).Encode(&models.Employee{})
}

// AddEmployee adds employee to the database
func AddEmployee(w http.ResponseWriter, r *http.Request) {
	var emp models.Employee
	// Using decoder because we are reading from the HTTP Stream
	err := json.NewDecoder(r.Body).Decode(&emp)
	if err != nil {
		// Handle error
		json.NewEncoder(w).Encode(`{'error': 'Error in decoding JSON'}`)
		return
	}
	emps = append(emps, emp)
	json.NewEncoder(w).Encode(emp)
}

func main() {
	router := mux.NewRouter().PathPrefix("/api").Subrouter().StrictSlash(false)

	emps = append(emps, models.Employee{ID: "1", FirstName: "Karan", LastName: "Nadagoudar", Age: 25})
	emps = append(emps, models.Employee{ID: "2", FirstName: "John", LastName: "Wick", Age: 25})

	// Routes for the employee resource
	router.HandleFunc("/", Homepage).Methods("GET")
	router.HandleFunc("/employees", GetEmployees).Methods("GET")
	router.HandleFunc("/employees/{id}", GetEmployee).Methods("GET")
	router.HandleFunc("/employees", AddEmployee).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", router))
}
