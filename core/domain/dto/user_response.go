package dto

import (
	"time"

	"github.com/WatkornFeng/go-hexa/core/domain"
)

// dto => Data Transfer Object
type UserResponse struct {
	UserID uint   `json:"user_id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
}

func NewUserResponse(user *domain.User) *UserResponse {
	return &UserResponse{
		UserID: user.ID,
		Name:   user.Name,
		Email:  user.Email,
	}
}
func NewListUsersResponse(users []domain.User) []UserResponse {
	listUsersResponse := make([]UserResponse, len(users))
	for i, user := range users {
		listUsersResponse[i] = UserResponse{
			UserID: user.ID,
			Name:   user.Name,
			Email:  user.Email,
		}
	}
	return listUsersResponse
}

type UpdateUserResponse struct {
	UserID   uint      `json:"user_id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	UpdateAt time.Time `json:"update_at"`
}

func NewUpdateUserResponse(user *domain.User) *UpdateUserResponse {
	return &UpdateUserResponse{
		UserID:   user.ID,
		Name:     user.Name,
		Email:    user.Email,
		UpdateAt: user.UpdatedAt,
	}
}
