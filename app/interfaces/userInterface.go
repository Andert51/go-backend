package interfaces

import "go-backend/app/models"

type IUser interface {
	CreateUser(user models.User) error
	GetAllUsers() ([]models.User, error)
	GetUserById(id string) (*models.User, error)
	GetUserByUsername(username string) (*models.User, error)
	GetUserByRol(rol string) ([]models.User, error)
	UpdateUser(id string, user models.User) error
	DeleteUser(id string) error
}
