package main

import (
	"GoSocial/internal/handlers"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.POST("/login", handlers.Login)
	e.POST("/registration", handlers.Registration)
	e.Logger.Fatal(e.Start(":3336"))
}
