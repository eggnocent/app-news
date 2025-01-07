package service

import (
	"app-news/internal/adapter/repository"
	"app-news/internal/core/domain/entity"
	"app-news/lib/conv"
	"context"

	"github.com/gofiber/fiber/v2/log"
)

type UserService interface {
	UpdatePassword(ctx context.Context, newPass string, id int64) error
	GetUserByID(ctx context.Context, id int64) (*entity.UserEntity, error)
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

// GetUserByID implements UserService.
func (u *userService) GetUserByID(ctx context.Context, id int64) (*entity.UserEntity, error) {
	result, err := u.userRepo.GetUserByID(ctx, id)
	if err != nil {
		code = "[SERVICE] GetUserByID - 1"
		log.Errorw(code, err)
		return nil, err
	}
	return result, nil
}

// UpdatePassword implements UserService.
func (u *userService) UpdatePassword(ctx context.Context, newPass string, id int64) error {
	password, err := conv.HashPassword(newPass)
	if err != nil {
		code = "[SERVICE] UpdatePassword - 2"
		log.Errorw(code, err)
		return err
	}

	err = u.userRepo.UpdatePassword(ctx, password, id)
	if err != nil {
		code = "[SERVICE] UpdatePassword - 3"
		log.Errorw(code, err)
		return err
	}
	return nil
}
