package handler

import (
	"app-news/internal/adapter/handler/request"
	"app-news/internal/adapter/handler/response"
	"app-news/internal/core/domain/entity"
	"app-news/internal/core/service"
	validatorLib "app-news/lib/validator"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

type UserHander interface {
	UpdatePassword(c *fiber.Ctx) error
	GetUserByID(c *fiber.Ctx) error
}

type userHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) UserHander {
	return &userHandler{userService: userService}
}

// GetUserByID implements UserHander.
func (u *userHandler) GetUserByID(c *fiber.Ctx) error {
	claims := c.Locals("user").(*entity.JwtData)
	if claims.UserId == 0 {
		code := "[HANDLER] GetUserByID - 1"
		log.Errorw(code, err)
		errorResp.Meta.Status = false
		errorResp.Meta.Message = "Unauthorized access"

		return c.Status(fiber.StatusUnauthorized).JSON(errorResp)
	}

	user, err := u.userService.GetUserByID(c.Context(), int64(claims.UserId))
	if err != nil {
		code := "[HANDLER] GetUserByID - 2"
		log.Errorw(code, err)
		errorResp.Meta.Status = false
		errorResp.Meta.Message = err.Error()

		return c.Status(fiber.StatusInternalServerError).JSON(errorResp)
	}

	defaultSuccessResponse.Meta.Status = true
	defaultSuccessResponse.Meta.Message = "Success"
	resp := response.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
	defaultSuccessResponse.Data = resp

	return c.JSON(defaultSuccessResponse)
}

// UpdatePassword implements UserHander.
func (u *userHandler) UpdatePassword(c *fiber.Ctx) error {
	claims := c.Locals("user").(*entity.JwtData)
	if claims.UserId == 0 {
		code := "[HANDLER] UpdatePassword - 1"
		log.Errorw(code, err)
		errorResp.Meta.Status = false
		errorResp.Meta.Message = "Unauthorized access"

		return c.Status(fiber.StatusUnauthorized).JSON(errorResp)
	}

	var req request.UpdatePasswordRequest
	if err := c.BodyParser(&req); err != nil {
		code := "[HANDLER] UpdatePassword - 2"
		log.Errorw(code, err)
		errorResp.Meta.Status = false
		errorResp.Meta.Message = "Invalid request body"

		return c.Status(fiber.StatusBadRequest).JSON(errorResp)
	}

	if err = validatorLib.ValidateStruct(&req); err != nil {
		code := "[HANDLER] UpdatePassword - 3"
		log.Errorw(code, err)
		errorResp.Meta.Status = false
		errorResp.Meta.Message = err.Error()

		return c.Status(fiber.StatusBadRequest).JSON(errorResp)
	}

	err = u.userService.UpdatePassword(c.Context(), req.NewPassword, int64(claims.UserId))
	if err != nil {
		code := "[HANDLER] UpdatePassword - 4"
		log.Errorw(code, err)
		errorResp.Meta.Status = false
		errorResp.Meta.Message = err.Error()

		return c.Status(fiber.StatusInternalServerError).JSON(errorResp)
	}

	defaultSuccessResponse.Meta.Status = true
	defaultSuccessResponse.Meta.Message = "Password updated successfully"
	defaultSuccessResponse.Data = nil
	return c.JSON(defaultSuccessResponse)
}
