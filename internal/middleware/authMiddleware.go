package middleware

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"strings"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Чистый токен без Bearer
		clearToken := strings.ReplaceAll(c.Request().Header.Get("Authorization"), "Bearer ", "")

		// Чек и валидация JWT токена
		token, err := jwt.Parse(clearToken, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("неожиданный метод подписи: %v", token.Header["alg"])
			}

			return []byte("very-secret-jwt-key"), nil
		})

		if err != nil {
			return echo.ErrUnauthorized
		}

		// Чек что токен действителен
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Set("userId", claims["userId"]) //Сохранить ID пользователя для дальнейшего использования
			return next(c)
		}

		return echo.ErrUnauthorized
	}
}
