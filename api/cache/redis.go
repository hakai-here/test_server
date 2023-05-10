package cache

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

var rdb *redis.Client
var ctx = context.Background()

func InitRedis() (string, error) { // initilizing Redis DB
	if rdb != nil {
		return "Already present", nil
	}
	rdb = redis.NewClient(&redis.Options{
		Addr:     viper.GetString("REDIS_URL"),
		Password: "",
		DB:       0, // use default DB
	})
	if _, err := rdb.Ping(ctx).Result(); err != nil {
		return "", err
	}

	return "connect successfully", nil
}

func GetValue(key string) (string, error) { // getting data from the cache database
	val, err := rdb.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", nil
	}
	return val, err
}

func SetValue(key string, value []byte, time time.Duration) error { // setting data to the database for cacheing
	if err := rdb.Set(ctx, key, value, time).Err(); err != nil {
		return err
	}
	return nil
}
