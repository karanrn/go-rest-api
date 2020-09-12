package employee

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Employees type is array/list of Employee
type Employees struct {
	Employees []Employee
}

// Employee type holds information about an employee
type Employee struct {
	ID        int64  `json:"id,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Age       int32  `json:"age,omitempty"`
}

// Emps is slice of Employee type
var Emps []Employee

// GetEmployees pulls/gets all the employees
func GetEmployees(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Emps)
}

// GetEmployee pulls/gets particular employee basis Id
func GetEmployee(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.ParseInt(params["id"], 10, 64)
	for _, emp := range Emps {
		if emp.ID == id {
			json.NewEncoder(w).Encode(emp)
			return
		}
	}
	json.NewEncoder(w).Encode(&Employee{})
}

// AddEmployee adds employee to the database
func AddEmployee(w http.ResponseWriter, r *http.Request) {
	var emp Employee
	// Using decoder because we are reading from the HTTP Stream
	err := json.NewDecoder(r.Body).Decode(&emp)
	if err != nil {
		// Handle error
		json.NewEncoder(w).Encode(`{'error': 'Error in decoding JSON'}`)
		return
	}
	Emps = append(Emps, emp)
	json.NewEncoder(w).Encode(emp)
}
