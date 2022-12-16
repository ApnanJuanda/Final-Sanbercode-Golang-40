package controller

import "github.com/gin-gonic/gin"

type TransactionController interface {
	AddTransaction(ctx *gin.Context)
	GetTransaction(ctx *gin.Context)
}