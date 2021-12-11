package models

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/rapando/hash-engine/utils"
)

type base struct {
	ctx   context.Context
	db    *sql.DB
	cache *redis.Client
}

func Process() {
	start := time.Now()
	defer func() {
		fmt.Printf("\nDONE : %s\n", time.Since(start))
	}()
	var err error
	var base base

	err = godotenv.Load()
	if err != nil {
		log.Printf("unable to load dotenv because %v", err)
		return
	}

	base.ctx = context.Background()
	base.db, err = utils.DbConnect(os.Getenv("DB_URI"))
	if err != nil {
		panic("db error")
	}

	base.cache, _ = utils.ConnectToCache(base.ctx,
		os.Getenv("REDIS_HOST"),
		os.Getenv("REDIS_PASSWORD"),
		0)

	var chars = "`1234567890-=\\][poiuytrewqasdfghjkl;'/.," +
		"mnbvcxz~!@#$%^&*()_+|}{POIUYTREWQASDFGHJKL:\"?><MNBVCXZ "
	chars = "abcde"
	var n = int64(len(chars))

	combinations := GetNoOfCombinations(n)
	log.Printf("a string with %d characters has %d combinations", n, combinations.Int64())

	var wg sync.WaitGroup
	wg.Add(2)
	log.Println("getting combinations and pusblishing them to redis")

	go GetCombinations(&wg, base.ctx, chars, base.cache)
	log.Println("done publishing combinations")

	/// subscribe
	go MD5Subscriber(&wg, base.ctx, base.cache, base.db)
	wg.Wait()
}
