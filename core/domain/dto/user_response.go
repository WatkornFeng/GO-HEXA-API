package dto

import "github.com/WatkornFeng/go-hexa/core/domain"

// dto => Data Transfer Object
type UserResponse struct {
	UserID uint   `json:"customer_id"`
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
