package handler

import (
	"app-news/internal/adapter/handler/request"
	"app-news/internal/adapter/handler/response"
	"app-news/internal/core/domain/entity"
	"app-news/internal/core/service"
	"app-news/lib/conv"
	validatorLib "app-news/lib/validator"

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
	claims := c.Locals("user").(*entity.JwtData)
	userID := claims.UserId
	if userID == 0 {
		code = "[HANDLER] GetCategories - 1"
		log.Errorw(code, err)
		errorResp.Meta.Status = false
		errorResp.Meta.Message = "Unauthorized access"

		return c.Status(fiber.StatusUnauthorized).JSON(errorResp)
	}

	results, err := ch.categoryService.GetCategories(c.Context())
	if err != nil {
		code = "[HANDLER] GetCategories - 2"
		log.Errorw(code, err)
		errorResp.Meta.Status = false
		errorResp.Meta.Message = err.Error()

		return c.Status(fiber.StatusInternalServerError).JSON(errorResp)
	}

	categoryResponses := []response.SuccessCategoryResponse{}
	for _, result := range results {
		categoryResponse := response.SuccessCategoryResponse{
			ID:            result.ID,
			Title:         result.Title,
			Slug:          result.Slug,
			CreatedByName: result.User.Name,
		}
		categoryResponses = append(categoryResponses, categoryResponse)
	}
	defaultSuccessResponse.Meta.Status = true
	defaultSuccessResponse.Pagination = nil
	defaultSuccessResponse.Meta.Message = "Categories fetched successfully"
	defaultSuccessResponse.Data = categoryResponses

	return c.JSON(defaultSuccessResponse)
}

// GetCategoryByID implements CategoryHandler.
func (ch *categoryHandler) GetCategoryByID(c *fiber.Ctx) error {
	claims := c.Locals("user").(*entity.JwtData)
	userID := claims.UserId
	if userID == 0 {
		code = "[HANDLER] GetCategoriesbyID - 1"
		log.Errorw(code, err)
		errorResp.Meta.Status = false
		errorResp.Meta.Message = "Unauthorized access"

		return c.Status(fiber.StatusUnauthorized).JSON(errorResp)
	}

	idParam := c.Params("categoryID")
	id, err := conv.StringtoInt(idParam)
	if err != nil {
		code = "[HANDLER] GetCategoriesbyID - 2"
		log.Errorw(code, err)
		errorResp.Meta.Status = false
		errorResp.Meta.Message = err.Error()

		return c.Status(fiber.StatusBadRequest).JSON(errorResp)
	}

	result, err := ch.categoryService.GetCategoryByID(c.Context(), id)
	if err != nil {
		code = "[HANDLER] GetCategoriesbyID - 3"
		log.Errorw(code, err)
		errorResp.Meta.Status = false
		errorResp.Meta.Message = err.Error()

		return c.Status(fiber.StatusInternalServerError).JSON(errorResp)
	}

	categoryResponse := response.SuccessCategoryResponse{
		ID:            result.ID,
		Title:         result.Title,
		Slug:          result.Slug,
		CreatedByName: result.User.Name,
	}

	defaultSuccessResponse.Meta.Status = true
	defaultSuccessResponse.Pagination = nil
	defaultSuccessResponse.Meta.Message = "Category fetched successfully"
	defaultSuccessResponse.Data = categoryResponse

	return c.JSON(defaultSuccessResponse)
}

// CreateCategory implements CategoryHandler.
func (ch *categoryHandler) CreateCategory(c *fiber.Ctx) error {
	var req request.CategoryReqest
	claims := c.Locals("user").(*entity.JwtData)
	userID := claims.UserId

	if userID == 0 {
		code = "[HANDLER] CreateCategory - 1"
		log.Errorw(code, err)
		errorResp.Meta.Status = false
		errorResp.Meta.Message = "Unauthorized access"

		return c.Status(fiber.StatusUnauthorized).JSON(errorResp)
	}

	if err = c.BodyParser(&req); err != nil {
		code = "[HANDLER] CreateCategory - 2"
		log.Errorw(code, err)
		errorResp.Meta.Status = false
		errorResp.Meta.Message = "Invalid request body"

		return c.Status(fiber.StatusUnauthorized).JSON(errorResp)
	}

	if err = validatorLib.ValidateStruct(req); err != nil {
		code = "[HANDLER] CreateCategory - 3"
		log.Errorw(code, err)
		errorResp.Meta.Status = false
		errorResp.Meta.Message = err.Error()

		return c.Status(fiber.StatusBadRequest).JSON(errorResp)
	}

	reqEntity := entity.CategoryEntity{
		Title: req.Title,
		User: entity.UserEntity{
			ID: int64(userID),
		},
	}

	err = ch.categoryService.CreateCategory(c.Context(), reqEntity)
	if err != nil {
		code = "[HANDLER] CreateCategory - 4"
		log.Errorw(code, err)
		errorResp.Meta.Status = false
		errorResp.Meta.Message = err.Error()

		return c.Status(fiber.StatusInternalServerError).JSON(errorResp)
	}

	defaultSuccessResponse.Data = nil
	defaultSuccessResponse.Pagination = nil
	defaultSuccessResponse.Meta.Status = true
	defaultSuccessResponse.Meta.Message = "Category created successfully"
	return c.Status(fiber.StatusCreated).JSON(defaultSuccessResponse)
}

// UpdateCategory implements CategoryHandler.
func (ch *categoryHandler) UpdateCategory(c *fiber.Ctx) error {
	var req request.CategoryReqest
	claims := c.Locals("user").(*entity.JwtData)
	userID := claims.UserId
	if userID == 0 {
		code = "[HANDLER] UpdateCategory - 1"
		log.Errorw(code, err)
		errorResp.Meta.Status = false
		errorResp.Meta.Message = "Unauthorized access"

		return c.Status(fiber.StatusUnauthorized).JSON(errorResp)
	}

	if err = c.BodyParser(&req); err != nil {
		code = "[HANDLER] UpdateCategory - 2"
		log.Errorw(code, err)
		errorResp.Meta.Status = false
		errorResp.Meta.Message = "Invalid request body"

		return c.Status(fiber.StatusUnauthorized).JSON(errorResp)
	}

	if err = validatorLib.ValidateStruct(req); err != nil {
		code = "[HANDLER] UpdateCategory - 3"
		log.Errorw(code, err)
		errorResp.Meta.Status = false
		errorResp.Meta.Message = err.Error()

		return c.Status(fiber.StatusBadRequest).JSON(errorResp)
	}

	idParam := c.Params("categoryID")
	id, err := conv.StringtoInt(idParam)
	if err != nil {
		code = "[HANDLER] UpdateCategory - 4"
		log.Errorw(code, err)
		errorResp.Meta.Status = false
		errorResp.Meta.Message = err.Error()

		return c.Status(fiber.StatusBadRequest).JSON(errorResp)
	}

	reqEntity := entity.CategoryEntity{
		ID:    id,
		Title: req.Title,
		User: entity.UserEntity{
			ID: int64(userID),
		},
	}

	err = ch.categoryService.UpdateCategory(c.Context(), reqEntity)
	if err != nil {
		code = "[HANDLER] UpdateCategory - 5"
		log.Errorw(code, err)
		errorResp.Meta.Status = false
		errorResp.Meta.Message = err.Error()

		return c.Status(fiber.StatusInternalServerError).JSON(errorResp)
	}

	defaultSuccessResponse.Data = nil
	defaultSuccessResponse.Pagination = nil
	defaultSuccessResponse.Meta.Status = true
	defaultSuccessResponse.Meta.Message = "Category updated successfully"

	return c.JSON(defaultSuccessResponse)
}

// DeleteCategory implements CategoryHandler.
func (ch *categoryHandler) DeleteCategory(c *fiber.Ctx) error {
	claims := c.Locals("user").(*entity.JwtData)
	userID := claims.UserId

	if userID == 0 {
		code = "[HANDLER] DeleteCategory - 1"
		log.Errorw(code, err)
		errorResp.Meta.Status = false
		errorResp.Meta.Message = "Unauthorized access"

		return c.Status(fiber.StatusUnauthorized).JSON(errorResp)
	}

	idParam := c.Params("categoryID")
	id, err := conv.StringtoInt(idParam)
	if err != nil {
		code = "[HANDLER] DeleteCategory - 4"
		log.Errorw(code, err)
		errorResp.Meta.Status = false
		errorResp.Meta.Message = err.Error()

		return c.Status(fiber.StatusBadRequest).JSON(errorResp)
	}

	err = ch.categoryService.DeleteCategory(c.Context(), id)
	if err != nil {
		code = "[HANDLER] DeleteCategory - 5"
		log.Errorw(code, err)
		errorResp.Meta.Status = false
		errorResp.Meta.Message = err.Error()

		return c.Status(fiber.StatusInternalServerError).JSON(errorResp)
	}

	defaultSuccessResponse.Data = nil
	defaultSuccessResponse.Pagination = nil
	defaultSuccessResponse.Meta.Status = true
	defaultSuccessResponse.Meta.Message = "Category deleted successfully"
	return c.JSON(defaultSuccessResponse)
}
