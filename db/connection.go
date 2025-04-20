package db

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func ConnectDB() (*sql.DB, error) {
	godotenv.Load()
	var dbConnectionString = os.Getenv("DBCONNECTION")

	db, err := sql.Open("postgres", dbConnectionString)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to " + dbConnectionString)

	return db, nil
}
