package repositories

import (
	"context"
	"go-backend/app/models"

	"cloud.google.com/go/firestore"
)

type userRepository struct {
	Client *firestore.Client
}

func (r *userRepository) CreateUser(user models.User) error {
	_, _, err := r.Client.Collection("users").Add(context.Background(), user)
	return err
}
