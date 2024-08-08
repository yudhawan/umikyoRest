package middleware

import (
	"net/http"
	"umikyoRest/libs"
)

func AuthorizationMiddleware(next http.Handler) http.Handler {
	header := true
	if header {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// fmt.Printf("isi headernya : %v", r.Header)
			next.ServeHTTP(w, r)
		})
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		libs.SendErrorResponse(w, "Authorizatio Failed", 500)
	})

}
