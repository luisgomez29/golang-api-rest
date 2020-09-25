package repositories

import (
	"github.com/luisgomez29/golang-api-rest/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type (
	ProductRepository interface {
		All() ([]*models.User, error)
	}

	productDB struct {
		conn *gorm.DB
	}
)

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productDB{db}
}

func (db *productDB) All() ([]*models.User, error) {
	var u []*models.User
	result := db.conn.Debug().Preload(clause.Associations).Where("id >= 1").Order("id desc").Find(&u)
	//result := db.conn.Debug().Find(&u)
	if result.Error != nil {
		return nil, result.Error
	}
	return u, nil
}
