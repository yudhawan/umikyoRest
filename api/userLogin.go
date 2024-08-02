package api

import (
	"context"
	"encoding/json"
	"net/http"
	"umikyoRest/libs"

	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

type UserBody struct {
	email    string
	password string
}
type UserHasLogin struct {
	user_data any
	token     string
}

func userLogin(w http.ResponseWriter, r *http.Request) {
	var user_body UserBody
	json.NewDecoder(r.Body).Decode(&user_body)
	collect := libs.Collection("Users", w, bson.D{{"email", user_body.email}})
	var userLogin UserBody
	var user_authenticated UserHasLogin
	user_authenticated.token = "tokeeennn"
	collect.All(context.Background(), &user_authenticated.user_data)
	if err := bcrypt.CompareHashAndPassword([]byte(userLogin.password), []byte(user_body.password)); err != nil {
		json.NewEncoder(w).Encode(user_authenticated)
		return
	}
	libs.SendErrorResponse(w, "Authentication Failed", http.StatusInternalServerError)

}
