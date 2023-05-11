package db

import (
	"demoproject/api/models"
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB = nil
var err error

func ConnectDB() error {
	username := viper.GetString("DB_USER") // get values from .env
	password := viper.GetString("DB_PASS")
	dbName := viper.GetString("DB_NAME")
	dbPort := viper.GetString("DB_PORT")
	dbHost := viper.GetString("DB_HOST")
	dsn := fmt.Sprintf("host=%s user=%s dbname=%s port=%s sslmode=disable password=%s", dbHost, username, dbName, dbPort, password)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{SkipDefaultTransaction: true}) // connecting postgres database , transaction skiped
	if err != nil {
		return err
	}
	// migrating the databases
	db.AutoMigrate(models.User{}, models.TodoActivity{})

	return nil
}
