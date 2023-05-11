package cache

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

var ctx = context.Background()
var err error

func createConnections(rdb *redis.Client, Db int) (*redis.Client, error) {
	if rdb != nil {
		return rdb, nil
	}

	rdb = redis.NewClient(&redis.Options{
		Addr:     viper.GetString("REDIS_URL"),
		Password: "",
		DB:       Db,
	})
	if _, err = rdb.Ping(ctx).Result(); err != nil {
		return rdb, err

	}
	return rdb, nil
}

func InitConnection() error {
	rdbauth, err = createConnections(rdbauth, 0) // creating connection to sessions(auth) db
	if err != nil {
		return fmt.Errorf("unable to connect to auth cache database : %s", err.Error())
	}

	rdb, err = createConnections(rdb, 1) // creating connection to cache db
	if err != nil {
		return fmt.Errorf("unable to connect to auth cache database : %s", err.Error())
	}
	return nil
}
