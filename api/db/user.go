package db

import (
	"demoproject/api/models"

	"gorm.io/gorm"
)

func CheckUsername(username string) (bool, error) { // check if username already exists in db
	var signup models.User
	result := db.Where("username = ?", username).First(&signup)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return false, result.Error
	}

	return result.RowsAffected > 0, nil
}

func InsertUser(userDetials models.User) error {
	if err := db.Create(userDetials).Error; err != nil { // inserting data to the postgres database
		return err
	}
	return nil
}

func GetUserDetials(username string) (models.User, error) {
	var data models.User
	if err := db.First(&data, "username = ?", username).Error; err != nil {
		return data, err
	}
	return data, nil
}
