package controller

import "github.com/gin-gonic/gin"

type ProductController interface {
	AddProduct(ctx *gin.Context)
	GetProducts(ctx *gin.Context)
	GetProductById(ctx *gin.Context)
	UpdateProductById(ctx *gin.Context)
	DeleteProductById(ctx *gin.Context)
}