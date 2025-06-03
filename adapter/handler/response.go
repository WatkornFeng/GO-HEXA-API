package handler

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/WatkornFeng/go-hexa/core/domain"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func SuccessResponse(c *fiber.Ctx, status int, message string, data any) error {
	resp := fiber.Map{
		"success": true,
		"message": message,
	}

	if data != nil {
		resp["data"] = data
	}
	return c.Status(status).JSON(resp)
}

func handleSuccess(c *fiber.Ctx, message string, data any) error {
	if data == nil && message == "" {
		return c.SendStatus(fiber.StatusNoContent)
	}

	return SuccessResponse(c, fiber.StatusOK, message, data)
}
func ErrorResponse(c *fiber.Ctx, status int, errMessage string, details any) error {
	resp := fiber.Map{
		"success": true,
		"error":   errMessage,
	}
	if details != nil {
		resp["details"] = details
	}

	return c.Status(status).JSON(resp)
}
func FormatValidationError(err error) map[string]string {
	errors := make(map[string]string)

	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, fieldErr := range validationErrors {
			field := strings.ToLower(fieldErr.Field())
			tag := fieldErr.Tag()
			param := fieldErr.Param()

			switch tag {
			case "required":
				errors[field] = "is required"
			case "email":
				errors[field] = "must be a valid email address"
			case "min":
				errors[field] = fmt.Sprintf("must be at least %s characters", param)
			case "max":
				errors[field] = fmt.Sprintf("must be at most %s characters", param)
			case "alpha":
				errors[field] = "must contain only letters"
			default:
				errors[field] = fmt.Sprintf("is not valid (%s)", tag)
			}
		}
	} else {
		// use for development
		slog.Error("Unexpected this error type in validation", "error", err)
		// errors["error"] = err.Error()
	}

	return errors
}

func validationError(c *fiber.Ctx, err error) error {
	validationDetails := FormatValidationError(err)
	return ErrorResponse(c, fiber.StatusBadRequest, "Validation failed", validationDetails)
}
func bodyParseError(c *fiber.Ctx) error {
	return ErrorResponse(c, fiber.StatusBadRequest, "Invalid request body", nil)
}

func handleError(c *fiber.Ctx, err error) error {
	if appErr, ok := err.(*domain.AppError); ok {
		return ErrorResponse(c, appErr.Code, appErr.Message, nil)
	}
	// Default fallback
	return ErrorResponse(c, fiber.StatusInternalServerError, "Internal server error", nil)
}
