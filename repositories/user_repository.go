package repositories

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/luisgomez29/golang-api-rest/models"
	"github.com/luisgomez29/golang-api-rest/utils"
	"gorm.io/gorm"
)

type UserRepository interface {
	All() ([]*models.User, error)
	FindById(uint32) (*models.User, error)
	Create(*models.User) (*models.User, error)
	Update(uint32, *models.User) (*models.User, error)
	Delete(uint32) (int64, error)
}

type database struct {
	conn *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &database{db}
}

func (db *database) All() ([]*models.User, error) {
	var users []*models.User
	db.conn.Select(utils.Fields(&models.User{})).Find(&users)
	return users, nil
}

func (db *database) FindById(id uint32) (*models.User, error) {
	u := new(models.User)
	fields := utils.Fields(&models.User{})
	if err := db.conn.Select(fields).Take(u, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, echo.ErrNotFound
	}
	return u, nil
}

func (db *database) Create(user *models.User) (*models.User, error) {
	err := db.conn.Create(user).Error
	if err != nil {
		return nil, echo.ErrInternalServerError
	}
	user.FirstName = user.FirstName + " " + user.LastName
	return user, nil
}

func (db *database) Update(id uint32, user *models.User) (*models.User, error) {
	u, err := db.FindById(id)
	if err != nil {
		return nil, err
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
	rs := db.conn.Debug().Select("id").Take(&models.User{}, id).Delete(&models.User{})
	if errors.Is(rs.Error, gorm.ErrRecordNotFound) {
		return 0, echo.ErrNotFound
	}
	return rs.RowsAffected, nil
}
