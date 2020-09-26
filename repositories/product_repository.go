package repositories

import (
	"github.com/go-pg/pg/v10"
	"github.com/luisgomez29/golang-api-rest/models"
)

type (
	ProductRepository interface {
		All() ([]*models.Product, error)
	}

	productDB struct {
		conn *pg.DB
	}
)

func NewProductRepository(db *pg.DB) ProductRepository {
	return &productDB{db}
}

func (db *productDB) All() ([]*models.Product, error) {
	var u []*models.Product
	err := db.conn.Model(&u).Column("product.*").Relation("User").Select()
	if err != nil {
		panic(err)
	}
	return u, nil
}
