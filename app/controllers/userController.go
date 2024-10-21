package controllers

import (
	"encoding/json"
	"go-backend/app/models"
	"go-backend/app/services"
	"net/http"

	"github.com/gorilla/mux"
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

func (c *UserController) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := c.Service.Repo.GetAllUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(users)

}

func (c *UserController) GetUserById(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	user, err := c.Service.GetUserById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(user)
}

func (c *UserController) GetUserByRol(w http.ResponseWriter, r *http.Request) {
	rol := mux.Vars(r)["rol"]
	user, err := c.Service.GetUserByRol(rol)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(user)
}

func (c *UserController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := c.Service.UpdateUser(id, user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (c *UserController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	if err := c.Service.DeleteUser(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
