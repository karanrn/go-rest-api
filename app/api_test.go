package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"

	emp "github.com/karanrn/go-rest-api/app/employee"
)

func TestGetEmployees(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/employees", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(emp.GetEmployees)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `[{"id":1,"first_name":"Karan","last_name":"Nadagoudar","age":24},{"id":2,"first_name":"Karan","last_name":"Nadagoudar","age":25}]` + "\n"

	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %q want %q",
			rr.Body.String(), expected)
	}
}

func TestGetEmployee(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/api/employees/{id:[0-9]+}", emp.GetEmployee)

	req, err := http.NewRequest("GET", "/api/employees/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"id":1,"first_name":"Karan","last_name":"Nadagoudar","age":24}` + "\n"

	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %q want %q",
			rr.Body.String(), expected)
	}
}
