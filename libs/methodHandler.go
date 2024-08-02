package libs

import (
	"net/http"
)

var method = []string{"GET", "POST", "PUT", "DELETE"}

func MethodHandler(m string, fn http.HandlerFunc) http.HandlerFunc {
	for _, val := range method {
		if val == m {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				fn.ServeHTTP(w, r)
			})
		}
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		SendErrorResponse(w, "Method is not Allowed", 405)
	})
}
