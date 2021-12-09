package md5engine

import (
	"context"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"fmt"
	"sync"

	"github.com/go-redis/redis/v8"
	"github.com/rapando/hash-engine/database"
)

func MD5Subscriber(ctx context.Context, cache *redis.Client, db *sql.DB) {
	subscriber := cache.Subscribe(ctx, "passwords")
	var wg sync.WaitGroup
	defer wg.Wait()
	for {
		msg, err := subscriber.ReceiveMessage(ctx)
		if err != nil {
			panic(err)
		}
		password := msg.Payload
		fmt.Printf("received : md5 : %s\n", password)
		hash := md5.Sum([]byte(password))
		hashed := hex.EncodeToString(hash[:])
		wg.Add(1)
		go database.Save(&wg, "md5", password, hashed, db)
	}
}
