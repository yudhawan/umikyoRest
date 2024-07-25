package api

import (
	"context"
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
	if r.Method != "GET" {
		libs.SendErrorResponse(w, "", 405)
		return
	}

	collection := libs.DBCollection("Users")
	if collection == nil {
		libs.SendErrorResponse(w, "Database collection not found", http.StatusInternalServerError)
		return
	}

	cursor, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		libs.SendErrorResponse(w, "Error Database Collection Users", 502)
		return
	}
	defer cursor.Close(context.Background())

	var users []bson.M

	if err = cursor.All(context.Background(), &users); err != nil {
		libs.SendErrorResponse(w, "Error decoding users", http.StatusInternalServerError)
		return
	}

	response := ResultResponse{
		Length: len(users),
		Page:   1,
		Offset: 0,
		Users:  users,
	}
	// fmt.Printf("all %v", response)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		libs.SendErrorResponse(w, "Error when encoding response", 502)
		return
	}
}
