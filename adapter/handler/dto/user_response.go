package dto

// dto => Data Transfer Object
type UserResponse struct {
	UserID uint   `json:"customer_id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
}
