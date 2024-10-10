package controllers

import (
	"encoding/json"
	"go-backend/app/models"
	"go-backend/app/services"
	"net/http"
)

type UserController struct {
	Service *services.UserService
}

func (c *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)
	err := c.Service.CreateUser(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
