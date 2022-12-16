package controller

import (
	"fmt"
	"net/http"
	"project/model"
	"project/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CartControllerImpl struct {
	cartService service.CartService
}

func NewCartController(cartService service.CartService) CartController {
	return &CartControllerImpl{
		cartService: cartService,
	}
}

func (cc *CartControllerImpl) AddCart(ctx *gin.Context) {
	var addCart model.AddCart
	if err := ctx.ShouldBindJSON(&addCart); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	customerId := ctx.Request.Context().Value("customer_id")
	responseCart, err := cc.cartService.AddCart(&addCart, customerId.(int))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, responseCart)
}

func (cc *CartControllerImpl) GetCartByCustomerId(ctx *gin.Context) {
	customerId := ctx.Request.Context().Value("customer_id")
	responseCart, err := cc.cartService.GetCartByCustomerId(customerId.(int))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}
	ctx.JSON(http.StatusOK, responseCart)
}

func (cc *CartControllerImpl) UpdateCartProductById(ctx *gin.Context) {
	var addCartProduct model.AddCartProduct
	cartId, err := strconv.Atoi(ctx.Query("cartId"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request",
		})
		return
	}
	cartProductId, err := strconv.Atoi(ctx.Query("cartProductId"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request",
		})
		return
	}

	if err := ctx.ShouldBindJSON(&addCartProduct); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	count, err := cc.cartService.UpdateCartProductById(cartId, cartProductId, &addCartProduct)
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

func (cc *CartControllerImpl) DeleteCartProductById(ctx *gin.Context) {
	cartId, err := strconv.Atoi(ctx.Query("cartId"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request",
		})
		return
	}
	cartProductId, err := strconv.Atoi(ctx.Query("cartProductId"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request",
		})
		return
	}
	count, err := cc.cartService.DeleteCartProductById(cartId, cartProductId)
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
