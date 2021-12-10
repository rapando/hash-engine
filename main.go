package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
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

func main() {
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

	base.cache, _ = utils.ConnectToCache(base.ctx, "localhost:6379", "", 0)

	var chars = []rune("`1234567890-=\\][poiuytrewqasdfghjkl;'/.," +
		"mnbvcxz~!@#$%^&*()_+|}{POIUYTREWQASDFGHJKL:\"?><MNBVCXZ")
	var n = len(chars)

	// var stopChan = make(chan bool)
	// for i := 1; i <= n; i++ {
	// 	section := chars[0:i]
	// 	go permute(base.ctx, section, 0, len(section)-1, base.cache)
	// }
	// go md5engine.MD5Subscriber(base.ctx, base.cache, base.db)
	// <-stopChan
}

func getNoOfPossibilities(length int) (possibilities int) {
	if length == 1 {
		return 1
	}

	for i := 1; i < length; i++ {
		possibilities += length * factorial(i)
	}
	possibilities += factorial(length)
	return possibilities
}

func factorial(n int) (f int) {
	f = 1
	for i := n; i >= 1; i-- {
		f *= n
	}
	return f
}
func differentFlagPermutations()
