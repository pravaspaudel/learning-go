package config

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func ConfigDB() *sql.DB {
	var err error

	Db, err := sql.Open("postgres", AppConfig.DB_CONNECTION_URL)

	if err != nil {
		log.Fatal("failed to connect to database", err)
	}

	if err := Db.Ping(); err != nil {
		log.Fatalln("failed to ping the database", err)
	}
	return Db
}
