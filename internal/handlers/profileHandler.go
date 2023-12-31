package handlers

import (
	"GoSocial/internal/repository"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"sync"
)

var lock sync.Mutex

// GetUser Get user profile from database/*
func GetUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	user, err := repository.GetUser(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "User not found",
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"user": user,
	})
}

func EditUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	user, err := repository.GetUser(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "User not found",
		})
	}

	data, _ := c.FormParams()

	user, err = repository.UpdUser(id, data)

	return c.JSON(http.StatusOK, echo.Map{
		"message": "User updated",
		"user":    user,
	})
}

func DeleteUser() {

}
