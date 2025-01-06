package auth

import (
	"app-news/config"
	"app-news/internal/core/domain/entity"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2/log"
	"github.com/golang-jwt/jwt/v5"
)

type Jwt interface {
	GenerateToken(data *entity.JwtData) (string, int64, error)
	VerifyAccessToken(token string) (*entity.JwtData, error)
}

type Options struct {
	SigningKey string
	Issuer     string
}

// GenerateToken implements Jwt.
func (o *Options) GenerateToken(data *entity.JwtData) (string, int64, error) {
	now := time.Now().Local()
	expiresAt := now.Add(time.Hour * 24)
	data.RegisteredClaims.ExpiresAt = jwt.NewNumericDate(expiresAt)
	data.RegisteredClaims.Issuer = o.Issuer
	data.RegisteredClaims.NotBefore = jwt.NewNumericDate(now)
	acToken := jwt.NewWithClaims(jwt.SigningMethodHS256, data)
	accessToken, err := acToken.SignedString([]byte(o.SigningKey))
	if err != nil {
		return "", 0, err
	}

	return accessToken, expiresAt.Unix(), nil
}

// VerifyAccessToken implements Jwt.
func (o *Options) VerifyAccessToken(token string) (*entity.JwtData, error) {
	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			log.Error("[JWT] Invalid signing method")
			return nil, fmt.Errorf("Invalid signing method")
		}
		return []byte(o.SigningKey), nil
	})

	if err != nil {
		log.Errorf("[JWT] Error parsing token: %v", err)
		return nil, err
	}

	// Pastikan token valid
	if !parsedToken.Valid {
		log.Error("[JWT] Token is invalid or expired")
		return nil, fmt.Errorf("Token is invalid or expired")
	}

	// Ambil klaim dari token
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		log.Error("[JWT] Failed to parse claims as MapClaims")
		return nil, fmt.Errorf("Failed to parse claims")
	}

	log.Infof("[JWT] Parsed claims: %+v", claims)

	// Ambil UserId dari klaim
	userID, ok := claims["user_id"].(float64)
	if !ok {
		log.Error("[JWT] Missing or invalid user_id in claims")
		return nil, fmt.Errorf("Invalid or missing user_id in token claims")
	}

	// Return JwtData
	return &entity.JwtData{
		UserId: userID,
	}, nil
}

func NewJwt(cfg *config.Config) Jwt {
	opt := new(Options)
	opt.SigningKey = cfg.App.JwtSecretKey
	opt.Issuer = cfg.App.JwtIssuer

	return opt
}
