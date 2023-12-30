package handlers

import (
	"GoSocial/internal/models"
	"GoSocial/internal/repository"
	"GoSocial/internal/validators"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"net/http"
	"time"
)

type jwtCustomClaims struct {
	Username string `json:"username"`
	Admin    bool   `json:"user"`
	jwt.RegisteredClaims
}

func Login(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	user, _ := repository.FindUser(email, password)

	if user == nil {
		return echo.ErrNotFound
	}

	claims := &jwtCustomClaims{
		"Alexander Pistoletov",
		false,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
		"login": true,
	})
}

func Registration(c echo.Context) error {
	username := c.FormValue("username")
	email := c.FormValue("email")
	password, _ := models.EncryptPassword(c.FormValue("password"))

	err := validators.UserRegValidation(username, email, password)

	if err != nil {
		var validationErrors []string
		for _, err := range err.(validator.ValidationErrors) {
			validationErrors = append(validationErrors, err.Error())
		}
		return c.JSON(http.StatusBadRequest, validationErrors)
	}

	checkUsernameAndEmail, _ := validators.CheckUsernameAndEmail(username, email)
	if checkUsernameAndEmail != nil {
		var duplicateErrors []string
		for _, v := range checkUsernameAndEmail {
			duplicateErrors = append(duplicateErrors, v)
		}
		return c.JSON(http.StatusBadRequest, duplicateErrors)
	}

	user, err := repository.Reg(username, email, password)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, user)
}
