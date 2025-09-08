package database

import (
	"errors"
	"fmt"
	mesagging "keycloak-demo/kafka/messaging"
	"keycloak-demo/minio"
	"keycloak-demo/model"
	"math/rand"
	"time"

	"gorm.io/gorm"
)

type OrderDB struct {
	DB *gorm.DB
}

type IOrderDB interface {
	PlaceOrder(order *model.ORDER, msg *mesagging.Messaging) (*model.ORDER, error)
	GetOrdersByUserID(userID string) ([]model.ORDER, error)
	UpdateOrderEvent(OrderEve *model.ORDER, msg *mesagging.Messaging) error
}

func NewOrderDB(db *gorm.DB) IOrderDB {
	return &OrderDB{db}
}

func (f *OrderDB) PlaceOrder(order *model.ORDER, msg *mesagging.Messaging) (*model.ORDER, error) {
	order.Status = "Placed"
	result := f.DB.Create(order)
	if result.Error != nil {
		return nil, errors.New("unable to create order: ")
	}
	msg.ChMessaging <- order.ToBytes()
	go f.UpdateOrderEvent(order, msg)

	return order, nil
}

func (f *OrderDB) GetOrdersByUserID(userID string) ([]model.ORDER, error) {
	var orders []model.ORDER
	result := f.DB.Where("user_id = ?", userID).Find(&orders)
	if result.Error != nil {
		return nil, errors.New("unable to find orders by user ID")
	}
	return orders, nil
}

func (f *OrderDB) UpdateOrderEvent(OrderEve *model.ORDER, msg *mesagging.Messaging) error {
	var status string
	k := rand.Intn(10)
	time.Sleep(time.Second * 5)
	if k%2 == 0 {
		status = "confirmed"
	} else {
		status = "cancelled"
	}
	url := minio.UploadPDF(OrderEve)
	OrderEve.Contact_Url = url

	fmt.Printf("url return is :%S", OrderEve.Contact_Url)

	result := f.DB.Model(&model.ORDER{}).
		Where("id = ?", OrderEve.Id).
		Updates(map[string]interface{}{
			"status":       status,
			"confirmed_at": time.Now().Unix(),
			"contact_url":  OrderEve.Contact_Url,
		})

	if result.Error != nil {
		return errors.New("unable to update order event: " + result.Error.Error())
	}
	OrderEve.Status = status
	OrderEve.Confirmed_at = time.Now().Unix()
	msg.ChMessaging <- OrderEve.ToBytes()

	return nil
}
