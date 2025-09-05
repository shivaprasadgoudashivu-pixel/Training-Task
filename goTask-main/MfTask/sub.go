package main

// import (
// 	"context"
// 	"fmt"
// 	"log"

// 	"github.com/redis/go-redis/v9"
// )

// var Sl = []string{"ABB", "ACC", "ABB-EQ"}

// func subNav() {
// 	ctx := context.Background()
// 	rdb := redis.NewClient(&redis.Options{Addr: "localhost:6379"})
// 	defer rdb.Close()

// 	// Subscribe to specific channels
// 	pubsub := rdb.PSubscribe(ctx, "__keyspace@0__:nav:latest:*")
// 	// pubsub := rdb.Subscribe(ctx, "ABB", "ACC", "ABB-EQ")
// 	defer pubsub.Close()

// 	// Confirm subscription
// 	_, err := pubsub.Receive(ctx)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// ch := pubsub.Channel()

// 	fmt.Println("Listening for updates...")
// 	for msg := range pubsub.Channel() {
// 		// msg.Channel = "__keyspace@0__:nav:latest:ABB"
// 		// msg.Payload = "set"
// 		key := msg.Channel[len("__keyspace@0__:"):] // strip prefix
// 		if msg.Payload == "set" {
// 			val, err := rdb.Get(ctx, key).Result()
// 			if err != nil {
// 				log.Println("GET failed:", err)
// 			} else {
// 				fmt.Printf("Update %s → %s\n", key, val)
// 			}
// 		}
// 	}
// }
