package main

import (
	"fmt"
	"os"
	"project/controller"
	"project/database"
	"project/middleware"
	"project/service"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// ENV Configuration
	err := godotenv.Load("config/.env")
	if err != nil {
		fmt.Println("failed load file environtment")
	} else {
		fmt.Println("success read file environtment")
	}

	router := gin.Default()
	DB := database.NewConnDB()

	// Supplier
	supplierService := service.NewSupplierService(DB)
	supplierController := controller.NewSupplierController(supplierService)
	router.POST("/suppliers", middleware.WithAuh(), middleware.WithAuh(), supplierController.AddSupplier)
	router.GET("/suppliers", supplierController.GetSuppliers)
	router.GET("/suppliers/:id", supplierController.GetSupplierById)
	router.PUT("/suppliers/:id", middleware.WithAuh(), supplierController.UpdateSupplierById)
	router.DELETE("/suppliers/:id", middleware.WithAuh(), supplierController.DeleteSupplierById)

	// Category
	categoryService := service.NewCategorySevice(DB)
	categoryController := controller.NewCategoryController(categoryService)
	router.POST("/categories", middleware.WithAuh(), categoryController.AddCategory)
	router.GET("/categories", categoryController.GetCategories)
	router.GET("/categories/:id", categoryController.GetCategoryById)
	router.PUT("/categories/:id", middleware.WithAuh(), categoryController.UpdateCategoryById)
	router.DELETE("/categories/:id", middleware.WithAuh(), categoryController.DeleteCategoryById)

	// Product
	productService := service.NewProductService(DB)
	productController := controller.NewProductController(productService)
	router.POST("/products", middleware.WithAuh(), productController.AddProduct)
	router.GET("/products", productController.GetProducts)
	router.GET("/products/:id", productController.GetProductById)
	router.PUT("/products/:id", middleware.WithAuh(), productController.UpdateProductById)
	router.DELETE("/products/:id", middleware.WithAuh(), productController.DeleteProductById)

	// Customer
	customerService := service.NewCustomerService(DB)
	customerController := controller.NewCustomerController(customerService)
	router.POST("customer/register", customerController.RegisterCustomer)
	router.POST("customer/login", customerController.LoginCustomer)
	router.PUT("customer/update", customerController.UpdateCustomer)
	router.DELETE("customer/delete/:id", customerController.DeleteCustomer)
	
	// Feedback
	feedbackService := service.NewFeedbackService(DB)
	feedbackController := controller.NewFeedbackController(feedbackService)
	router.POST("/reviews", middleware.WithAuh(), feedbackController.AddFeedback)
	router.GET("/reviews", feedbackController.GetFeedbacks)
	router.GET("/reviews/:id", feedbackController.GetFeedbackById)
	router.PUT("/reviews/:id", middleware.WithAuh(), feedbackController.UpdateFeedbackById)
	router.DELETE("/reviews/:id", middleware.WithAuh(), feedbackController.DeleteFeedbackById)

	// PaymentType
	paymentTypeService := service.NewPaymentTypeService(DB)
	paymentTypeController := controller.NewPaymentTypeController(paymentTypeService)
	router.POST("/paymentType", middleware.WithAuh(), paymentTypeController.AddPaymentType)
	router.GET("/paymentType", paymentTypeController.GetPaymentTypes)
	router.GET("/paymentType/:id", paymentTypeController.GetPaymentTypeById)
	router.PUT("/paymentType/:id", middleware.WithAuh(), paymentTypeController.UpdatePaymentTypeById)
	router.DELETE("/paymentType/:id", middleware.WithAuh(), paymentTypeController.DeletePaymentTypeById)
	
	// Cart
	cartService := service.NewCartService(DB)
	cartController := controller.NewCartController(cartService)
	router.POST("/cart", middleware.WithAuh(), cartController.AddCart)
	router.GET("/cart", middleware.WithAuh(), cartController.GetCartByCustomerId)
	router.PUT("/cart", middleware.WithAuh(), cartController.UpdateCartProductById)
	router.DELETE("/cart", middleware.WithAuh(), cartController.DeleteCartProductById)

	// Transaction
	transactionService := service.NewTransactionService(DB)
	transactionController := controller.NewTransactionController(transactionService)
	router.POST("/transactions", middleware.WithAuh(), transactionController.AddTransaction)
	router.GET("/transactions/:id", middleware.WithAuh(), transactionController.GetTransaction)

	
	router.Run(":" + os.Getenv("PORT"))
}
