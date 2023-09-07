package db

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/alessandra1408/crud-golang/configs"
	_ "github.com/lib/pq"
)

func NewPostgresConnection() (*sql.DB, error) {
	conf := configs.GetDB()

	postgresURL := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", conf.Host, conf.Port, conf.User, conf.Pass, conf.Database)

	conn, cErr := sql.Open("postgres", postgresURL)
	if cErr != nil {
		return nil, errors.New("error to connect with database")
	}

	pErr := conn.Ping()
	if pErr != nil {
		return nil, errors.New("error on Ping database")
	}

	return conn, nil
}
