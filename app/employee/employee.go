package employee

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/karanrn/go-rest-api/app/database"
)

// Employee type holds information about an employee
type Employee struct {
	ID        int64  `json:"id,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Age       int64  `json:"age,omitempty"`
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

	w.WriteHeader(http.StatusOK)
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

	w.WriteHeader(http.StatusOK)
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

	w.WriteHeader(http.StatusCreated)
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

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(fmt.Sprintf(`{'message': %v empolyee deleted}`, empID))
}

// LoadCSV adds employee details from the uploaded file
func LoadCSV(w http.ResponseWriter, r *http.Request) {
	// Parse multipart upload, maximum of 10 MB file
	// Left shift 20 for MB
	r.ParseMultipartForm(5 << 20)

	file, handler, err := r.FormFile("EmployeesFile")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(`{'error': 'Error Retrieving the file.'}`)
		panic(err.Error())
	}
	defer file.Close()

	fmt.Printf("Uploaded File: %+v\n", handler.Filename)

	// Create temporary file in tempdir
	tempFile, err := ioutil.TempFile("tempdir", "upload-*.csv")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(`{'error': 'Internal error, Upload the file again.'}`)
		panic(err.Error())
	}
	defer tempFile.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(`{'error': 'Internal error, Upload the file again.'}`)
		panic(err.Error())
	}
	tempFile.Write(fileBytes)

	// Read CSV and add employees
	csvFile, err := os.Open(tempFile.Name())
	defer csvFile.Close()
	if err != nil {
		fmt.Printf("Error opening file %v : %v\n", tempFile.Name(), err)
	}

	csvReader := csv.NewReader(csvFile)
	csvLines, err := csvReader.ReadAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(`{'error': 'Error reading CSV file, Upload the file again.'}`)
		panic(err.Error())
	}

	var data []Employee
	for _, line := range csvLines[1:] {
		id, _ := strconv.ParseInt(line[0], 10, 64)
		age, _ := strconv.ParseInt(line[3], 10, 64)

		row := Employee{
			ID:        id,
			FirstName: line[1],
			LastName:  line[2],
			Age:       age,
		}

		data = append(data, row)
	}

	// Prepare sql
	db := database.DBConn()
	defer db.Close()

	bulkInsertStmt, err := db.Prepare("INSERT INTO employee (id, first_name, last_name, age) VALUES (?, ?, ?, ?)")
	if err != nil {
		panic(err.Error())
	}

	for _, row := range data {
		// Implement Upsert
		bulkInsertStmt.Exec(row.ID, row.FirstName, row.LastName, row.Age)
	}

	// Remove temp file
	err = os.Remove(tempFile.Name())
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Fprintf(w, "File upload successful")
}
