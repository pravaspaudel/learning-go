package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/pravaspaudel/09_crud/config"
)

var Db *sql.DB

func ConnectDb() {

	db, err := sql.Open("postgres", config.App.DB_URL)

	if err != nil {
		log.Fatal("failed to start a connection")
		return
	}

	Db = db

	if err := Db.Ping(); err != nil {
		log.Fatal("error connecting to the database", err)
	}

	log.Println("database connected successfully")
}
