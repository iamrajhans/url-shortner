package utils

import (
	"os"

	"github.com/redis/go-redis/v9"
)

func GetRedisClient() *redis.Client {
	redisClient := redis.NewClient(&redis.Options{

		Addr:     os.Getenv("REDIS_HOST"), // e.g., "localhost:6379"
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})
	return redisClient
}
