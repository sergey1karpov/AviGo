package repository

import (
	"GoSocial/internal/models"
	"GoSocial/internal/transport"
)

func GetUser(id int) (*models.User, error) {
	db, _ := transport.ConnectDB()
	defer db.Close()

	user := models.User{}
	err := db.QueryRow("SELECT * FROM users WHERE id = $1", id).Scan(&user.Id, &user.Username, &user.Password, &user.Email, &user.Role)

	return &user, err
}

func UpdUser(id int, data map[string][]string) (*models.User, error) {
	db, _ := transport.ConnectDB()
	defer db.Close()

	user := models.User{}
	user.Id = id
	user.Username = data["username"][0]
	user.Email = data["email"][0]

	stm := `UPDATE users SET username = $1, email = $2 WHERE id = $3;`
	_, err := db.Exec(stm, data["username"][0], data["email"][0], id)

	if err != nil {
		return nil, err
	}

	return &user, err
}

func DelUser(id int) error {
	db, _ := transport.ConnectDB()
	defer db.Close()

	stm := `DELETE FROM users WHERE id = $1`
	_, err := db.Exec(stm, id)

	if err != nil {
		return err
	}

	return err
}
