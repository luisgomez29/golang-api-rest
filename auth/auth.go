package auth

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/labstack/echo/v4"
	"github.com/luisgomez29/golang-api-rest/config"
	"github.com/luisgomez29/golang-api-rest/models"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

// JWTResponse define la respuesta para JWT
type JWTResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
	*models.User `json:"user"`
}

// VerifyPassword verifica el hash de la contraseña con su equivalente en texto plano
func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

// GenerateTokens genera el token de acceso y refresh.
func GenerateTokens(user *models.User) (map[string]string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"id":  user.ID,
		"exp": time.Now().Add(time.Hour * 1).Unix(),
	})
	t, err := token.SignedString(config.SECRETKEY)
	if err != nil {
		return nil, err
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"id":      user.ID,
		"refresh": true,
		"exp":     time.Now().Add(time.Hour * 2).Unix(),
	})
	rt, err := refreshToken.SignedString(config.SECRETKEY)
	if err != nil {
		return nil, err
	}

	return map[string]string{
		"token":         t,
		"refresh_token": rt,
	}, nil
}

// RefreshToken genera el nuevo token de acceso y refresh
func RefreshToken(r *http.Request) (*JWTResponse, error) {
	claims, err := VerifyToken(r)
	if err != nil {
		return nil, err
	}
	if claims["refresh"] != nil {
		user := &models.User{ID: uint32(claims["id"].(float64))}
		tokens, err := GenerateTokens(user)
		if err != nil {
			return nil, echo.ErrInternalServerError
		}
		return &JWTResponse{
			Token:        tokens["token"],
			RefreshToken: tokens["refresh_token"],
			User:         user,
		}, nil
	}
	return nil, echo.NewHTTPError(http.StatusUnauthorized, "invalid token")
}

// VerifyToken verifica el token. Retorna el payload si el token es valido, error en caso contrario
func VerifyToken(r *http.Request) (jwt.MapClaims, error) {
	token, _ := request.ParseFromRequest(r, request.OAuth2Extractor, func(token *jwt.Token) (interface{}, error) {
		return config.SECRETKEY, nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, echo.ErrBadRequest
}

// TokenPayload obtiene el payload del token
func TokenPayload(c echo.Context) jwt.MapClaims {
	user := c.Get("user").(*jwt.Token)
	return user.Claims.(jwt.MapClaims)
}

// UserIDFromToken obtiene el ID del usuario del token.
func UserIDFromToken(c echo.Context) float64 {
	return TokenPayload(c)["id"].(float64)
}
