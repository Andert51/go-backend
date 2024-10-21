package services

import (
	"errors"
	"go-backend/app/models"
	"go-backend/app/repositories"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	Repo *repositories.UserRepository
}

// func (s *UserService) isUserDuplicated(user string) (bool, error) {
// 	users, err := s.Repo.GetAllUsers()
// 	if err != nil {
// 		return false, err
// 	}

// 	for _, u := range users {
// 		if u.Username == user {
// 			return true, nil
// 		}
// 	}

// 	return false, nil
// }

func (s *UserService) isNameDuplicated(name, patsurn, matsurn string) (bool, error) {
	users, err := s.Repo.GetAllUsers()
	if err != nil {
		return false, err
	}

	for _, u := range users {
		if u.Name == name && u.Patsurn == patsurn && u.Matsurn == matsurn {
			return true, nil
		}
	}

	return false, nil
}

func (s *UserService) CreateUser(user models.User) error {
	isDuplicated, err := s.Repo.GetUserByUsername(user.Username)
	if err != nil {
		return err
	}

	if isDuplicated != nil {
		return errors.New("user already exists")
	}

	isNameDuplicated, err := s.isNameDuplicated(user.Name, user.Patsurn, user.Matsurn)
	if err != nil {
		return err
	}

	if isNameDuplicated {
		return errors.New("name already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)

	return s.Repo.CreateUser(user)

}

func (s *UserService) UpdateUser(id string, user models.User) error {
	userId, err := s.Repo.GetUserById(id)
	if err != nil {
		return err
	}
	if userId == nil {
		return errors.New("user not found")
	}
	userId.Name = user.Name
	userId.Patsurn = user.Patsurn
	userId.Matsurn = user.Matsurn
	userId.Addres = user.Addres
	userId.Phone = user.Phone
	userId.City = user.City
	userId.State = user.State
	userId.Username = user.Username
	userId.Rol = user.Rol
	userId.Image = user.Image

	if user.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil
		}
		userId.Password = string(hashedPassword)
	}
	return s.Repo.UpdateUser(id, *userId)

}

func (s *UserService) DeleteUser(id string) error {
	userId, err := s.Repo.GetUserById(id)
	if err != nil {
		return err
	}
	if userId == nil {
		return errors.New("user not found")
	}
	return s.Repo.DeleteUser(id)
}

func (s *UserService) GetUserById(id string) (*models.User, error) {
	return s.Repo.GetUserById(id)
}

func (s *UserService) GetUserByUsername(username string) (*models.User, error) {
	return s.Repo.GetUserByUsername(username)
}

func (s *UserService) GetUserByRol(rol string) ([]models.User, error) {
	return s.Repo.GetUserByRol(rol)
}
