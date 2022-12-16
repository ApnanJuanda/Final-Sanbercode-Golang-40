package controller

import "github.com/gin-gonic/gin"

type CustomerController interface {
	RegisterCustomer(ctx *gin.Context)
	LoginCustomer(ctx *gin.Context)
	UpdateCustomer(ctx *gin.Context)
	DeleteCustomer(ctx *gin.Context)
}