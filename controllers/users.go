package controllers

import (
	"encoding/json"
	"go-todo-rest-api/models"
	u "go-todo-rest-api/utils"
	"net/http"
)

// CreateAccount is used to create a new user account
var CreateAccount = func(w http.ResponseWriter, r *http.Request) {

	user := &models.User{}
	err := json.NewDecoder(r.Body).Decode(user) //decode the request body into struct and failed if any error occur
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp := user.Create()
	u.Respond(w, resp)
}

// Authenticate will verify the provided email address and
// password are correct.
var Authenticate = func(w http.ResponseWriter, r *http.Request) {

	user := &models.User{}
	err := json.NewDecoder(r.Body).Decode(user) //decode the request body into struct and failed if any error occur
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp := models.Authenticate(user.Email, user.Password)
	u.Respond(w, resp)
}
