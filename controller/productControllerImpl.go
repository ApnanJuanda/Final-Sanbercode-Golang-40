package controller

import (
	"fmt"
	"net/http"
	"project/model"
	"project/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductControllerImpl struct {
	productService service.ProductService
}

func NewProductController(productService service.ProductService) ProductController {
	return &ProductControllerImpl{
		productService: productService,
	}
}

func (pc *ProductControllerImpl) AddProduct(ctx *gin.Context) {
	var addProduct model.AddProduct
	if err := ctx.ShouldBindJSON(&addProduct); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	//validasi product has been exists
	// supplierExist, err := sc.supplierService.GetSupplierByName(addSupplier.Name)
	// if err == nil {
	// 	ctx.JSON(http.StatusOK, supplierExist)
	// 	return
	// }
	
	product, err := pc.productService.AddProduct(&addProduct)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, product)
}

func (pc *ProductControllerImpl) GetProducts(ctx *gin.Context) {
	products, err := pc.productService.GetProducts()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}
	ctx.JSON(http.StatusOK, products)	
}

func (pc *ProductControllerImpl) GetProductById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request",
		})
		return
	}
	product, err := pc.productService.GetProductById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": err,
		})
		return
	}
	ctx.JSON(http.StatusOK, product)
}

func (pc *ProductControllerImpl) UpdateProductById(ctx *gin.Context) {
	var addProduct model.AddProduct
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request",
		})
		return
	}
	if err := ctx.ShouldBindJSON(&addProduct); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	count, err := pc.productService.UpdateProductById(id, &addProduct)
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

func (pc *ProductControllerImpl) DeleteProductById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request",
		})
		return
	}
	count, err := pc.productService.DeleteProductById(id)
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