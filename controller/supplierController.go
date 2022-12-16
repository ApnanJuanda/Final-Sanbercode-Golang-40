package controller

import "github.com/gin-gonic/gin"

type SupplierController interface {
	AddSupplier(ctx *gin.Context)
	GetSuppliers(ctx *gin.Context)
	GetSupplierById(ctx *gin.Context)
	UpdateSupplierById(ctx *gin.Context)
	DeleteSupplierById(ctx *gin.Context)
}