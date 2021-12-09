package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"github.com/rapando/hash-engine/md5engine"
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
	base.ctx = context.Background()
	var dbURI = "sam:therealsam@tcp(localhost:3306)/hash_engine_db"
	base.db, err = utils.DbConnect(dbURI)
	if err != nil {
		panic("db error")
	}

	base.cache, _ = utils.ConnectToCache(base.ctx, "localhost:6379", "", 0)

	var chars = []rune("`1234567890-=\\][poiuytrewqasdfghjkl;'/.," +
		"mnbvcxz~!@#$%^&*()_+|}{POIUYTREWQASDFGHJKL:\"?><MNBVCXZ")
	var n = len(chars)

	var stopChan = make(chan bool)
	for i := 1; i <= n; i++ {
		section := chars[0:i]
		go permute(base.ctx, section, 0, len(section)-1, base.cache)
	}
	go md5engine.MD5Subscriber(base.ctx, base.cache, base.db)
	<-stopChan
}

func permute(ctx context.Context, a []rune, l, r int, cache *redis.Client) {
	var s string
	if l == r {
		s = string(a)
		fmt.Println(s)
		_ = cache.Publish(ctx, "passwords", s)
	} else {
		for i := l; i <= r; i++ {
			temp := a[l]
			a[l] = a[i]
			a[i] = temp

			permute(ctx, a, l+1, r, cache)

			temp = a[l]
			a[l] = a[i]
			a[i] = temp
		}
	}
}
