package service

import "project/model"

type CartService interface {
	AddCart(addCart *model.AddCart, customerId int) (*model.ResponseCart, error)
	GetCartById(cartId int) (*[]model.CartProduct, error)
	GetCartByCustomerId(customerId int) (*model.ResponseCart, error)
	UpdateCartProductById(cartId, cartProductId int, addCartProduct *model.AddCartProduct) (int, error)
	DeleteCartProductById(cartId, cartProductId int) (int, error)
}