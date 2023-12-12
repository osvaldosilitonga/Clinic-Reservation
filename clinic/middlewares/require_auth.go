package middlewares

import (
	"clinic/utils"
	"fmt"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func RequireAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie("Authorization")
		if err != nil {
			return utils.ErrorMessage(c, &utils.ApiUnauthorized, "Authorization cookie not found")
		}

		tokenString := cookie.Value
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("failed to verify token signature")
			}
			return []byte(os.Getenv("JWT_SECRET")), nil
		})
		if err != nil {
			return utils.ErrorMessage(c, &utils.ApiUnauthorized, err.Error())
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Set("user", claims)
			return next(c)
		}
		return utils.ErrorMessage(c, &utils.ApiUnauthorized, "Please log in to access this page")
	}
}
