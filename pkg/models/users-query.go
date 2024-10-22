package models

import (
	"log"

	"github.com/v1bh475u/LibMan_MVC/pkg/types"
)

func FetchUser(username string) (types.User, error) {
	log.Println("Fetching user")
	db, err := connection()
	if err != nil {
		log.Printf("Error connecting to database: %v", err)
		return types.User{}, err
	}
	defer db.Close()

	getUser := `SELECT * FROM users WHERE username = ?`
	result, err := db.Query(getUser, username)
	if err != nil {
		log.Printf("Error querying database: %v", err)
		return types.User{}, err
	}
	var user types.User
	for result.Next() {
		err := result.Scan(&user.ID, &user.Username, &user.Password, &user.Role)
		if err != nil {
			log.Printf("Error scanning database: %v", err)
			return types.User{}, err
		}
	}
	return user, nil
}

func InsertUser(user types.User) error {
	db, err := connection()
	if err != nil {
		log.Printf("Error connecting to database: %v", err)
		return err
	}
	defer db.Close()

	insertUser := `INSERT INTO users (Username, Password, Role ) VALUES (?, ?, ?)`
	_, err = db.Exec(insertUser, user.Username, user.Password, user.Role)
	if err != nil {
		log.Printf("Error inserting into database: %#v", err)
		return err
	}
	return nil
}

func update_user(username, role string) error {
	db, err := connection()
	if err != nil {
		log.Printf("Error connecting to database: %v", err)
		return err
	}
	defer db.Close()

	updateUser := `UPDATE users SET Role = ? WHERE Username = ?`
	_, err = db.Exec(updateUser, role, username)
	if err != nil {
		log.Printf("Error updating database: %v", err)
		return err
	}
	return nil
}
