package routes

import (
	"go-backend/app/controllers"
	"go-backend/app/repositories"
	"go-backend/app/services"
	"net/http"

	"github.com/gorilla/mux"
)

func InitializeRoutes(router *mux.Router) {
	repo := &repositories.UserRepository{}
	service := &services.UserService{Repo: repo}
	controller := &controllers.UserController{Service: service}

	router.HandleFunc("/users", controller.CreateUser).Methods(http.MethodPost)
	router.HandleFunc("/users", controller.GetAllUsers).Methods(http.MethodGet)
	router.HandleFunc("/users/{id}", controller.GetUserById).Methods(http.MethodGet)
	router.HandleFunc("/users/rol/{rol}", controller.GetUserByRol).Methods(http.MethodGet)
	router.HandleFunc("/users/{id}", controller.UpdateUser).Methods(http.MethodPut)
	router.HandleFunc("/users/{id}", controller.DeleteUser).Methods(http.MethodDelete)
}
