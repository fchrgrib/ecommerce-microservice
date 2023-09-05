package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
)

type Database struct {
	Conn *sql.DB
}

func (db *Database) Connect() {
	postgresConn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	_db, err := sql.Open("postgres", postgresConn)

	if err != nil {
		log.Fatal(err)
	}

	if _, err := _db.Exec("CREATE TABLE IF NOT EXISTS payment(id PRIMARY KEY VARCHAR(225), customer_id VARCHAR(225), customer_name VARCHAR(225), product_id VARCHAR(225), product_name VARCHAR(225), price VARCHAR(225), date_order TIMESTAMP, service_shipment_name VARCHAR(225), date_shipment TIMESTAMP, shipment_methode VARCHAR(225))"); err != nil {
		log.Fatal(err)
	}

	db.Conn = _db
}
