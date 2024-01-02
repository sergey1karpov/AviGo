package handlers

import (
	"GoSocial/internal/repository"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"sync"
)

var lock sync.Mutex

// GetUser Get user profile from database/*
func GetUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	tokenId := c.Get("userId").(float64)
	intTokenId := int(tokenId)

	if id != intTokenId {
		return echo.ErrForbidden
	}

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

	tokenId := c.Get("userId").(float64)
	intTokenId := int(tokenId)

	if id != intTokenId {
		return echo.ErrForbidden
	}

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

func DeleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	tokenId := c.Get("userId").(float64)
	intTokenId := int(tokenId)

	if id != intTokenId {
		return echo.ErrForbidden
	}

	err := repository.DelUser(id)

	fmt.Println(err)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": fmt.Sprintf("User with id %d deleted", id),
	})
}
