package handler

import (
	"github.com/WatkornFeng/go-hexa/adapter/handler/dto"
	"github.com/WatkornFeng/go-hexa/core/port"
	"github.com/gofiber/fiber/v2"
)

type userHandlder struct {
	userSrv port.UserService
}

func NewUserHandler(userSrv port.UserService) *userHandlder {
	return &userHandlder{userSrv: userSrv}
}

func (h userHandlder) GetUsers(c *fiber.Ctx) error {
	users, err := h.userSrv.GetUsers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch users",
		})
	}

	usersResponse := make([]dto.UserResponse, len(users))
	for i, user := range users {
		usersResponse[i] = dto.UserResponse{
			UserID: user.ID,
			Name:   user.Name,
			Email:  user.Email,
		}
	}
	return c.JSON(usersResponse)
}
