package handler

import (
	"app-news/internal/adapter/handler/request"
	"app-news/internal/adapter/handler/response"
	"app-news/internal/core/domain/entity"
	"app-news/internal/core/service"
	validatorLib "app-news/lib/validator"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

var err error
var code string
var errorResp response.ErrorResponseDefault
var validate = validator.New()

type AuthHandler interface {
	Login(c *fiber.Ctx) error
}

type authHandler struct {
	authService service.AuthService
}

func NewAuthHandler(authService service.AuthService) AuthHandler {
	return &authHandler{authService: authService}
}

// Login implements AuthHandler.
func (a *authHandler) Login(c *fiber.Ctx) error {
	req := request.LoginRequest{}
	resp := response.SuccessAuthResponse{}

	if err = c.BodyParser(&req); err != nil {
		code = "[HANDLER] Login -1"
		log.Errorw(code, err)
		errorResp.Meta.Status = false
		errorResp.Meta.Message = err.Error()
	}

	if err = validatorLib.ValidateStruct(req); err != nil {
		code = "[HANDLER] Login -2"
		log.Errorw(code, err)
		errorResp.Meta.Status = false
		errorResp.Meta.Message = err.Error()

		return c.Status(fiber.StatusBadRequest).JSON(errorResp)
	}

	reqLogin := entity.LoginRequest{
		Email:    req.Email,
		Password: req.Password,
	}

	result, err := a.authService.GetUserByEmail(c.Context(), reqLogin)
	if err != nil {
		code = "[HANDLER] Login -3"
		log.Errorw(code, err)
		errorResp.Meta.Status = false
		errorResp.Meta.Message = "Invalid email or password"
		if err.Error() == "invalid password" {
			return c.Status(fiber.StatusUnauthorized).JSON(errorResp)
		}
		return c.Status(fiber.StatusInternalServerError).JSON(errorResp)
	}

	resp.Meta.Status = true
	resp.Meta.Message = "Login successfull"
	resp.AccessToken = result.AccessToken
	resp.ExpiresAt = result.ExpiresAt

	return c.JSON(resp)
}
