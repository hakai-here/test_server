package cache

import (
	"demoproject/api/models"
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"
)

var rdb *redis.Client

func TodoSetKey(key string, value []models.TodoActivity, time time.Duration) error {
	jsondata, err := json.Marshal(value) // marshelling the data to json format to store in redisdb
	if err != nil {
		return err
	}
	if err := rdb.Set(ctx, key, jsondata, time).Err(); err != nil {
		return err
	}
	return nil
}

func TodoGetKey(key string) ([]models.TodoActivity, error) {
	var data []models.TodoActivity
	value, err := rdb.Get(ctx, key).Result()
	if err != nil {
		return data, err
	}

	if err := json.Unmarshal([]byte(value), &data); err != nil { // unmarshelling the data for validating the session
		return data, err
	}
	return data, err
}

func DeleteKey(key string) error {
	if _, err := rdb.Del(ctx, key).Result(); err != nil {
		return err
	}
	return nil
}
