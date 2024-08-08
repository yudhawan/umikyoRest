package api

import (
	"encoding/json"
	"net/http"
	"umikyoRest/libs"

	"go.mongodb.org/mongo-driver/bson"
)

type ResultResponse struct {
	Length int      `json:"length"`
	Page   int      `json:"page"`
	Offset int      `json:"offset"`
	Users  []bson.M `json:"users"`
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	data := libs.Field{FieldName: "Users"}
	users := data.GetAll()

	response := ResultResponse{
		Length: len(users),
		Page:   1,
		Offset: 0,
		Users:  users,
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		libs.SendErrorResponse(w, "Error when encoding response", 502)
		return
	}
}
