package main

import (
	"context"
	"keycloak-demo/database"
	"keycloak-demo/handler"
	"keycloak-demo/kafka/consumer"
	mesagging "keycloak-demo/kafka/messaging"
	"keycloak-demo/model"
	"keycloak-demo/redis"
	"os"
	"runtime"
	"strings"
	"sync"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

var (
	DSN           string
	PORT          string
	SEEDS         string
	Topic         string
	ConsumerGroup string
)

// var wg *sync.WaitGroup

func main() {

	service := "mf-service"
	wg := new(sync.WaitGroup)
	if SEEDS == "" {
		SEEDS = "kafka1, kafka2, kafka3"
	}

	DSN = os.Getenv("DSN")
	if DSN == "" {
		DSN = `host=localhost user=app password=app123 dbname=usersdb port=5432 sslmode=disable`
		log.Info().Msg(DSN)
	}
	PORT = os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	db, err := database.GetConnection(DSN)

	if err != nil {
		log.Fatal().
			Err(err).
			Str("service", service).
			Msgf("unable to connect to the database %s", service)
	}
	log.Info().Str("service", service).Msg("database connection is established")

	//redis connection
	// rdb := redis.NewClient(&redis.Options{
	// 	Addr: "redis:6379", // or your VM IP if not forwarded

	// 	DialTimeout:  2 * time.Second,
	// 	ReadTimeout:  1 * time.Second,
	// 	WriteTimeout: 1 * time.Second,
	// 	PoolSize:     20,
	// 	MinIdleConns: 4,
	// })
	// defer rdb.Close()

	// if err := rdb.Ping(ctx).Err(); err != nil {
	// 	log.Print(err)
	// } else {
	// 	println("redis Connected ..")
	// }

	rdb := redis.ConnRedis()

	Init(db)
	go redis.PubNav(rdb, wg)

	ctx := context.Background()

	msgOrderPlaced := mesagging.NewMessaging("ordersv1", strings.Split(SEEDS, ","))
	go msgOrderPlaced.ProduceRecords()
	go consumer.ConsumeTopic(db)

	app := fiber.New()

	ODHandler := handler.NewOrderHandler(database.NewOrderDB(db))

	OR_group := app.Group("/api/v1/orders")
	OR_group.Post("/", ODHandler.Place_order(msgOrderPlaced, rdb, ctx))
	OR_group.Get("/:order_id", ODHandler.Get_Order_By_ID)

	err = app.Listen(":" + PORT)

	// err = http.ListenAndServe(":"+PORT, nil)

	if err != nil {
		println(err.Error())
		runtime.Goexit()
	}

	wg.Wait()
}

func Init(db *gorm.DB) {
	db.AutoMigrate(&model.ORDER{}, &model.HOLDINGS{})
}
