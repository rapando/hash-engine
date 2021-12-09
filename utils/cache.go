package utils

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

// ConnectToCache
// connects to single redis instance
// Takes context, host, password and db
// Returns client, error
func ConnectToCache(ctx context.Context, host, password string, db int) (
	client *redis.Client, err error) {
	log.Printf("connecting to redis")
	client = redis.NewClient(&redis.Options{
		Addr:     host,
		Password: password,
		DB:       db,
		PoolSize: 100,
	})
	if err = client.Ping(ctx).Err(); err != nil {
		log.Printf("unable tp ping redis because %v", err)
		return nil, err
	}
	log.Printf("connected to redis")
	return client, nil
}

func Publish(ctx context.Context, data string, client *redis.Client) (err error) {
	return client.Publish(ctx, "passwords", data).Err()
}
