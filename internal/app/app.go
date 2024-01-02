package app

import (
	"GoSocial/internal/handlers"
	customMiddleware "GoSocial/internal/middleware"
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
	r.Use(customMiddleware.AuthMiddleware)

	r.GET("/:id", handlers.GetUser)
	r.PATCH("/:id/update-profile", handlers.EditUser)
	r.DELETE("/:id/delete-profile", handlers.DeleteUser)

	e.Logger.Fatal(e.Start(":3337"))
}
