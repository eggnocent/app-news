package handler

import (
	"app-news/internal/adapter/handler/response"
	"app-news/internal/core/domain/entity"
	"app-news/internal/core/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

var defaultSuccessResponse response.DefaultSuccessResponse

type CategoryHandler interface {
	GetCategories(c *fiber.Ctx) error
	GetCategoryByID(c *fiber.Ctx) error
	CreateCategory(c *fiber.Ctx) error
	UpdateCategory(c *fiber.Ctx) error
	DeleteCategory(c *fiber.Ctx) error
}

type categoryHandler struct {
	categoryService service.CategoryService
}

func NewCategoryHandler(caegoryService service.CategoryService) CategoryHandler {
	return &categoryHandler{categoryService: caegoryService}
}

// GetCategories implements CategoryHandler.
func (ch *categoryHandler) GetCategories(c *fiber.Ctx) error {
	claims := c.Locals("use").(entity.JwtData)
	userID := claims.UserId
	if userID == 0 {
		code = "[HANDLER] GetCategories - 1"
		log.Errorw(code, err)
		errorResp.Meta.Message = "Unauthorized acceess"
		return c.Status(fiber.StatusUnauthorized).JSON(errorResp)
	}
	results, err := ch.categoryService.GetCategories(c.Context())
	if err != nil {
		code = "[HANDLER] GetCategories - 2"
		log.Errorw(code, err)
		errorResp.Meta.Message = "Error getting categories"
		return c.Status(fiber.StatusInternalServerError).JSON(errorResp)
	}
	categoryResponses := []response.SuccessCategoryResponse{}
	for _, result := range results {
		categoryRepsonse := response.SuccessCategoryResponse{
			ID:            result.ID,
			Title:         result.Title,
			Slug:          result.Slug,
			CreatedByName: result.User.Name,
		}
		categoryResponses = append(categoryResponses, categoryRepsonse)
	}
	defaultSuccessResponse.Meta.Status = true
	defaultSuccessResponse.Meta.Message = "Categories fetched successfully"
	defaultSuccessResponse.Data = categoryResponses

	return c.JSON(defaultSuccessResponse)
}

// GetCategoryByID implements CategoryHandler.
func (ch *categoryHandler) GetCategoryByID(c *fiber.Ctx) error {
	panic("unimplemented")
}

// CreateCategory implements CategoryHandler.
func (ch *categoryHandler) CreateCategory(c *fiber.Ctx) error {
	panic("unimplemented")
}

// UpdateCategory implements CategoryHandler.
func (ch *categoryHandler) UpdateCategory(c *fiber.Ctx) error {
	panic("unimplemented")
}

// DeleteCategory implements CategoryHandler.
func (ch *categoryHandler) DeleteCategory(c *fiber.Ctx) error {
	panic("unimplemented")
}
