package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

type Database struct {
	DB *sql.DB
}

func (db *Database) Connect() {
	mysqlConnString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	_db, err := sql.Open("mysql", mysqlConnString)
	if err != nil {
		log.Fatal(err)
		return
	}

	if _, err := _db.Exec("CREATE TABLE IF NOT EXISTS users(id VARCHAR(225) PRIMARY KEY, user_id VARCHAR(225) UNIQUE, user_name VARCHAR(225), phone_number VARCHAR(225), email VARCHAR(225) UNIQUE, address VARCHAR(225), born VARCHAR(225), created_at TIME, updated_at TIME)"); err != nil {
		log.Fatal(err)
		return
	}

	db.DB = _db
}
