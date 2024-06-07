package main

import (
	"context"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()
var rdb *redis.Client

// initializeRedis initializes the Redis client
func initializeRedis() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // No password set
		DB:       0,  // Use default DB
	})

	// Ping Redis to check the connection
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatal(err)
	}
}

// setKey sets a key-value pair in Redis with expiration time
func setKey(key string, value string, expiration time.Duration) {
	err := rdb.Set(ctx, key, value, expiration)
	if err.Err() != nil {
		log.Fatal(err.Err())
	}
}

// getKey retrieves the value associated with the given key from Redis
func getKey(key string) (string, error) {
	val, err := rdb.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}
