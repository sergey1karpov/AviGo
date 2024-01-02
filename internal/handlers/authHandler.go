package handlers

import (
	"GoSocial/internal/models"
	"GoSocial/internal/repository"
	"GoSocial/internal/validators"
	"GoSocial/pkg/mail/mailhog"
	"GoSocial/pkg/mail/templates"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"net/http"
	"time"
)

func Login(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	user, _ := repository.FindUser(email, password)

	fmt.Println(user)

	if user == nil {
		return echo.ErrNotFound
	}

	claims := &models.JwtCustomClaims{
		UserId:   user.Id,
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte("very-secret-jwt-key"))
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

	template := templates.Welcome()
	defer mailhog.SendMail(email, template)

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
