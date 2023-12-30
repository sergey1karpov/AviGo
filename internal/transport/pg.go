package transport

import (
	c "GoSocial/internal/config"
	"database/sql"
	"fmt"
	"log"
)

var config = c.Database{
	DB:        "postgres",
	DBUser:    "postgres",
	DBPass:    "secret",
	DBPort:    5432,
	DBHost:    "localhost",
	DBSslmode: "disable",
}

func ConnectDB() (*sql.DB, error) {
	conf := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		config.DBHost, config.DBPort, config.DBUser, config.DBPass, config.DB, config.DBSslmode)

	db, err := sql.Open("postgres", conf)

	if err != nil {
		log.Fatalln(err)
	}

	return db, err
}
