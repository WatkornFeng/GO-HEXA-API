package handler

import (
	"github.com/WatkornFeng/go-hexa/core/domain"
	"github.com/WatkornFeng/go-hexa/core/port"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type userHandlder struct {
	userSrv port.UserService
}

func NewUserHandler(userSrv port.UserService) *userHandlder {
	return &userHandlder{userSrv: userSrv}
}

func (h *userHandlder) GetUsers(c *fiber.Ctx) error {
	ctx := c.UserContext()
	users, err := h.userSrv.GetUsers(ctx)
	if err != nil {
		return handleError(c, err)
	}

	// rsp := dto.NewListUsersResponse(users)

	return handleSuccess(c, "Get list of users success", users)
}

type registerRequest struct {
	Name  string `json:"name" validate:"required,max=15,alpha"`
	Email string `json:"email" validate:"required,email"`
	// Password string `json:"password" validate:"required,min=8"`
}

var validate = validator.New()

func (h *userHandlder) Register(c *fiber.Ctx) error {
	var req registerRequest

	if err := c.BodyParser(&req); err != nil {
		return bodyParseError(c)
	}

	if err := validate.Struct(req); err != nil {
		return validationError(c, err)
	}

	ctx := c.UserContext()
	userReq := domain.User{
		Name:  req.Name,
		Email: req.Email,
	}
	user, err := h.userSrv.Register(ctx, &userReq)
	if err != nil {
		return handleError(c, err)
	}

	// rsp := dto.NewUserResponse(user)

	return handleSuccess(c, "Create user success", user)

}
