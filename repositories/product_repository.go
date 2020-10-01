package repositories

import (
	"github.com/luisgomez29/golang-api-rest/models"
	"gorm.io/gorm"
)

type ProductRepository interface {
	All() ([]*models.Product, error)
}

type productDB struct {
	conn *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productDB{db}
}

func (db *productDB) All() ([]*models.Product, error) {
	rows, err := db.conn.Debug().Model(&models.Product{}).
		Select(`products.id, products.name, products.created_at, products.updated_at, products.user_id, users.id,
			users.first_name, users.last_name, CONCAT(users.first_name, ' ', users.last_name) AS full_name, users.email,
			users.created_at, users.updated_at`).Joins("INNER JOIN users ON products.user_id = users.id").Rows()

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var p []*models.Product
	for rows.Next() {
		var product models.Product
		var user models.User
		if err := rows.Scan(
			&product.ID, &product.Name, &product.CreatedAt, &product.UpdatedAt, &product.UserID, &user.ID,
			&user.FirstName, &user.LastName, &user.FullName, &user.Email, &user.CreatedAt, &user.UpdatedAt,
		); err != nil {
			return nil, err
		}
		product.User = user
		p = append(p, &product)
	}
	return p, nil
}
