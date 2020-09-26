package repositories

import (
	"github.com/go-pg/pg/v10"
	"github.com/labstack/echo/v4"
	"github.com/luisgomez29/golang-api-rest/models"
	"time"
)

type (
	UserRepository interface {
		All() ([]*models.User, error)
		FindById(uint32) (*models.User, error)
		Create(*models.User) (*models.User, error)
		Update(*models.User) (*models.User, error)
		Delete(uint32) error
	}

	userDB struct {
		conn *pg.DB
	}
)

func NewUserRepository(db *pg.DB) UserRepository {
	return &userDB{db}
}

func (db *userDB) All() ([]*models.User, error) {
	var users []*models.User
	err := db.conn.Model(&users).Select()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (db *userDB) FindById(id uint32) (*models.User, error) {
	u := new(models.User)
	if err := db.conn.Model(u).Where("id = ?", id).Select(); err != nil {
		return nil, echo.ErrNotFound
	}
	return u, nil
}

func (db *userDB) Create(user *models.User) (*models.User, error) {
	_, err := db.conn.Model(user).Insert()
	if err != nil {
		return nil, echo.ErrInternalServerError
	}
	return user, nil
}

func (db *userDB) Update(u *models.User) (*models.User, error) {
	user, err := db.FindById(u.ID)
	if err != nil {
		return nil, err
	}
	user.FirstName = u.FirstName
	user.LastName = u.LastName
	user.Email = u.Email
	user.UpdatedAt = time.Now()
	_, err = db.conn.Model(user).Where("id = ?", user.ID).Update()
	return u, err
}

func (db *userDB) Delete(id uint32) error {
	user, err := db.FindById(id)
	if err != nil {
		return err
	}
	_, err = db.conn.Model(user).Where("id = ?", user.ID).Delete()
	return err
}
