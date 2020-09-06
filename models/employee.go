package models

type Employees struct {
	Employees []Employee
}

type Employee struct {
	Id string `json:"id,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName string `json:"last_name,omitempty"`
	Age int `json:"age,omitempty"`
}