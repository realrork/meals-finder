package server

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	_ "github.com/joho/godotenv/autoload"
)

var dbConnInstance *pgx.Conn

func NewConnection() *pgx.Conn {
	if dbConnInstance != nil {
		return dbConnInstance
	}
	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_DATABASE"),
	)
	conn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		log.Fatal(err)
	}
	dbConnInstance = conn

	return dbConnInstance
}

var dbConnInstanceTest *pgx.Conn

func NewConnectionTest() *pgx.Conn {
	if dbConnInstance != nil {
		return dbConnInstance
	}

	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		os.Getenv("TEST_DB_USERNAME"),
		os.Getenv("TEST_DB_PASSWORD"),
		os.Getenv("TEST_DB_HOST"),
		os.Getenv("TEST_DB_PORT"),
		os.Getenv("TEST_DB_DATABASE"),
	)

	conn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		log.Fatal(err)
	}
	dbConnInstance = conn

	return dbConnInstance
}
