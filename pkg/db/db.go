package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() error {
	DB_DSN, er := os.LookupEnv("DB_DSN")
	if !er {
		log.Fatal("Missing env variable DB_DSN")
	}

	conn, err := sql.Open("postgres", DB_DSN)
	if err != nil {
		panic(err)
	}

	return conn.Ping()
}
