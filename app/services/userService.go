package services

import (
	"errors"
	"go-backend/app/models"

	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	Repo *repositories.userRepository
}

func (s *userService) isUserDuplicated(user string) (bool, error) {
	users, err := s.Repo.GetAllUsers()
	if err != nil {
		return false, err
	}

	for _, u := range users {
		if u.User == user {
			return true, nil
		}
	}

	return false, nil
}

func (s *userService) isNameDuplicated(name, patsurn, matsurn string) (bool, error) {
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

func (s *userService) CreateUser(user models.User) error {
	isDuplicated, err := s.isUserDuplicated(user.User)
	if err != nil {
		return err
	}

	if isDuplicated {
		return errors.New("User already exists")
	}

	isNameDuplicated, err := s.isNameDuplicated(user.Name, user.Patsurn, user.Matsurn)
	if err != nil {
		return err
	}

	if isNameDuplicated {
		return errors.New("Name already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)

	user.Image = "default.png"

	err = s.Repo.CreateUser(user)
	if err != nil {
		return err
	}
	return nil

}
