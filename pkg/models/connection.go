package models

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"

	_ "github.com/go-sql-driver/mysql"
)

func dsn() string {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Error loading .env file. %v", err)
	}
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	dbname := os.Getenv("DB_NAME")
	hostname := os.Getenv("DB_HOST")
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbname)
}

func connection() (*sql.DB, error) {
	fmt.Printf("DSN: %s\n", dsn())
	db, err := sql.Open("mysql", dsn())
	fmt.Printf("DSN: %s\n", dsn())
	if err != nil {
		log.Printf("Error opening database: %v", err)
		return nil, err
	}
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(20)
	db.SetConnMaxIdleTime(time.Minute * 10)

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	err = db.PingContext(ctx)
	if err != nil {
		log.Printf("Error pinging database: %v", err)
		return nil, err
	}

	log.Println("Connected to database")

	return db, err
}
