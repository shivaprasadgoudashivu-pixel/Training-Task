package redis

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

// var wg *sync.WaitGroup

func PubNav(rdb *redis.Client, wg *sync.WaitGroup) {
	// var Sl = []string{"ABB", "ACC", "ABB-EQ"}
	// wg = new(sync.WaitGroup)
	// rdb := redis.NewClient(&redis.Options{
	// 	Addr: "localhost:6379", // or your VM IP if not forwarded

	// 	DialTimeout:  2 * time.Second,
	// 	ReadTimeout:  1 * time.Second,
	// 	WriteTimeout: 1 * time.Second,
	// 	PoolSize:     20,
	// 	MinIdleConns: 4,
	// })
	// defer rdb.Close()

	// if err := rdb.Ping(ctx).Err(); err != nil {
	// 	log.Fatal(err)
	// } else {
	// 	println("Connected")
	// }

	// wg.Add(1)
	go func() {
		// for i := 0; i < 3; i++ {
		for {
			time.Sleep(time.Second * 3)
			// key := fmt.Sprintf("nav:latest:%s", Sl[i])
			val := rand.Intn(100)
			err := rdb.Set(ctx, "ABB", val, 0).Err()
			if err != nil {
				log.Println("SET failed:", err)
			} else {
				fmt.Printf("Updated %s → %d\n", "ABB", val)
			}
			// if i == len(Sl)-1 {
			// 	i = 0
			// }

		}
		// wg.Done()
	}()
	// wg.Wait()

}
