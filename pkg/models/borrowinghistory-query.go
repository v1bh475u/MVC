package models

import (
	"fmt"
	"strings"

	"github.com/v1bh475u/LibMan_MVC/pkg/types"
)

func FetchBorrowingHistory(username string, bookname string) []types.BorrowingHistory {
	db, err := connection()
	if err != nil {
		fmt.Printf("Error connecting to database: %v", err)
		return nil
	}
	defer db.Close()

	query := `SELECT * FROM borrowing_history`
	conditions := []string{}
	params := []interface{}{}

	if username != "" {
		conditions = append(conditions, `Username = ?`)
		params = append(params, username)
	}
	if bookname != "" {
		conditions = append(conditions, `Title = ?`)
		params = append(params, bookname)
	}
	if len(conditions) > 0 {
		query += ` WHERE ` + strings.Join(conditions, ` AND `)
	}
	result, err := db.Query(query, params...)
	if err != nil {
		fmt.Printf("Error querying database: %v", err)
		return nil
	}
	var borrowingHistory []types.BorrowingHistory
	for result.Next() {
		var bh types.BorrowingHistory
		err := result.Scan(&bh.ID, &bh.BookID, &bh.Title, &bh.Username, &bh.Borrowed_date, &bh.Returned_date)
		if err != nil {
			fmt.Printf("Error scanning database: %v", err)
			return nil
		}
		borrowingHistory = append(borrowingHistory, bh)
	}
	return borrowingHistory
}

func InsertBorrowingHistory(bh types.BorrowingHistory) error {
	db, err := connection()
	if err != nil {
		fmt.Printf("Error connecting to database: %v", err)
		return err
	}
	defer db.Close()

	insertBorrowingHistory := `INSERT INTO borrowing_history (BookID, Title, Username, Borrowed_date, Returned_date) VALUES (?, ?, ?, ?, ?)`
	_, err = db.Exec(insertBorrowingHistory, bh.BookID, bh.Title, bh.Username, bh.Borrowed_date, bh.Returned_date)
	if err != nil {
		fmt.Printf("Error inserting into database: %v", err)
		return err
	}
	return nil
}

func UpdateBorrowingHistory(ID int, Returned_date string) error {
	db, err := connection()
	if err != nil {
		fmt.Printf("Error connecting to database: %v", err)
		return err
	}
	defer db.Close()

	updateBorrowingHistory := `UPDATE borrowing_history SET Returned_date = ? WHERE ID = ?`
	_, err = db.Exec(updateBorrowingHistory, Returned_date, ID)
	if err != nil {
		fmt.Printf("Error updating database: %v", err)
		return err
	}
	return nil
}
