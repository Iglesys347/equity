package db

import (
	"database/sql"
	"fmt"
	"github.com/Iglesys347/equity/logger"
	_ "github.com/lib/pq"
)

const DB_USERNAME string = "postgres"
const DB_PASSWORD string = "mysecretpassword"
const DB_NAME string = "equity"
const DB_HOST string = "localhost"
const DB_PORT uint16 = 5432

var db *sql.DB

var l = logger.Get()

func Connect() (*sql.DB, error) {
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d sslmode=disable", DB_USERNAME, DB_PASSWORD, DB_NAME, DB_HOST, DB_PORT)
	database, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	db = database
	return db, nil
}
