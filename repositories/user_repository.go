package repositories

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/luisgomez29/golang-api-rest/models"
	"gorm.io/gorm"
)

type (
	UserRepository interface {
		Create(*models.User) (*models.User, error)
		Update(uint32, *models.User) (*models.User, error)
		//Delete(uint32) (int, error)
		All() ([]*models.User, error)
		//FindById(uint32) (*models.User, error)
	}

	database struct {
		conn *gorm.DB
	}
)

func NewUserRepository(db *gorm.DB) *database {
	return &database{db}
}

func (db *database) All() ([]*models.User, error) {
	var users []*models.User
	err := db.conn.Find(&users).Error
	return users, err
}

func (db *database) Create(user *models.User) (*models.User, error) {
	err := db.conn.Create(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (db *database) Update(id uint32, user *models.User) (*models.User, error) {
	u := new(models.User)
	if err := db.conn.Take(u, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, echo.ErrNotFound
	}

	u.FirstName = user.FirstName
	u.LastName = user.LastName
	u.Email = user.Email
	u.Password = user.Password

	if err := db.conn.Save(u).Error; err != nil {
		return nil, echo.ErrInternalServerError
	}
	return u, nil
}
