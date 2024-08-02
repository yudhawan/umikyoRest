package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func AuthorizationMiddleware(next http.Handler) http.Handler {
	header := true
	if header {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Printf("isi headernya : %v", r.Header)
			next.ServeHTTP(w, r)
		})
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode("Authorization Failed")
	})

}
