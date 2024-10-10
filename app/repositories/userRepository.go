package repositories

import (
	"context"
	"go-backend/app/models"

	"cloud.google.com/go/firestore"
)

type UserRepository struct {
	Client *firestore.Client
}

func (r *UserRepository) CreateUser(user models.User) error {
	_, _, err := r.Client.Collection("users").Add(context.Background(), user)
	return err
}
