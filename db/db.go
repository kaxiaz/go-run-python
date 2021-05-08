package db

import (
	"database/sql"
	"fmt"
	"log"

	"go-run-python/config"

	_ "github.com/lib/pq"
)

func OpenConnection() *sql.DB {
	config := config.GetConfig()
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.GetString("db.host"), config.GetInt("db.port"), config.GetString("db.user"),
		config.GetString("db.password"), config.GetString("db.dbname"))

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func CloseConnection(db *sql.DB) {
	db.Close()
}
