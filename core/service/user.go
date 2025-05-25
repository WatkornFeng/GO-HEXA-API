package service

import (
	"log"

	"github.com/WatkornFeng/go-hexa/core/domain"
	"github.com/WatkornFeng/go-hexa/core/port"
)

type userService struct {
	repo port.UserRepository
}

func NewUserService(repo port.UserRepository) port.UserService {
	return &userService{repo: repo}
}

func (us *userService) GetUsers() ([]domain.User, error) {
	users, err := us.repo.GetAll()
	if err != nil {
		log.Println("Error retrieving users:", err)
		return nil, err
	}

	return users, nil
}
