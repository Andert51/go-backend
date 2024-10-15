package main

import (
	"go-backend/app/routes"
	"go-backend/config"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	config.InitializeFirebaseApp()
	router := mux.NewRouter()
	routes.InitializeRoutes(router)
	log.Println("Server Running in port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
