package routes

import (
	"encoding/json"
	"net/http"
)

func registerUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	response := map[string]string{"message": "Success terr"}
	// w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
