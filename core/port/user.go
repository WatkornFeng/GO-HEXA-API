package port

import (
	"context"

	"github.com/WatkornFeng/go-hexa/core/domain"
	"github.com/WatkornFeng/go-hexa/core/domain/dto"
)

type UserRepository interface {
	GetAll(ctx context.Context) ([]domain.User, error)
	CreateUser(ctx context.Context, user *domain.User) (*domain.User, error)
	FindUserByEmail(ctx context.Context, email string) (*domain.User, error)
	GetUserByID(ctx context.Context, id uint64) (*domain.User, error)
	UpdateUserByID(ctx context.Context, id uint64, updateData *domain.User) (*domain.User, error)

	DeleteUserByID(ctx context.Context, id uint64) (bool, error)
}

type UserService interface {
	GetUsers(ctx context.Context) ([]dto.UserResponse, error)
	GetUser(ctx context.Context, id uint64) (*dto.UserResponse, error)
	Register(ctx context.Context, user *domain.User) (*dto.UserResponse, error)
	UpdateUser(ctx context.Context, id uint64, updateData *domain.User) (*dto.UpdateUserResponse, error)

	DeleteUser(ctx context.Context, id uint64) error
}
