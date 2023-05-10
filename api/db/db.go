package db

import (
	"demoproject/api/cache"
	"demoproject/api/structs"
	"encoding/json"
	"fmt"
	"time"

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
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{}) // connecting postgres database
	if err != nil {
		return err
	}
	db.AutoMigrate(&structs.Proceedingentry{})
	return nil
}

func InsertMany(data []structs.Proceedingentry) { // insert data to database <only used  after xml parsing>
	db.Create(data)
}

func GetQueriedEntry(query string) ([]structs.Proceedingentry, error) { // getting the specific value from the database
	value, err := cache.GetValue(query) // primarly it will search in cache and if not found will go to the nain database
	var entry structs.Proceedingentry
	var entries []structs.Proceedingentry
	if value == "" && err == nil {
		db.First(&entry, "number = ?", query)
		if entry.Number == 0 {
			return entries, fmt.Errorf("no data found")
		}
		entries = append(entries, entry)
		jsonData, err := json.Marshal(entries) // marshelling the data
		if err != nil {
			return entries, fmt.Errorf("unable to marshal")
		}
		cache.SetValue(query, jsonData, time.Hour) // adding the data to redis database
		return entries, nil
	} else if err != nil {
		return entries, err
	}

	if err := json.Unmarshal([]byte(value), &entries); err != nil { // unmarshelling the data collected from cache
		return entries, fmt.Errorf("unable to unmarshal the data")
	}

	return entries, nil
}

func GetAllEntrys() ([]structs.Proceedingentry, error) { // getting all value from database

	value, err := cache.GetValue("/") // checking the redis db

	var entries []structs.Proceedingentry
	if err == nil && value == "" {

		db.Find(&entries)
		jsonData, err := json.Marshal(entries) // marshelling the data for cache
		if err != nil {
			fmt.Println("Failed to marshal JSON:", err)
		}
		if err := cache.SetValue("/", jsonData, time.Hour); err != nil { // setting the cache
			return []structs.Proceedingentry{}, nil
		}
		return []structs.Proceedingentry{}, fmt.Errorf("no data found")
	}
	if err := json.Unmarshal([]byte(value), &entries); err != nil { // unmarshelling the data collected from cache
		return []structs.Proceedingentry{}, fmt.Errorf("unable to unmarshal")
	}
	return entries, nil
}
