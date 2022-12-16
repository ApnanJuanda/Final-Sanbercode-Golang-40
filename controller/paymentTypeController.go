package controller

import "github.com/gin-gonic/gin"

type PaymentTypeController interface {
	AddPaymentType(ctx *gin.Context)
	GetPaymentTypes(ctx *gin.Context)
	GetPaymentTypeById(ctx *gin.Context)
	UpdatePaymentTypeById(ctx *gin.Context)
	DeletePaymentTypeById(ctx *gin.Context)
}