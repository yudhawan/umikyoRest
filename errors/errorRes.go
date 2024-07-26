package errors

import (
	"encoding/json"
	"net/http"
)

type authError struct {
	msg  string
	stat string
}

var responseEncode []authError

func ErrorAuthorization(w http.ResponseWriter, req *http.Request) {
	responseEncode := append(responseEncode, authError{msg: "Authorization Failed", stat: "Cannot Reach Anything"})
	w.WriteHeader(401)
	json.NewEncoder(w).Encode(responseEncode)
}

func ErrorAuthentication(text string, w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(404)
	json.NewEncoder(w).Encode(text)
}
