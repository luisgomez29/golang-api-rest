package repositories

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/luisgomez29/golang-api-rest/models"
	"gorm.io/gorm"
)

type (
	UserRepository interface {
		All() ([]*models.User, error)
		FindById(uint32) (*models.User, error)
		Create(*models.User) (*models.User, error)
		Update(uint32, *models.User) (*models.User, error)
		Delete(uint32) (int64, error)
	}

	database struct {
		conn *gorm.DB
	}
)

func NewUserRepository(db *gorm.DB) UserRepository {
	return &database{db}
}

func (db *database) All() ([]*models.User, error) {
	var users []*models.User
	err := db.conn.Find(&users).Error
	return users, err
}

func (db *database) FindById(id uint32) (*models.User, error) {
	u := new(models.User)
	if err := db.conn.Take(u, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, echo.ErrNotFound
	}
	return u, nil
}

func (db *database) Create(user *models.User) (*models.User, error) {
	err := db.conn.Create(user).Error
	if err != nil {
		return nil, echo.ErrInternalServerError
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

func (db *database) Delete(id uint32) (int64, error) {
	rs := db.conn.Take(&models.User{}, id).Delete(&models.User{}, id)
	if errors.Is(rs.Error, gorm.ErrRecordNotFound) {
		return 0, echo.ErrNotFound
	}
	return rs.RowsAffected, nil
}
