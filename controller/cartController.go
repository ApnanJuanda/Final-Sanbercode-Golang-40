package controller

import "github.com/gin-gonic/gin"

type CartController interface {
	AddCart(ctx *gin.Context)
	GetCartByCustomerId(ctx *gin.Context)
	UpdateCartProductById(ctx *gin.Context)
	DeleteCartProductById(ctx *gin.Context)
}
