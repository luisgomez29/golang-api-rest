package repositories

import (
	"github.com/luisgomez29/golang-api-rest/models"
	"github.com/luisgomez29/golang-api-rest/utils"
	"gorm.io/gorm"
)

type (
	ProductRepository interface {
		All() ([]*models.Product, error)
	}

	productDB struct {
		conn *gorm.DB
	}
)

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productDB{db}
}

func (db *productDB) All() ([]*models.Product, error) {
	var p []*models.Product
	fields := utils.Fields(&models.Product{})
	err := db.conn.Debug().Preload("User").Select(fields).Find(&p).Error
	if err != nil {
		return nil, err
	}
	return p, nil
}
