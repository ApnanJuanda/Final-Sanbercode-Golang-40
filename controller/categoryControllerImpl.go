package controller

import (
	"fmt"
	"net/http"
	"project/model"
	"project/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CategoryControllerImpl struct {
	categoryService service.CategoryService
}

func NewCategoryController(categoryService service.CategoryService) CategoryController {
	return &CategoryControllerImpl{
		categoryService: categoryService,
	}
}

func (cc *CategoryControllerImpl) AddCategory(ctx *gin.Context) {
	var addCategory model.AddCategory
	if err := ctx.ShouldBindJSON(&addCategory); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	categoryExist, err := cc.categoryService.GetCategoryByName(addCategory.Name)
	if err == nil {
		ctx.JSON(http.StatusOK, categoryExist)
		return
	}

	category, err := cc.categoryService.AddCategory(&addCategory)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, category)
}

func (cc *CategoryControllerImpl) GetCategories(ctx *gin.Context) {
	categories, err := cc.categoryService.GetCategories()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}
	ctx.JSON(http.StatusOK, categories)
}

func (cc *CategoryControllerImpl) GetCategoryById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request",
		})
		return
	}
	category, err := cc.categoryService.GetCategoryById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": err,
		})
		return
	}
	ctx.JSON(http.StatusOK, category)
}

func (cc *CategoryControllerImpl) UpdateCategoryById(ctx *gin.Context) {
	var addCategory model.AddCategory
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request",
		})
		return
	}
	if err := ctx.ShouldBindJSON(&addCategory); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	count, err := cc.categoryService.UpdateCategoryById(id, &addCategory)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	message := fmt.Sprintf("Updated data amount %d", count)
	ctx.JSON(http.StatusOK, gin.H{
		"message": message,
	})
}

func (cc *CategoryControllerImpl) DeleteCategoryById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request",
		})
		return
	}
	count, err := cc.categoryService.DeleteCategoryById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	message := fmt.Sprintf("Deleted data amount %d", count)
	ctx.JSON(http.StatusOK, gin.H{
		"message": message,
	})
}