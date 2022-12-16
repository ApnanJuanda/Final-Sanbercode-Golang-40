package controller

import "github.com/gin-gonic/gin"

type CategoryController interface {
	AddCategory(ctx *gin.Context)
	GetCategories(ctx *gin.Context)
	GetCategoryById(ctx *gin.Context)
	UpdateCategoryById(ctx *gin.Context)
	DeleteCategoryById(ctx *gin.Context)
}
