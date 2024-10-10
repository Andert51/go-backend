package controllers

import (
	"encoding/json"
	"go-backend/app/models"
	"net/http"
)

type userController struct {
	Service *services.userService
}

func (c *userController) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)
	err := c.Service.CreateUser(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
