package controller

import (
	"fmt"
	"net/http"
	"project/model"
	"project/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CustomerControllerImpl struct {
	customerService service.CustomerService
}

func NewCustomerController(customerService service.CustomerService) CustomerController {
	return &CustomerControllerImpl{
		customerService: customerService,
	}
}

func (cc *CustomerControllerImpl) RegisterCustomer(ctx *gin.Context) {
	var customerRegister model.CustomerRegister
	if err := ctx.ShouldBindJSON(&customerRegister); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	token, err := cc.customerService.RegisterCustomer(&customerRegister)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"token": token,
	})
}

func (cc *CustomerControllerImpl) LoginCustomer(ctx *gin.Context) {
	var customerLogin model.CustomerLogin
	if err := ctx.ShouldBindJSON(&customerLogin); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	token, err := cc.customerService.LoginCustomer(&customerLogin)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
	})	
	
}

func (cc *CustomerControllerImpl) UpdateCustomer(ctx *gin.Context) {
	var customerRegister model.CustomerRegister
	customerId := ctx.Request.Context().Value("customer_id")
	if customerId == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request",
		})
		return
	}
	if err := ctx.ShouldBindJSON(&customerRegister); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	
	count, err := cc.customerService.UpdateCustomer(customerId.(int), &customerRegister)
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

func (cc *CustomerControllerImpl) DeleteCustomer(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request",
		})
		return
	}
	count, err := cc.customerService.DeleteCustomer(id)
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


