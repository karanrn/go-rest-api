package authentication

import (
	"fmt"
	"net/http"
)

// AuthMiddleware implements authentication/authorization
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") == "JOHN_SNOW" {
			next.ServeHTTP(w, r)
		} else {
			fmt.Println("Missing/Wrong token")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Missing/Wrong Token"))
		}
	})
}
