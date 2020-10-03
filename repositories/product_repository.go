package repositories

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/luisgomez29/golang-api-rest/models"
	"gorm.io/gorm"
)

type ProductRepository interface {
	All() ([]*models.Product, error)
	FindById(uint32) (*models.Product, error)
	Create(*models.Product) (*models.Product, error)
	Update(uint32, *models.Product) (*models.Product, error)
	Delete(uint32) (int64, error)
}

type productDB struct {
	conn *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productDB{db}
}

func (db *productDB) All() ([]*models.Product, error) {
	rows, err := db.conn.Model(&models.Product{}).
		Select(`products.id, products.name, products.created_at, products.updated_at, products.user_id, users.id,
			users.first_name, users.last_name, CONCAT(users.first_name, ' ', users.last_name) AS full_name, users.email,
			users.created_at, users.updated_at`).Joins("INNER JOIN users ON products.user_id = users.id").
		Limit(100).Rows()

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*models.Product
	for rows.Next() {
		var p models.Product
		var u models.User
		if err := rows.Scan(
			&p.ID, &p.Name, &p.CreatedAt, &p.UpdatedAt, &p.UserID, &u.ID, &u.FirstName, &u.LastName, &u.FullName,
			&u.Email, &u.CreatedAt, &u.UpdatedAt,
		); err != nil {
			return nil, err
		}
		p.User = u
		products = append(products, &p)
	}
	return products, nil
}

func (db *productDB) FindById(id uint32) (*models.Product, error) {
	row := db.conn.Model(&models.Product{}).
		Select(`products.id, products.name, products.created_at, products.updated_at, products.user_id, users.id,
			users.first_name, users.last_name, CONCAT(users.first_name, ' ', users.last_name) AS full_name, users.email,
			users.created_at, users.updated_at`).Joins("INNER JOIN users ON products.user_id = users.id").
		Where("products.id = ?", id).Row()

	var p models.Product
	var u models.User
	if err := row.Scan(
		&p.ID, &p.Name, &p.CreatedAt, &p.UpdatedAt, &p.UserID, &u.ID, &u.FirstName, &u.LastName, &u.FullName, &u.Email,
		&u.CreatedAt, &u.UpdatedAt,
	); err != nil {
		return nil, echo.ErrNotFound
	}
	p.User = u
	return &p, nil
}

func (db *productDB) Create(p *models.Product) (*models.Product, error) {
	err := db.conn.Create(p).Error
	if err != nil {
		return nil, echo.ErrInternalServerError
	}
	return p, nil
}

func (db *productDB) Update(id uint32, product *models.Product) (*models.Product, error) {
	p, err := db.FindById(id)
	if err != nil {
		return nil, err
	}
	p.Name = product.Name
	if err := db.conn.Save(&p).Error; err != nil {
		return nil, echo.ErrInternalServerError
	}
	return p, nil
}

func (db *productDB) Delete(id uint32) (int64, error) {
	rs := db.conn.Select("id").Take(&models.Product{}, id).Delete(&models.Product{})
	if errors.Is(rs.Error, gorm.ErrRecordNotFound) {
		return 0, echo.ErrNotFound
	}
	return rs.RowsAffected, nil
}
