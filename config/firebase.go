package config

// Go puede importar automaticamente los paquetes necesarios

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

var FirebaseApp *firebase.App

func InitializeFirebaseApp() {
	opt := option.WithCredentialsFile("./firebaseServiceAccount.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatal("Failed to Initialize Firebase App: %v ", err)
	}
	FirebaseApp = app
}

func GetAuthClient(app *firebase.App) *auth.Client {
	client, err := app.Auth(context.Background())
	if err != nil {
		log.Fatal("Failed to Initialize Firebase Auth: %v ", err)
	}
	return client
}
