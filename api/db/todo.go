package db

import (
	"demoproject/api/cache"
	"demoproject/api/models"
	"log"
	"time"
)

func InsertTodo(data models.TodoActivity) error {
	return db.Create(data).Error // inserting the data to the database
}

func GetTodo(userid string, query string) ([]models.TodoActivity, error) {
	var entries []models.TodoActivity
	if query != "" {
		if err := db.Find(&entries, "user_id = ? AND id = ?", userid, query).Error; err != nil {
			return entries, err
		}
		if len(entries) > 0 {
			if err := cache.TodoSetKey(query, entries, 2*time.Hour); err != nil {
				log.Println(err.Error()) // unable to cache . but will not stop the server
			}
		}
	} else {
		if err := db.Find(&entries, "user_id = ?", userid).Error; err != nil {
			return entries, err
		}
	}

	return entries, nil
}

func Deletetodo(id string) error {
	return db.Delete(&models.TodoActivity{}, "id = ?", id).Error
}

func Update(data models.TodoActivity) error {
	return db.Save(data).Error
}
