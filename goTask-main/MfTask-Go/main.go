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
	"github.com/gofiber/fiber/v2/middleware/cors"
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
	// cfg := config.Load()
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

	rdb := redis.ConnRedis()

	Init(db)
	go redis.PubNav(rdb, wg)

	ctx := context.Background()

	msgOrderPlaced := mesagging.NewMessaging("orders.placed", strings.Split(SEEDS, ","))
	msgOrderConfirmed := mesagging.NewMessaging("order.confirmed", strings.Split(SEEDS, ","))
	go msgOrderPlaced.ProduceRecords()
	go msgOrderConfirmed.ProduceRecords()
	go consumer.ConsumeTopic(db, msgOrderConfirmed)

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173", // your frontend URL
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
	}))

	ODHandler := handler.NewOrderHandler(database.NewOrderDB(db))

	OR_group := app.Group("/api/v1/orders")
	OR_group.Post("/", ODHandler.Place_order(msgOrderPlaced, rdb, ctx))
	OR_group.Get("/:order_id", ODHandler.Get_Order_By_ID)
	// app.Post("/login", handler.LoginHandler(cfg))

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
