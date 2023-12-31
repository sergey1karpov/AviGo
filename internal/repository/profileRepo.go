package repository

import (
	"GoSocial/internal/models"
	"GoSocial/internal/transport"
)

func GetUser(id int) (*models.User, error) {
	db, _ := transport.ConnectDB()
	defer db.Close()

	user := models.User{}
	err := db.QueryRow("SELECT * FROM users WHERE id = $1", id).Scan(&user.Id, &user.Username, &user.Email, &user.Password, &user.Role)

	return &user, err
}
