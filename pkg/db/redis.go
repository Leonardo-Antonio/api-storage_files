package db

import (
	"github.com/go-redis/redis/v8"
)

func RedisGetClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}
