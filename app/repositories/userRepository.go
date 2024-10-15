package repositories

import (
	"context"
	"go-backend/app/models"
	"go-backend/config"

	"google.golang.org/api/iterator"
)

type UserRepository struct {
}

func (r *UserRepository) CreateUser(user models.User) error {
	ctx := context.Background()
	client, err := config.FirebaseApp.Firestore(ctx)
	if err != nil {
		return err
	}
	defer client.Close()
	_, _, err = client.Collection("users_go").Add(ctx, user)
	return err
}

func (r *UserRepository) GetAllUsers() ([]models.User, error) {
	ctx := context.Background()
	client, err := config.FirebaseApp.Firestore(ctx)
	if err != nil {
		return nil, err
	}
	var users []models.User
	iter := client.Collection("users_go").Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		var user models.User
		doc.DataTo(&user)
		users = append(users, user)
	}
	return users, nil
}

func (r *UserRepository) GetUserById(id string) (*models.User, error) {
	ctx := context.Background()
	client, err := config.FirebaseApp.Firestore(ctx)
	if err != nil {
		return nil, err
	}
	defer client.Close()
	doc, err := client.Collection("users_go").Doc(id).Get(ctx)
	if err != nil {
		return nil, err
	}
	var user models.User
	doc.DataTo(&user)
	user.ID = doc.Ref.ID
	return &user, nil
}

func (r *UserRepository) GetUserByUsername(username string) (*models.User, error) {
	ctx := context.Background()
	client, err := config.FirebaseApp.Firestore(ctx)
	if err != nil {
		return nil, err
	}
	defer client.Close()
	iter := client.Collection("users_go").Where("Username", "==", username).Documents(ctx)
	doc, err := iter.Next()

	if err != iterator.Done {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	var user models.User
	doc.DataTo(&user)
	user.ID = doc.Ref.ID
	return &user, nil

}

func (r *UserRepository) GetUserByRole(rol string) ([]models.User, error) {
	ctx := context.Background()
	client, err := config.FirebaseApp.Firestore(ctx)
	if err != nil {
		return nil, err
	}
	var users []models.User
	iter := client.Collection("users_go").Where("Rol", "==", rol).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		var user models.User
		doc.DataTo(&user)
		users = append(users, user)
	}
	return users, nil
}

func (r *UserRepository) UpdateUser(id string, user models.User) error {
	ctx := context.Background()
	client, err := config.FirebaseApp.Firestore(ctx)
	if err != nil {
		return err
	}
	defer client.Close()
	_, err = client.Collection("users_go").Doc(id).Set(ctx, user)
	return err
}

func (r *UserRepository) DeleteUser(id string) error {
	ctx := context.Background()
	client, err := config.FirebaseApp.Firestore(ctx)
	if err != nil {
		return err
	}
	defer client.Close()
	_, err = client.Collection("users_go").Doc(id).Delete(ctx)
	return err
}
