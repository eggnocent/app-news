package service

import (
	"app-news/internal/adapter/repository"
	"app-news/internal/core/domain/entity"
	"app-news/lib/conv"
	"context"

	"github.com/gofiber/fiber/v2/log"
)

type CategoryService interface {
	GetCategories(ctx context.Context) ([]entity.CategoryEntity, error)
	GetCategoryByID(ctx context.Context, id int64) ([]entity.CategoryEntity, error)
	CreateCategory(ctx context.Context, req entity.CategoryEntity) error
	UpdateCategory(ctx context.Context, req entity.CategoryEntity) error
	DeleteCategory(ctx context.Context, id int64) error
}

type categoryService struct {
	categoryRepository repository.CategoryRespository
}

// GetCategories implements CategoryService.
func (c *categoryService) GetCategories(ctx context.Context) ([]entity.CategoryEntity, error) {
	results, err := c.categoryRepository.GetCategories(ctx)
	if err != nil {
		code = "[SERVICE] GetCategories - 1"
		log.Errorw(code, err)
		return nil, err
	}
	return results, nil
}

// GetCategoryByID implements CategoryService.
func (c *categoryService) GetCategoryByID(ctx context.Context, id int64) ([]entity.CategoryEntity, error) {
	result, err := c.categoryRepository.GetCategoryByID(ctx, id)
	if err != nil {
		code = "[SERVICE] GetCategoryByID - 1"
		log.Errorw(code, err)
		return nil, err
	}
	return result, nil
}

// CreateCategory implements CategoryService.
func (c *categoryService) CreateCategory(ctx context.Context, req entity.CategoryEntity) error {
	slug, err := conv.GenerateSlug(req.Title)
	if err != nil {
		code = "[SERVICE] CreateCategory - 1"
		log.Errorw(code, err)
		return err
	}
	req.Slug = slug

	err = c.categoryRepository.CreateCategory(ctx, req)
	if err != nil {
		code = "[SERVICE] CreateCategory - 2"
		log.Errorw(code, err)
		return err
	}
	return nil
}

// UpdateCategory implements CategoryService.
func (c *categoryService) UpdateCategory(ctx context.Context, req entity.CategoryEntity) error {
	panic("unimplemented")
}

// DeleteCategory implements CategoryService.
func (c *categoryService) DeleteCategory(ctx context.Context, id int64) error {
	panic("unimplemented")
}

func NewCategoryService(categoryRepo repository.CategoryRespository) CategoryService {
	return &categoryService{categoryRepository: categoryRepo}
}
