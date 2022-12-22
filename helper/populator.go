package helper

import (
	"errors"
	"project/model"

	"golang.org/x/crypto/bcrypt"
)

func ProductPopulator(addProduct *model.AddProduct) (*model.Product, error) {
	var imageValidated string
	if addProduct.Name == "" {
		return nil, errors.New("name is required")
	}
	if addProduct.Price <= 0 {
		return nil, errors.New("price must greater than zero")
	}
	if addProduct.StockProduct <= 0 {
		return nil, errors.New("stock_product must greater than zero")
	}
	isValid, _ := ValidateImageUrl(addProduct.ImageUrl)
	if isValid {
		imageValidated = addProduct.ImageUrl
	} else {
		return nil, errors.New("image_url is not valid")
	}
	return &model.Product{
		Name:         addProduct.Name,
		Price:        addProduct.Price,
		StockProduct: addProduct.StockProduct,
		ImageUrl:     imageValidated,
		CategoryId:   addProduct.CategoryId,
		SupplierId:   addProduct.SupplierId,
	}, nil

}

func CustomerPopulator(customerRegister *model.CustomerRegister) (*model.Customer, error) {
	var hash []byte
	if customerRegister.Name == "" {
		return nil, errors.New("name is required")
	}

	if customerRegister.Email == "" {
		return nil, errors.New("email is required")
	}

	if customerRegister.Password == "" {
		return nil, errors.New("password is required")
	}

	if len(customerRegister.Password) < 7 {
		return nil, errors.New("password length must be 6 character or more")
	} else {
		hash, _ = bcrypt.GenerateFromPassword([]byte(customerRegister.Password), bcrypt.DefaultCost)
	}

	if customerRegister.PhoneNumber == "" {
		return nil, errors.New("phone_number is required")
	}

	if customerRegister.Address == "" {
		return nil, errors.New("address is required")
	}

	return &model.Customer{
		Name:        customerRegister.Name,
		Email:       customerRegister.Email,
		Password:    string(hash),
		PhoneNumber: customerRegister.PhoneNumber,
		Address:     customerRegister.Address,
	}, nil
}

func FeedbackPopulator(addFeedback *model.AddFeedback, customerId int) (*model.Feedback, error) {
	if addFeedback.Review == "" {
		return nil, errors.New("review is required")
	}

	return &model.Feedback{
		CustomerId: customerId,
		ProductId:  addFeedback.ProductId,
		Review:     addFeedback.Review,
	}, nil
}

func CartProductPopulator(addCartProduct *model.AddCartProduct) (*model.CartProduct, error) {
	if addCartProduct.Quantity <= 0 {
		return nil, errors.New("quantity must greater than zero")
	}
	return &model.CartProduct{
		ProductId: addCartProduct.ProductId,
		Quantity:  addCartProduct.Quantity,
	}, nil
}
