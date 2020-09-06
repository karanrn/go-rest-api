package models

type Employees struct {
	Employees []Employee
}

type Employee struct {
	Id string
	FirstName string
	LastName string
	Age int
}