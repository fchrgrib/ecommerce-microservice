package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

type Database struct {
	DB *sql.DB
}

func (db *Database) Connect() {
	host, pass, userName, dbName, port := os.Getenv("DB_HOST"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_USER"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT")

	_db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s", host, port, userName, pass, dbName))

	if err != nil {
		log.Fatal(err)
		return
	}

	if _, err := _db.Exec("CREATE TABLE IF NOT EXISTS users(id VARCHAR(225) PRIMARY KEY, user_id VARCHAR(225) UNIQUE, user_name VARCHAR(225), phone_number VARCHAR(225), email VARCHAR(225) UNIQUE, address VARCHAR(225), born VARCHAR(225), created_at TIME, updated_at TIME)"); err != nil {
		log.Fatal(err)
		return
	}

	if _, err := _db.Exec("CREATE TABLE IF NOT EXISTS history_order_product(id VARCHAR(225) PRIMARY KEY, user_id VARCHAR(225), product_id VARCHAR(225), order_date TIME, FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE)"); err != nil {
		log.Fatal(err)
		return
	}

	db.DB = _db
}
