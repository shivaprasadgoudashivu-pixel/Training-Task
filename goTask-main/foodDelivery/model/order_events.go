<<<<<<< HEAD
package model

import "encoding/json"

type OrderEvents struct {
	Id        uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Event     string `json:"event"`
	Order_Id  int    `json:"orderid"`
	EventTime string `json:"event_time" gorm:"timestamptz"`
	COMMONMODEL
}

func (oe *OrderEvents) ToBytes() []byte {
	bytes, _ := json.Marshal(oe)
	return bytes
}
=======
package model

import "encoding/json"

type OrderEvents struct {
	Id        uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Event     string `json:"event"`
	Order_Id  int    `json:"orderid"`
	EventTime string `json:"event_time" gorm:"timestamptz"`
	COMMONMODEL
}

func (oe *OrderEvents) ToBytes() []byte {
	bytes, _ := json.Marshal(oe)
	return bytes
}
>>>>>>> 6b7cfda (changes on MFtask)
