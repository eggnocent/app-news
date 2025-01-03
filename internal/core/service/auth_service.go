package service

import (
	"app-news/config"
	"app-news/internal/adapter/repository"
	"app-news/internal/core/domain/entity"
	"app-news/lib/auth"
	"app-news/lib/conv"
	"context"
	"errors"
	"time"

	"github.com/gofiber/fiber/v2/log"
	"github.com/golang-jwt/jwt/v5"
)

var err error
var code string

type AuthService interface {
	GetUserByEmail(ctx context.Context, req entity.LoginRequest) (*entity.AccessToken, error)
}

type authService struct {
	authRepository repository.AuthRepository
	cfg            *config.Config
	jwtToken       auth.Jwt
}

func NewAuthRepository(authRepository repository.AuthRepository, cfg *config.Config, jwtToken auth.Jwt) AuthService {
	return &authService{
		authRepository: authRepository,
		cfg:            cfg,
		jwtToken:       jwtToken,
	}
}

func (a *authService) GetUserByEmail(ctx context.Context, req entity.LoginRequest) (*entity.AccessToken, error) {
	result, err := a.authRepository.GetUserByEmail(ctx, req)
	if err != nil {
		code = "[SERVICE] GetUserByEmail -1"
		log.Errorw(code, err)
	}

	if checkPass := conv.CheckPasswordHash(req.Password, result.Password); !checkPass {
		code = "[SERVICE] GetUserByEmail -2"
		err = errors.New("invalid password")
		log.Errorw(code, err)
		return nil, err
	}

	jwtData := entity.JwtData{
		UserId: float64(result.ID),
		RegisteredClaims: jwt.RegisteredClaims{
			NotBefore: jwt.NewNumericDate(time.Now().Add(time.Hour * 2)),
			ID:        string(result.ID),
		},
	}

	accessToken, expiresAt, err := a.jwtToken.GenerateToken(&jwtData)
	if err != nil {
		code = "[SERVICE] GetUserByEmail -3"
		log.Errorw(code, err)
		return nil, err
	}

	resp := entity.AccessToken{
		AccessToken: accessToken,
		ExpiresAt:   expiresAt,
	}

	return &resp, nil
}
