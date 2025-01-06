package middleware

import (
	"app-news/config"
	"app-news/internal/adapter/handler/response"
	"app-news/lib/auth"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

type Middleware interface {
	CheckToken() fiber.Handler
}

type Options struct {
	authJwt auth.Jwt
}

// CheckToken implements Middleware.
func (o *Options) CheckToken() func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		var errorResponse response.ErrorResponseDefault
		authHandler := c.Get("Authorization")
		if authHandler == "" {
			log.Error("[MIDDLEWARE] Missing authorization header")
			errorResponse.Meta.Status = false
			errorResponse.Meta.Message = "Missing authorization Header"
			return c.Status(fiber.StatusUnauthorized).JSON(errorResponse)
		}

		tokenString := strings.Split(authHandler, "Bearer ")[1]
		claims, err := o.authJwt.VerifyAccessToken(tokenString)
		if err != nil {
			log.Errorf("[MIDDLEWARE] Invalid or expired token: %v", err)
			errorResponse.Meta.Status = false
			errorResponse.Meta.Message = "Invalid or expired token"
			return c.Status(fiber.StatusUnauthorized).JSON(errorResponse)
		}

		log.Infof("[MIDDLEWARE] Token validated. Claims: %+v", claims)
		c.Locals("user", claims) // Set klaim ke context

		return c.Next()
	}
}

func NewMiddleware(cfg *config.Config) Middleware {
	opt := new(Options)
	opt.authJwt = auth.NewJwt(cfg)

	return opt
}
