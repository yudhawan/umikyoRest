package middleware

import (
	"encoding/json"
	"net/http"
)

func AuthorizationMiddleware(next http.HandlerFunc) http.HandlerFunc {
	header := true
	if header {
		return func(w http.ResponseWriter, r *http.Request) {
			next(w, r)
		}
	}
	return func(w http.ResponseWriter, r *http.Request) {
		// libs.SendErrorResponse(w,"Not Allowed to Request",Sta)
		json.NewEncoder(w).Encode("Not ALlowed to Request")
	}
}
