package service

import (
	"database/sql"
	"errors"
	"fmt"
	"project/helper"
	"project/model"
)

type CartServiceImpl struct {
	DB *sql.DB
}

func NewCartService(DB *sql.DB) CartService {
	return &CartServiceImpl{
		DB: DB,
	}
}

func (cs *CartServiceImpl) AddCart(addCart *model.AddCart, customerId int) (*model.ResponseCart, error) {
	var responseCart = model.ResponseCart{}
	var newCart = model.Cart{}

	if addCart.CartProducts == nil {
		return nil, errors.New("product can not empty")
	}

	//check apakah sudah ada cart untuk customer tersebut
	cartExisting, err := cs.GetCartByCustomerId(customerId)
	if err != nil {
		//save cart (id, customerId)
		sqlStatement = `INSERT INTO cart (customer_id) VALUES ($1) Returning *`
		err = cs.DB.QueryRow(sqlStatement, customerId).Scan(&newCart.ID, &newCart.CustomerId, &newCart.CreatedAt, &newCart.UpdatedAt)
		if err != nil {
			return nil, err
		}
		cs.SaveCartProduct(addCart, newCart.ID)
	} else {
		cs.SaveCartProduct(addCart, cartExisting.ID)
	}

	// get All cartProduct by cartId
	cartProducts, err := cs.GetCartById(newCart.ID)
	var responseCartProducts = []model.ResponseCartProduct{}
	if err != nil {
		return nil, err
	}

	// populator from cartProducts to responseCartProducts
	for _, value := range *cartProducts {
		var responseCartProduct = model.ResponseCartProduct{}
		responseCartProduct.ID = value.ID
		responseCartProduct.Quantity = value.Quantity
		responseCartProduct.TotalPrice = value.TotalPrice
		responseCartProduct.CartId = value.CartId

		//get product by id
		productService := NewProductService(cs.DB)
		product, _ := productService.GetProductById(value.ProductId)
		if err != nil {
			return nil, err
		}
		responseCartProduct.Product = *product
		responseCartProducts = append(responseCartProducts, responseCartProduct)
	}

	responseCart.ID = newCart.ID
	responseCart.CartProducts = responseCartProducts
	return &responseCart, nil
}

func (cs *CartServiceImpl) GetCartById(cartId int) (*[]model.CartProduct, error) {
	var cartProducts = []model.CartProduct{}

	sqlStatement = `SELECT * FROM cart_product WHERE cart_id=($1)`

	rows, err := cs.DB.Query(sqlStatement, cartId)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var cartProduct = model.CartProduct{}

		err = rows.Scan(&cartProduct.ID, &cartProduct.ProductId, &cartProduct.Quantity, &cartProduct.TotalPrice,
			&cartProduct.CartId, &cartProduct.CreatedAt, &cartProduct.UpdatedAt)

		if err != nil {
			return nil, err
		}

		cartProducts = append(cartProducts, cartProduct)

	}
	return &cartProducts, nil
}

func (cs *CartServiceImpl) GetCartByCustomerId(customerId int) (*model.ResponseCart, error) {
	var cart = model.Cart{}

	sqlStatement := `SELECT * FROM cart WHERE customer_id=($1)`
	err = cs.DB.QueryRow(sqlStatement, customerId).Scan(&cart.ID, &cart.CustomerId, &cart.CreatedAt, &cart.UpdatedAt)
	if err != nil {
		return nil, err
	}

	// get cartProduct by cart.ID
	cartProducts, err := cs.GetCartById(cart.ID)
	if err != nil {
		return nil, err
	}

	// populator from cartProducts to responseCartProducts
	var responseCartProducts = []model.ResponseCartProduct{}
	for _, value := range *cartProducts {
		var responseCartProduct = model.ResponseCartProduct{}
		responseCartProduct.ID = value.ID
		responseCartProduct.Quantity = value.Quantity
		responseCartProduct.TotalPrice = value.TotalPrice
		responseCartProduct.CartId = value.CartId

		//get product by id
		productService := NewProductService(cs.DB)
		product, _ := productService.GetProductById(value.ProductId)
		if err != nil {
			return nil, err
		}
		responseCartProduct.Product = *product
		responseCartProducts = append(responseCartProducts, responseCartProduct)
	}

	// responseCartProduct
	var responseCart = model.ResponseCart{}
	responseCart.ID = cart.ID
	responseCart.CartProducts = responseCartProducts

	return &responseCart, nil
}

func (cs *CartServiceImpl) SaveCartProduct(addCart *model.AddCart, cartId int) error {
	//save masing-masing cart_product dengan id cart yang baru disimpan
	for _, addCartProduct := range addCart.CartProducts {

		// get productById -> dapatin harganya -> kali dengan quantitynya
		productService := NewProductService(cs.DB)
		product, _ := productService.GetProductById(addCartProduct.ProductId)
		if product == nil {
			return errors.New("product_id is not valid")
		}

		// populator cartProduct
		cartProduct, err := helper.CartProductPopulator(&addCartProduct)
		if err != nil {
			return err
		}
		if addCartProduct.Quantity > product.StockProduct {
			return errors.New("stock Product is not enough")
		}
		cartProduct.ProductId = addCartProduct.ProductId
		cartProduct.CartId = cartId
		cartProduct.TotalPrice = addCartProduct.Quantity * product.Price

		// save query
		sqlStatement = `INSERT INTO cart_product (product_id, quantity, total_price, cart_id) VALUES ($1, $2, $3, $4)`
		_, err = cs.DB.Exec(sqlStatement, cartProduct.ProductId, cartProduct.Quantity, cartProduct.TotalPrice, cartProduct.CartId)
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
	}

	return nil
}

func (cs *CartServiceImpl) UpdateCartProductById(cartId, cartProductId int, addCartProduct *model.AddCartProduct) (int, error) {
	if addCartProduct.Quantity < 0 {
		return 0, errors.New("invalid request body")
	}

	// get productById -> dapatin harganya -> kali dengan quantitynya
	productService := NewProductService(cs.DB)
	product, _ := productService.GetProductById(addCartProduct.ProductId)
	if product == nil {
		return 0, errors.New("product_id is not valid")
	}
	totalPrice := product.Price * addCartProduct.Quantity

	sqlStatement = `UPDATE cart_product SET quantity=$3, total_price=$4  WHERE cart_id=$1 AND id=$2;`
	
	result, err := cs.DB.Exec(sqlStatement, cartId, cartProductId, addCartProduct.Quantity, totalPrice)
	if err != nil {
		e := fmt.Sprintf("error: %v occurred while updating cart_product record with id: %d", err, cartProductId)
		return 0, errors.New(e)
	}
	count, err := result.RowsAffected()
	if err != nil {
		e := fmt.Sprintf("error occurred from database after update data: %v", err)
		return 0, errors.New(e) 
	}

	if count == 0 {
		e := "could not update the cart_product, please try again after sometime"
		return 0, errors.New(e) 
	}
	return int(count), nil
}

func (cs *CartServiceImpl) DeleteCartProductById(cartId, cartProductId int) (int, error) {
	sqlStatement = `DELETE FROM cart_product WHERE cart_id=$1 AND id=$2;`
	res, err := cs.DB.Exec(sqlStatement, cartId, cartProductId)
	if err != nil {
		e := fmt.Sprintf("error: %v occurred while delete cart_product record with id: %d", err, cartProductId)
		return 0, errors.New(e)
	}
	count, err := res.RowsAffected()
	if err != nil {
		e := fmt.Sprintf("error occurred from database after delete data: %v", err)
		return 0, errors.New(e)		
	}

	if count == 0 {
		e := fmt.Sprintf("could not delete the cart_product, there might be no data for ID %d", cartProductId)
		return 0, errors.New(e) 
	}
	return int(count), nil
}