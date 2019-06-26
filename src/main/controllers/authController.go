package controllers

import (
	"net/http"
	"encoding/json"
	"main/models"
	u "main/utils"
)

var Hello = func (w http.ResponseWriter, r *http.Request) {
	u.Respond(w, u.Message(true, "Hello world"), http.StatusOK)
}

var Secured = func (w http.ResponseWriter, r *http.Request) {
	u.Respond(w, u.Message(true, "Secured area"), http.StatusOK)		
}

var Register = func (w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	e := json.NewDecoder(r.Body).Decode(user)
	if e != nil {
		u.Respond(w, u.Message(false, "Invalid request"), http.StatusBadRequest)
		return
	}

	res := user.Create()
	u.Respond(w, res, http.StatusCreated)
}

var Login = func (w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	e := json.NewDecoder(r.Body).Decode(user)
	if e != nil {
		u.Respond(w, u.Message(false, "Invalid request"), http.StatusBadRequest)
		return
	}

	res := models.Login(user.Email, user.Password)
	u.Respond(w, res, http.StatusOK)
}