package handler

import (
	"strconv"

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

	return handleSuccess(c, "Get list of users success", users)
}

func (h *userHandlder) GetUser(c *fiber.Ctx) error {

	idParam := c.Params("userId")
	idUint64, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		return parameterError(c)
	}
	ctx := c.UserContext()
	user, err := h.userSrv.GetUser(ctx, idUint64)
	if err != nil {
		return handleError(c, err)
	}

	return handleSuccess(c, "Get user success", user)
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

	return handleSuccess(c, "Create user success", user)

}

type updateRequest struct {
	Name string `json:"name" validate:"required,max=15,alpha"`
}

func (h *userHandlder) UpdateUser(c *fiber.Ctx) error {
	var req updateRequest
	idParam := c.Params("userId")
	idUint64, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		return parameterError(c)
	}
	if err := c.BodyParser(&req); err != nil {
		return bodyParseError(c)
	}
	if err := validate.Struct(req); err != nil {
		return validationError(c, err)
	}
	ctx := c.UserContext()

	userReq := domain.User{
		Name: req.Name,
	}
	user, err := h.userSrv.UpdateUser(ctx, idUint64, &userReq)
	if err != nil {
		return handleError(c, err)
	}

	return handleSuccess(c, "Update user success", user)
}
