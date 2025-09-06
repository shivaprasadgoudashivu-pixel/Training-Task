package model

import (
	"encoding/json"
	"errors"
)

type ORDER struct {
	Id           int    `json:"id" gorm:"primaryKey;autoIncrement"`
	UserId       int    `json:"userId"`
	Scheme       string `json:sym`
	Price        int    `json:price`
	Units        int    `json:units`
	Status       string `json:status`
	Scheme_code  string `json:"schemeCode"`
	Nav_used     uint16 `json:"nav_used"`
	Placed_at    int64  `josn:"placed_at" gorm:"timestamptz"`
	Confirmed_at int64  `json:"confirm_at" gorm:"timestamptz"`
	Contact_Url  string `json: "contact_url"`
}

func (oe *ORDER) ToBytes() []byte {
	bytes, _ := json.Marshal(oe)
	return bytes
}

func (c *ORDER) Validate() error {

	if c.UserId == 0 {
		return errors.New("UserId can not be null")
	}
	if c.Units == 0 {
		return errors.New("Units can not be null")
	}
	if c.Scheme == "" {
		return errors.New("Scheme can not be null")
	}

	if c.Scheme_code == "" {
		return errors.New("Scheme_code can not be null")
	}
	return nil
}
