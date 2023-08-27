package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	DB *sql.DB
}

func (db *Database) Connect() {
	_db, err := sql.Open("sqlite3", "./sqlite.db")

	if err != nil {
		panic(err)
	}

	if _, err := _db.Exec("CREATE TABLE IF NOT EXISTS product(id VARCHAR(225) PRIMARY KEY, product_name VARCHAR(225), category VARCHAR(225), product_type VARCHAR(225), created_at TIME, updated_at TIME)"); err != nil {
		panic(err)
	}

	db.DB = _db
}
