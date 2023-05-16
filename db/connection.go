package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

import (
	"os"
)

var (
	host     = os.Getenv("host")
	password = os.Getenv("password")
	port     = 5432
	user     = "user"
	database = "udrive"
)

func GetConnection() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, database)
	_, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Printf("error connecting to database %v\n", err)
		panic(err)
	}

	dbConnection, err := GetConnection()
	if err != nil {
		fmt.Printf("error connecting to database %v\n", err)
		panic(err)
	}

	err = dbConnection.Ping()
	if err != nil {
		return nil, err
	}
	if err != nil {
		fmt.Printf("error pinging database %v\n", err)
		panic(err)
	}

	return dbConnection, nil
}
