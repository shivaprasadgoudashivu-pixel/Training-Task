package database

import (
	"errors"
	"keycloak-demo/model"

	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

type HOLDINGSDB struct {
	DB *gorm.DB
}

type IHoldingsDB interface {
	AddHoldings(holdings *model.HOLDINGS) error
	GetHoldings(Id uint16) (*model.HOLDINGS, error)
}

func NewHoldingsDB(db *gorm.DB) IHoldingsDB {
	return &HOLDINGSDB{db}
}

func (H *HOLDINGSDB) AddHoldings(holdings *model.HOLDINGS) error {

	err := H.DB.Create(holdings)
	if err != nil {
		log.Printf("error in adding holdings")
	}
	return errors.New("Error in adding holdings")

}

func (H *HOLDINGSDB) GetHoldings(Id uint16) (*model.HOLDINGS, error) {
	holdings := new(model.HOLDINGS)
	result := H.DB.Where("userId = ?", Id).Find(&holdings)
	if result.Error != nil {
		return nil, errors.New("unable to find orders by user ID")
	}
	return holdings, nil

}
