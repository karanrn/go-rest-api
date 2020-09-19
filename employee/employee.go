package employee

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/karanrn/go-rest-api/database"
)

// Employee type holds information about an employee
type Employee struct {
	ID        int64  `json:"id,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Age       int32  `json:"age,omitempty"`
}

// GetEmployees pulls/gets all the employees
func GetEmployees(w http.ResponseWriter, r *http.Request) {
	res := []Employee{}
	db := database.DBConn()
	defer db.Close()

	selectStmt, err := db.Query("SELECT id, first_name, last_name, age FROM employee ORDER BY id")
	if err != nil {
		panic(err.Error())
	}

	emp := Employee{}
	for selectStmt.Next() {
		err = selectStmt.Scan(&emp.ID, &emp.FirstName, &emp.LastName, &emp.Age)
		if err != nil {
			panic(err.Error())
		}
		res = append(res, emp)
	}
	json.NewEncoder(w).Encode(res)
}

// GetEmployee pulls/gets particular employee basis Id
func GetEmployee(w http.ResponseWriter, r *http.Request) {
	empID := mux.Vars(r)["id"]

	db := database.DBConn()
	defer db.Close()

	selectStmt, err := db.Query("SELECT id, first_name, last_name, age FROM employee WHERE id=?", empID)
	if err != nil {
		panic(err.Error())
	}

	emp := Employee{}
	for selectStmt.Next() {
		err = selectStmt.Scan(&emp.ID, &emp.FirstName, &emp.LastName, &emp.Age)
		if err != nil {
			panic(err.Error())
		}
	}
	json.NewEncoder(w).Encode(emp)
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

	// Add employee to database
	db := database.DBConn()
	defer db.Close()

	insertStmt, err := db.Prepare("INSERT INTO employee (id, first_name, last_name, age) VALUES (?, ?, ?, ?)")
	if err != nil {
		panic(err.Error())
	}
	_, err = insertStmt.Exec(emp.ID, emp.FirstName, emp.LastName, emp.Age)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(`{'error': 'Internal Error'}`)
		panic(err.Error())
	}

	json.NewEncoder(w).Encode(emp)
}

// DeleteEmployee removes employee basis emp_id
func DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	empID := mux.Vars(r)["id"]

	db := database.DBConn()
	defer db.Close()

	delStmt, err := db.Prepare("DELETE FROM employee WHERE id=?")
	if err != nil {
		panic(err.Error())
	}

	_, err = delStmt.Exec(empID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(`{'error': 'Internal Error'}`)
		panic(err.Error())
	}

	json.NewEncoder(w).Encode(fmt.Sprintf(`{'message': %v empolyee deleted}`, empID))
}
