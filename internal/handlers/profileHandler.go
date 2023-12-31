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
	lock.Lock()
	defer lock.Unlock()

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

func EditUser() {

}

func DeleteUser() {

}
