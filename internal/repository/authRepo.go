package repository

import (
	"GoSocial/internal/models"
	"GoSocial/internal/transport"
)

func Reg(username, email, password string) (*models.User, error) {
	db, _ := transport.ConnectDB()
	defer db.Close()

	user := models.User{Username: username, Email: email, Password: password}
	query := "INSERT INTO users (username, email, password) VALUES($1, $2, $3)"
	_, err := db.Exec(query, user.Username, user.Email, user.Password)

	return &user, err
}

func FindUser(email, password string) (*models.User, error) {
	db, _ := transport.ConnectDB()
	defer db.Close()

	u := models.User{}
	err := db.QueryRow("SELECT username, email, password FROM users WHERE email = $1", email).Scan(&u.Username, &u.Email, &u.Password)

	hashPass := &u.Password

	err = models.ComparePasswords([]byte(*hashPass), []byte(password))

	if err != nil {
		return nil, err
	}

	return &u, err
}
