package app

import (
	"GoSocial/internal/handlers"
	"GoSocial/internal/models"
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Start() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/login", handlers.Login)
	e.POST("/registration", handlers.Registration)

	r := e.Group("/user")

	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(models.JwtCustomClaims)
		},
		SigningKey: []byte("very-secret-jwt-key"),
	}
	r.Use(echojwt.WithConfig(config))

	r.GET("/:id", handlers.GetUser)
	r.PATCH("/:id/update-profile", handlers.EditUser)

	e.Logger.Fatal(e.Start(":3336"))
}
