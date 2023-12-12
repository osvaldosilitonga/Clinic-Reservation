package helpers

import (
	"clinic/models/dto"
	"clinic/utils"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func SignNewJWT(c echo.Context, id int, email, role string) error {
	claims := jwt.MapClaims{
		"exp":   time.Now().Add(4 * time.Hour).Unix(),
		"id":    id,
		"email": email,
		"role":  role,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return utils.ErrorMessage(c, &utils.ApiInternalServer, err.Error())
	}

	cookie := new(http.Cookie)
	cookie.Name = "Authorization"
	cookie.HttpOnly = true
	cookie.Path = "/"
	cookie.Value = tokenString
	cookie.SameSite = http.SameSiteLaxMode
	cookie.Expires = time.Now().Add(2 * time.Hour)
	c.SetCookie(cookie)

	return nil
}

func GetClaims(c echo.Context) (*dto.Claims, error) {
	claimsTmp := c.Get("user")
	if claimsTmp == nil {
		return nil, utils.ErrorMessage(c, &utils.ApiUnauthorized, "Failed to fetch user claims from JWT")
	}

	claims := claimsTmp.(jwt.MapClaims)
	return &dto.Claims{
		ID:    int(claims["id"].(float64)),
		Email: claims["email"].(string),
		Role:  claims["role"].(string),
	}, nil
}
