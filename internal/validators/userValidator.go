package validators

import (
	"GoSocial/internal/models"
	"GoSocial/internal/transport"
	"github.com/go-playground/validator/v10"
)

func UserRegValidation(username, email, password string) error {
	validate := validator.New()

	inputData := &models.User{
		Username: username,
		Email:    email,
		Password: password,
	}

	err := validate.Struct(inputData)

	return err
}

func CheckUsernameAndEmail(username, email string) ([]string, error) {
	db, err := transport.ConnectDB()
	defer db.Close()

	var validationErrors []string

	row := db.QueryRow("SELECT EXISTS (SELECT 1 FROM users WHERE username = $1)", username)

	var exists bool

	err = row.Scan(&exists)
	if err != nil {
		return nil, err
	}

	if exists {
		validationErrors = append(validationErrors, "Username already exists")
	}

	row = db.QueryRow("SELECT EXISTS (SELECT 1 FROM users WHERE email = $1)", email)

	err = row.Scan(&exists)
	if err != nil {
		return nil, err
	}

	if exists {
		validationErrors = append(validationErrors, "Email already exists")
	}

	return validationErrors, nil
}
