package port

import (
	"github.com/WatkornFeng/go-hexa/core/domain"
)

type UserRepository interface {
	GetAll() ([]domain.User, error)
}

type UserService interface {
	GetUsers() ([]domain.User, error)
}
