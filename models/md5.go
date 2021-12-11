package models

import (
	"context"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"fmt"
	"sync"

	"github.com/go-redis/redis/v8"
)

func MD5Subscriber(wg *sync.WaitGroup, ctx context.Context, cache *redis.Client, db *sql.DB) {
	defer wg.Done()
	subscriber := cache.Subscribe(ctx, "passwords")
	var wg2 sync.WaitGroup
	defer wg2.Wait()
	for {
		msg, err := subscriber.ReceiveMessage(ctx)
		if err != nil {
			panic(err)
		}
		password := msg.Payload
		fmt.Printf("received : md5 : %s\n", password)
		hash := md5.Sum([]byte(password))
		hashed := hex.EncodeToString(hash[:])
		wg2.Add(1)
		go Save(&wg2, "md5", password, hashed, db)
	}
}
