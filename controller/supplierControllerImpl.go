package controller

import (
	"fmt"
	"net/http"
	"project/model"
	"project/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SupplierControllerImpl struct {
	supplierService service.SupplierService
}

func NewSupplierController(supplierService service.SupplierService) SupplierController {
	return &SupplierControllerImpl{
		supplierService: supplierService,
	}
}

func (sc *SupplierControllerImpl) AddSupplier(ctx *gin.Context) {
	var addSupplier model.AddSupplier
	if err := ctx.ShouldBindJSON(&addSupplier); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	supplierExist, err := sc.supplierService.GetSupplierByName(addSupplier.Name)
	if err == nil {
		ctx.JSON(http.StatusOK, supplierExist)
		return
	}
	
	supplier, err := sc.supplierService.AddSupplier(&addSupplier)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, supplier)
}

func (sc *SupplierControllerImpl) GetSuppliers(ctx *gin.Context) {
	suppliers, err := sc.supplierService.GetSuppliers()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}
	ctx.JSON(http.StatusOK, suppliers)
}

func (sc *SupplierControllerImpl) GetSupplierById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request",
		})
		return
	}
	supplier, err := sc.supplierService.GetSupplierById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": err,
		})
		return
	}
	ctx.JSON(http.StatusOK, supplier)
}

func (sc *SupplierControllerImpl) UpdateSupplierById(ctx *gin.Context) {
	var addSupplier model.AddSupplier
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request",
		})
		return
	}
	if err := ctx.ShouldBindJSON(&addSupplier); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	count, err := sc.supplierService.UpdateSupplierById(id, &addSupplier)
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

func (sc *SupplierControllerImpl) DeleteSupplierById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request",
		})
		return
	}
	count, err := sc.supplierService.DeleteSupplierById(id)
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