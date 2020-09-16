package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"golang.org/x/time/rate"

	auth "github.com/karanrn/go-rest-api/authentication"
	emp "github.com/karanrn/go-rest-api/employee"
)

const (
	// PORT number for the web server
	PORT = ":8080"
)

// Homepage of the REST API
func Homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Homepage")
}

// Initalizing token bucket using rate
var limiters = map[string]*rate.Limiter{
	"Authorized":   rate.NewLimiter(2, 10),
	"Unauthorized": rate.NewLimiter(1, 3),
}

// Serves next request if token is available else rejects request
func rateLimitMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if auth.Authorize(r.Header.Get("Authorization")) {
			if !limiters["Authorized"].Allow() {
				http.Error(w, http.StatusText(429), http.StatusTooManyRequests)
				return
			}
		} else {
			if !limiters["Unauthorized"].Allow() {
				http.Error(w, http.StatusText(429), http.StatusTooManyRequests)
				return
			}
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	router := mux.NewRouter().PathPrefix("/api").Subrouter().StrictSlash(false)

	// Routes for the employee resource
	router.HandleFunc("/", Homepage).Methods("GET")
	router.HandleFunc("/employees", emp.GetEmployees).Methods("GET")
	router.HandleFunc("/employees/{id:[0-9]+}", emp.GetEmployee).Methods("GET")
	router.Handle("/employees", http.HandlerFunc(emp.AddEmployee)).Methods("POST")

	log.Fatal(http.ListenAndServe(PORT, rateLimitMiddleware(router)))
}
