package controller

import (
	"fmt"
	"net/http"
	"project/model"
	"project/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PaymentTypeControllerImpl struct {
	paymentTypeService service.PaymenTypeService
}

func NewPaymentTypeController(paymentTypeService service.PaymenTypeService) PaymentTypeController {
	return &PaymentTypeControllerImpl{
		paymentTypeService: paymentTypeService,
	}
}

func (pc *PaymentTypeControllerImpl) AddPaymentType(ctx *gin.Context) {
	var addPaymenType model.AddPaymentType
	if err := ctx.ShouldBindJSON(&addPaymenType); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// categoryExist, err := cc.categoryService.GetCategoryByName(addCategory.Name)
	// if err == nil {
	// 	ctx.JSON(http.StatusOK, categoryExist)
	// 	return
	// }

	paymenType, err := pc.paymentTypeService.AddPaymenType(&addPaymenType)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, paymenType)
}

func (pc *PaymentTypeControllerImpl) GetPaymentTypes(ctx *gin.Context) {
	paymentTypes, err := pc.paymentTypeService.GetPaymentTypes()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}
	ctx.JSON(http.StatusOK, paymentTypes)
}

func (pc *PaymentTypeControllerImpl) GetPaymentTypeById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request",
		})
		return
	}
	paymentType, err := pc.paymentTypeService.GetPaymentTypeById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": err,
		})
		return
	}
	ctx.JSON(http.StatusOK, paymentType)
}

func (pc *PaymentTypeControllerImpl) UpdatePaymentTypeById(ctx *gin.Context) {
	var addPaymenType model.AddPaymentType
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request",
		})
		return
	}
	if err := ctx.ShouldBindJSON(&addPaymenType); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	count, err := pc.paymentTypeService.UpdatePaymentTypeById(id, &addPaymenType)
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

func (pc *PaymentTypeControllerImpl) DeletePaymentTypeById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request",
		})
		return
	}
	count, err := pc.paymentTypeService.DeletePaymentTypeById(id)
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




