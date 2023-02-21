package database

import (
	"os"

	"github.com/redis/go-redis/v9"
)

var db *redis.Client

func Connect() *redis.Client {

	if db == nil {
		redis.SetLogger(nil)
		db = redis.NewClient(&redis.Options{
			Addr:     os.Getenv("REDIS_URL"),
			Password: os.Getenv("REDIS_PASSWORD"),
		})
	}

	return db
}
