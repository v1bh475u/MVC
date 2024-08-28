package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
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
	port := os.Getenv("DB_PORT")
	fmt.Printf("%s:%s@tcp(%s:%s)/%s\n", username, password, hostname, port, dbname)
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, hostname, port, dbname)
}

func connection() (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn())
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

func waitForDB() (*sql.DB, error) {
	var db *sql.DB
	var err error
	for i := 0; i < 10; i++ { // Retry up to 10 times
		db, err = connection()
		if err == nil {
			err = db.Ping()
			if err == nil {
				return db, nil
			}
		}
		log.Println("Waiting for database to be ready...")
		time.Sleep(5 * time.Second) // Wait for 5 seconds before retrying
	}
	return nil, fmt.Errorf("database not ready: %v", err)
}

func main() {
	// time.Sleep(10 * time.Second)
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Error loading .env file. %v", err)
	}
	AMDIN_PASS := os.Getenv("ADMIN_PASS")

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(AMDIN_PASS), 14)
	if err != nil {
		fmt.Println("Error hashing password:", err)
	}
	fmt.Println(string(hashedPassword))
	db, err := waitForDB()
	if err != nil {
		log.Printf("Error connecting to database: %v", err)
		return
	}
	defer db.Close()
	insertUser := `INSERT INTO users (Username, Password, Role ) VALUES (?, ?, ?)`
	_, err = db.Exec(insertUser, "admin", string(hashedPassword), "admin")
	if err != nil {
		log.Printf("Error inserting into database: %#v", err)
	}
}
