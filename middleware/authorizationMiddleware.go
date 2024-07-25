package middleware

import (
	"encoding/json"
	"net/http"
)

func AuthorizationMiddleware(next http.Handler) http.Handler {
	header := true
	if header {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			next.ServeHTTP(w, r)
		})
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode("Authorization Failed")
	})

}
