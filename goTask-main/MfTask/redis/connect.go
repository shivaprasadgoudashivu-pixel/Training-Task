package redis

import (
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

func ConnRedis() (rdb *redis.Client) {

	rdb = redis.NewClient(&redis.Options{
		Addr: "redis:6379",

		DialTimeout:  2 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
		PoolSize:     20,
		MinIdleConns: 4,
	})

	if err := rdb.Ping(ctx).Err(); err != nil {
		log.Print(err)
	} else {
		println("redis Connected ..")
	}

	return rdb

}
