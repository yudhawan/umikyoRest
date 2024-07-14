package routes

import (
	"encoding/json"
	"net/http"
)

func empty(w http.ResponseWriter, r *http.Request) {
	response := "TErrrrrr"
	json.NewEncoder(w).Encode(response)
}
