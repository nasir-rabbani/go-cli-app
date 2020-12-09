package services

import (
	"mycart/app/helpers/databasehelper"
	"mycart/app/helpers/loghelper"
	"mycart/app/models"
)

// RegisterUser -
func RegisterUser(name, password, role string) (*models.User, error) {

	db, err := databasehelper.GetConnectionByHostName("")
	if err != nil {
		loghelper.LogError(err)
		return nil, err
	}
	user := &models.User{
		Name:     name,
		Password: password,
		Role:     role,
	}

	err = db.Create(user).Error
	if err != nil {
		loghelper.LogError(err)
		return nil, err
	}
	return user, nil
}
