package main

import (
	"log"
	"net/http"
	"encoding/json"

	"github.com/gorilla/mux"

	"github.com/karanrn/go-rest-api/models"
)


var emps []models.Employee

func GetEmployees(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(emps)
}

func GetEmployee(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, emp := range emps {
		if emp.Id == params["id"] {
			json.NewEncoder(w).Encode(emp)
			return
		}
	}
	json.NewEncoder(w).Encode(&models.Employee{})
}

func AddEmployee(w http.ResponseWriter, r *http.Request) {
	var emp models.Employee
	_ = json.NewDecoder(r.Body).Decode(&emp)
	emps = append(emps, emp)
	json.NewEncoder(w).Encode(emp)
}

func main() {
	router := mux.NewRouter()
	emps = append(emps, models.Employee{Id:"1", FirstName:"Karan", LastName:"Nadagoudar", Age:25})
	emps = append(emps, models.Employee{Id:"2", FirstName:"John", LastName:"Wick", Age:25})

	router.HandleFunc("/employees", GetEmployees).Methods("GET")
	router.HandleFunc("/employees/{id}", GetEmployee).Methods("GET")
	router.HandleFunc("/employees", AddEmployee).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", router))
}