package cache

import (
	"demoproject/api/models"
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"
)

var rdbauth *redis.Client

func AuthSetKey(key string, value models.RedisType, time time.Duration) error {
	jsondata, err := json.Marshal(value) // marshelling the data to json format to store in redisdb
	if err != nil {
		return err
	}
	if err := rdbauth.Set(ctx, key, jsondata, time).Err(); err != nil {
		return err
	}
	return nil
}

func AuthGetKey(key string) (models.RedisType, error) {
	var data models.RedisType
	value, err := rdbauth.Get(ctx, key).Result()
	if err != nil {
		return data, err
	}

	if err := json.Unmarshal([]byte(value), &data); err != nil { // unmarshelling the data for validating the session
		return data, err
	}
	return data, err
}
