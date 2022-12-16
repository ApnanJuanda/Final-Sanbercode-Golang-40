package service

import (
	"database/sql"
	"errors"
	"fmt"
	"project/helper"
	"project/model"
)

type ProductServiceImpl struct {
	DB *sql.DB
}

func NewProductService(DB *sql.DB) ProductService {
	return &ProductServiceImpl{
		DB: DB,
	}
}

func (ps *ProductServiceImpl) AddProduct(addProduct *model.AddProduct) (*model.Product, error) {
	//validasi category_id and supplier_id
	categoryService := NewCategorySevice(ps.DB)
	category, _ := categoryService.GetCategoryById(addProduct.CategoryId)
	if category == nil {
		return nil, errors.New("category_id is not valid")
	}

	supplierService := NewSupplierService(ps.DB)
	supplier, _ := supplierService.GetSupplierById(addProduct.SupplierId)
	if supplier == nil {
		return nil, errors.New("supplier_id is not valid")
	}

	//populator
	product, err := helper.ProductPopulator(addProduct)
	if err != nil {
		return nil, err
	}

	//save query
	var newProduct = model.Product{}
	sqlStatement = `INSERT INTO product (name, price, stock_product, image_url, supplier_id, category_id )
	VALUES ($1, $2, $3, $4, $5, $6)
	Returning *
	`
	err = ps.DB.QueryRow(sqlStatement, product.Name, product.Price, product.StockProduct, product.ImageUrl, product.SupplierId, product.CategoryId).
		Scan(&newProduct.ID, &newProduct.Name, &newProduct.Price, &newProduct.StockProduct, &newProduct.CreatedAt, &newProduct.UpdatedAt,
			&newProduct.CategoryId, &newProduct.SupplierId, &newProduct.ImageUrl)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	return &newProduct, nil

}

func (ps *ProductServiceImpl) GetProducts() (*[]model.Product, error) {
	var products = []model.Product{}

	sqlStatement := `SELECT * FROM product`

	rows, err := ps.DB.Query(sqlStatement)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var product = model.Product{}

		err = rows.Scan(&product.ID, &product.Name, &product.Price, &product.StockProduct, &product.CreatedAt, &product.UpdatedAt,
			&product.CategoryId, &product.SupplierId, &product.ImageUrl)

		if err != nil {
			return nil, err
		}

		products = append(products, product)

	}
	return &products, nil
}

func (ps *ProductServiceImpl) GetProductById(id int) (*model.Product, error) {
	var product = model.Product{}
	sqlStatement := `SELECT * FROM product WHERE id=($1)`
	err = ps.DB.QueryRow(sqlStatement, id).Scan(&product.ID, &product.Name, &product.Price, &product.StockProduct, &product.CreatedAt, &product.UpdatedAt,
		&product.CategoryId, &product.SupplierId, &product.ImageUrl)
	if err != nil {
		return nil, err
	}
	return &product, err
}

func (ps *ProductServiceImpl) UpdateProductById(id int, addProduct *model.AddProduct) (int, error) {
	//validasi category_id and supplier_id
	categoryService := NewCategorySevice(ps.DB)
	category, _ := categoryService.GetCategoryById(addProduct.CategoryId)
	if category == nil {
		return 0, errors.New("category_id is not valid")
	}

	supplierService := NewSupplierService(ps.DB)
	supplier, _ := supplierService.GetSupplierById(addProduct.SupplierId)
	if supplier == nil {
		return 0, errors.New("supplier_id is not valid")
	}

	//populator
	product, err := helper.ProductPopulator(addProduct)
	if err != nil {
		return 0, err
	}

	//save query
	sqlStatement := `UPDATE product SET name = $2, price = $3,  stock_product = $4, image_url = $5, supplier_id = $6, category_id = $7
	WHERE id = $1;`
	result, err := ps.DB.Exec(sqlStatement, id, product.Name, product.Price, product.StockProduct, product.ImageUrl, product.SupplierId, product.CategoryId)

	if err != nil {
		e := fmt.Sprintf("error: %v occurred while updating product record with id: %d", err, id)
		return 0, errors.New(e)
	}
	count, err := result.RowsAffected()
	if err != nil {
		e := fmt.Sprintf("error occurred from database after update data: %v", err)
		return 0, errors.New(e) 
	}

	if count == 0 {
		e := "could not update the product, please try again after sometime"
		return 0, errors.New(e) 
	}
	return int(count), nil
}

func (ps *ProductServiceImpl) DeleteProductById(id int) (int, error) {
	sqlStatement = `DELETE FROM product WHERE id=$1;`
	res, err := ps.DB.Exec(sqlStatement, id)
	if err != nil {
		e := fmt.Sprintf("error: %v occurred while delete product record with id: %d", err, id)
		return 0, errors.New(e)
	}
	count, err := res.RowsAffected()
	if err != nil {
		e := fmt.Sprintf("error occurred from database after delete data: %v", err)
		return 0, errors.New(e)		
	}

	if count == 0 {
		e := fmt.Sprintf("could not delete the product, there might be no data for ID %d", id)
		return 0, errors.New(e) 
	}
	return int(count), nil
}