package models

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/v1bh475u/LibMan_MVC/pkg/types"
)

func FetchBorrowingHistory(username string, bookname string) []types.DBorrowingHistory {
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

	var borrowingHistory []types.DBorrowingHistory
	for result.Next() {
		var bh types.BorrowingHistory
		var returned_dateBytes []byte
		var borrowed_dateBytes []byte
		err := result.Scan(&bh.ID, &bh.BookID, &bh.Title, &bh.Username, &borrowed_dateBytes, &returned_dateBytes)
		if err != nil {
			fmt.Printf("Error scanning database: %v", err)
			return nil
		}
		bh.Borrowed_date, _ = time.Parse("2006-01-02", string(borrowed_dateBytes))
		if returned_dateBytes != nil {
			bh.Returned_date, _ = time.Parse("2006-01-02", string(returned_dateBytes))
		}
		dbh := types.DBorrowingHistory{ID: bh.ID, BookID: bh.BookID, Title: bh.Title, Username: bh.Username, Borrowed_date: bh.Borrowed_date.Format("Mon Jan _2 15:04:05 2006"), Returned_date: bh.Returned_date.Format("Mon Jan _2 15:04:05 2006")}
		borrowingHistory = append(borrowingHistory, dbh)
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
	var returned_date interface{} = bh.Returned_date
	if bh.Returned_date.IsZero() {
		returned_date = nil
	}
	insertBorrowingHistory := `INSERT INTO borrowing_history (BookID, Title, Username, Borrowed_date, Returned_date) VALUES (?, ?, ?, ?, ?)`
	_, err = db.Exec(insertBorrowingHistory, bh.BookID, bh.Title, bh.Username, bh.Borrowed_date, returned_date)
	if err != nil {
		fmt.Printf("Error inserting into database: %v", err)
		return err
	}
	return nil
}

func UpdateBorrowingHistory(ID sql.NullInt64, Returned_date time.Time, username string) error {
	db, err := connection()
	if err != nil {
		fmt.Printf("Error connecting to database: %v", err)
		return err
	}
	defer db.Close()

	updateBorrowingHistory := `UPDATE borrowing_history SET Returned_date = ? WHERE BookID = ? AND Username=?`
	_, err = db.Exec(updateBorrowingHistory, Returned_date, ID, username)
	if err != nil {
		fmt.Printf("Error updating database: %v", err)
		return err
	}
	return nil
}
