package authentication

import (
	"fmt"
	"net/http"
)

// Authentication implements authentication/authorization
func Authentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		/*
			Please use better authentication/authorization strategy.
			Like JWT or OAuth
		*/
		if r.Header.Get("Authorization") == "JOHN_SNOW" {
			next.ServeHTTP(w, r)
		} else {
			fmt.Println("Missing/Wrong token")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Missing/Wrong Token"))
		}
	})
}

// Authorize allows if user is valid
func Authorize(token string) bool {
	if token == "JOHN_SNOW" {
		return true
	}
	return false
}
