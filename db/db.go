package db

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

var DB *bun.DB

func GetDbDSN() (string, error) {
	var DB_HOST string
	var DB_PORT int
	var DB_USER string
	var DB_PASSWORD string
	var DB_NAME string

	DB_HOST = os.Getenv("DB_HOST")
	if DB_HOST == "" {
		DB_HOST = "localhost"
	}
	DB_PORT, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		fmt.Println("Error parsing DB_PORT to integer")
		return "", err
	}
	DB_USER = os.Getenv("DB_USER")
	if DB_USER == "" {
		return "", errors.New("no DB user provided in .env")
	}
	DB_PASSWORD = os.Getenv("DB_PASSWORD")
	if DB_PASSWORD == "" {
		return "", errors.New("no DB user provided in .env")
	}
	DB_NAME = os.Getenv("DB_NAME")
	if DB_NAME == "" {
		return "", errors.New("no DB user provided in .env")
	}
	var DSN string = fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME)
	return DSN, nil
}

func ConnectDb() error {
	var DSN string
	DSN, err := GetDbDSN()
	if err != nil {
		return err
	}
	var sqldb *sql.DB = sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(DSN)))
	DB = bun.NewDB(sqldb, pgdialect.New())
	return nil
}
