package services

import (
	"mycart/app/helpers/databasehelper"
	"mycart/app/helpers/loghelper"
	"mycart/app/models"
)

// AddCategory -
func AddCategory(name string) (*models.Category, error) {

	db, err := databasehelper.GetConnectionByHostName("")
	if err != nil {
		loghelper.LogError(err)
		return nil, err
	}
	category := &models.Category{
		Name: name,
	}

	err = db.Create(category).Error
	if err != nil {
		loghelper.LogError(err)
		return nil, err
	}
	return category, nil
}

// GetAllCategories -
func GetAllCategories() ([]models.Category, error) {

	db, err := databasehelper.GetConnectionByHostName("")
	if err != nil {
		loghelper.LogError(err)
		return nil, err
	}
	var categories []models.Category
	err = db.Find(&categories).Error
	if err != nil {
		loghelper.LogError(err)
		return nil, err
	}
	return categories, nil
}
