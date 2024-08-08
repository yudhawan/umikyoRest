package api

import (
	"encoding/json"
	"net/http"
	"umikyoRest/libs"

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
	userField := libs.Field{}
	userField.GetOne("email", user_body.email)
	// collect := libs.Collection("Users", w).FindOne(context.Background(), bson.D{{"email", user_body.email}})
	// collect := libs.Collection("Users", w, bson.D{{"email", user_body.email}})
	var userLogin UserBody
	token := "tokeeeen"
	// var user_authenticated UserHasLogin
	// user_authenticated.token = "tokeeennn"
	// collect.All(context.Background(), &user_authenticated.user_data)
	if err := bcrypt.CompareHashAndPassword([]byte(userLogin.password), []byte(user_body.password)); err != nil {
		user_authenticated := UserHasLogin{user_data: userField, token: token}
		json.NewEncoder(w).Encode(user_authenticated)
		return
	}
	libs.SendErrorResponse(w, "Authentication Failed", http.StatusInternalServerError)

}
