package handler

import (
	"context"
	"errors"
	"keycloak-demo/database"
	"keycloak-demo/grool"
	mesagging "keycloak-demo/kafka/messaging"
	"keycloak-demo/model"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
)

type OrderHandler struct {
	database.IOrderDB
}

type OrderHandlerInterface interface {
	Place_order(msg *mesagging.Messaging, rdb *redis.Client, ctx context.Context) func(c *fiber.Ctx) error
	Get_Order_By_ID(c *fiber.Ctx) error
}

func NewOrderHandler(iOrderDB database.IOrderDB) OrderHandlerInterface {
	return &OrderHandler{iOrderDB}
}

func (f *OrderHandler) Place_order(msg *mesagging.Messaging, rdb *redis.Client, ctx context.Context) func(c *fiber.Ctx) error {

	grlEngine, dctx, kb := grool.GrlExecute()
	return func(c *fiber.Ctx) error {
		order := new(model.ORDER)
		err := c.BodyParser(order)
		if err != nil {
			return err
		}
		err = order.Validate()

		if err != nil {
			return err
		}
		nav, err := rdb.Get(ctx, order.Scheme_code).Result()
		if err != nil {
			log.Print(err)
		}
		order.Nav_used, err = strconv.Atoi(nav)
		if err != nil {
			log.Print(err)
		}
		order.Amount = order.Nav_used * order.Units
		dctx.Add("ORDER", order)
		err = grlEngine.Execute(dctx, kb)

		if err != nil {
			panic(err)
		}

		if !order.UintsFlg {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error": "units should be >= 10",
			})
		}

		order.Status = "Placed"
		order.Placed_at = time.Now().Unix()

		order, err = f.PlaceOrder(order, msg)
		if err != nil {
			return err
		}

		return c.JSON(order)
	}

}

func (f *OrderHandler) Get_Order_By_ID(c *fiber.Ctx) error {
	id := c.Params("order_id")
	if id == "" {
		return errors.New("order ID is required")
	}
	order, err := f.GetOrdersByUserID(id)
	if err != nil {
		return errors.New("failed to retrieve order")
	}
	return c.JSON(order)

}
