package controller

import (
	"net/http"
	"project/model"
	"project/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TransactionControllerImpl struct {
	transactionService service.TransactionService
}

func NewTransactionController(transactionService service.TransactionService) TransactionController {
	return &TransactionControllerImpl{
		transactionService: transactionService,
	}
}

func (tc *TransactionControllerImpl) AddTransaction(ctx *gin.Context) {
	var addTransaction model.AddTransaction
	customerId := ctx.Request.Context().Value("customer_id")
	
	if err := ctx.ShouldBindJSON(&addTransaction); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	responseTransaction, err := tc.transactionService.AddTransaction(&addTransaction, customerId.(int))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, responseTransaction)
}

func (tc *TransactionControllerImpl) GetTransaction(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request",
		})
		return
	}
	transaction, err := tc.transactionService.GetTransaction(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": err,
		})
		return
	}
	ctx.JSON(http.StatusOK, transaction)
}

