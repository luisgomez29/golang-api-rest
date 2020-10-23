package controllers

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/luisgomez29/golang-api-rest/auth"
	"github.com/luisgomez29/golang-api-rest/models"
	"github.com/luisgomez29/golang-api-rest/utils"
	"gorm.io/gorm"
	"net/http"
)

type (
	AuthController interface {
		Login(echo.Context) error
		VerifyToken(echo.Context) error
		RefreshToken(echo.Context) error
	}

	loginDB struct {
		conn *gorm.DB
	}

	userLogin struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
)

func NewAuthController(db *gorm.DB) AuthController {
	return &loginDB{db}
}

func (db *loginDB) Login(c echo.Context) error {
	u := new(userLogin)
	if err := c.Bind(u); err != nil {
		return err
	}
	user := new(models.User)
	fields := utils.Fields(user)
	if err := db.conn.Select(fields).Where("email = ?", u.Email).Take(user).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return echo.NewHTTPError(http.StatusUnauthorized, "invalid email")
	}
	if err := auth.VerifyPassword(user.Password, u.Password); err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "invalid password")
	}
	user.Password = ""
	tokens, err := auth.GenerateTokens(user)
	if err != nil {
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, auth.JWTResponse{
		Token:        tokens["token"],
		RefreshToken: tokens["refresh_token"],
		User:         user,
	})
}

func (*loginDB) VerifyToken(c echo.Context) error {
	payload, err := auth.VerifyToken(c.Request())
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, payload)
}

func (*loginDB) RefreshToken(c echo.Context) error {
	res, err := auth.RefreshToken(c.Request())
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, res)
}
