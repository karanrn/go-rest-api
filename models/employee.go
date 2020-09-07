package models

// Employees type is array/list of Employee
type Employees struct {
	Employees []Employee
}

// Employee type holds information about an employee
type Employee struct {
	ID        string `json:"id,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Age       int    `json:"age,omitempty"`
}
