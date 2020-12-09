package services

import (
	"mycart/app/helpers/databasehelper"
	"mycart/app/helpers/loghelper"
	"mycart/app/models"
)

// AddProduct - TO add a new product
func AddProduct(name string, price float32, categoryID uint) (*models.Product, error) {

	db, err := databasehelper.GetConnectionByHostName("")
	if err != nil {
		loghelper.LogError(err)
		return nil, err
	}
	product := &models.Product{
		Name:       name,
		Price:      price,
		CategoryID: categoryID,
	}

	err = db.Create(product).Error
	if err != nil {
		loghelper.LogError(err)
		return nil, err
	}
	return product, nil
}

// GetAllProducts - To get all the products in the database
func GetAllProducts() ([]models.Product, error) {

	db, err := databasehelper.GetConnectionByHostName("")
	if err != nil {
		loghelper.LogError(err)
		return nil, err
	}
	var products []models.Product
	err = db.Find(&products).Error
	if err != nil {
		loghelper.LogError(err)
		return nil, err
	}
	return products, nil
}

// GetProductsByCategoryID - To get all the products of a certain category
func GetProductsByCategoryID(ID uint) ([]models.Product, error) {

	db, err := databasehelper.GetConnectionByHostName("")
	if err != nil {
		loghelper.LogError(err)
		return nil, err
	}
	var products []models.Product
	// err = db.Find(&products).Error                           // fix query
	err = db.Where("category_id = ?", ID).Find(&products).Error // fix query
	if err != nil {
		loghelper.LogError(err)
		return nil, err
	}
	return products, nil
}

// GetProductByID - To get the product by ID
func GetProductByID(ID uint) (models.Product, error) {

	db, err := databasehelper.GetConnectionByHostName("")
	if err != nil {
		loghelper.LogError(err)
		return models.Product{}, err
	}
	var product models.Product
	err = db.First(&product, ID).Error
	if err != nil {
		loghelper.LogError(err)
		return models.Product{}, err
	}
	return product, nil
}
