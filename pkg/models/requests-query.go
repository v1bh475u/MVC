package models

import (
	"fmt"
	"strings"
	"time"

	"github.com/v1bh475u/LibMan_MVC/pkg/types"
)

func FetchRequests(username, request, title, status string, ID int, User bool) []types.Request {
	db, err := connection()
	if err != nil {
		fmt.Printf("Error connecting to database: %v", err)
		return nil
	}
	defer db.Close()

	query := `SELECT * FROM requests`
	conditions := []string{}
	params := []interface{}{}
	if ID != 0 {
		conditions = append(conditions, `ID = ?`)
		params = append(params, ID)
	}
	if username != "" {
		conditions = append(conditions, `Username = ?`)
		params = append(params, username)
	}
	if request != "" {
		conditions = append(conditions, `Request = ?`)
		params = append(params, request)
	}
	if title != "" {
		conditions = append(conditions, `Title = ?`)
		params = append(params, title)
	}
	if status != "" {
		conditions = append(conditions, `Status = ?`)
		params = append(params, status)
	}
	if User {
		conditions = append(conditions, `User_status = ?`)
		params = append(params, "unseen")
	}
	if conditions != nil {
		query += ` WHERE ` + strings.Join(conditions, ` AND `)
	}
	result, err := db.Query(query, params...)
	if err != nil {
		fmt.Printf("Error querying database: %v", err)
		return nil
	}
	var requests []types.Request
	for result.Next() {
		var req types.Request
		err := result.Scan(&req.ID, &req.BookID, &req.Title, &req.Request, &req.Status, &req.User_status, &req.Date)
		if err != nil {
			fmt.Printf("Error scanning database: %v", err)
			return nil
		}
		requests = append(requests, req)
	}
	return requests
}

func InsertRequest(req types.Request) error {
	db, err := connection()
	if err != nil {
		fmt.Printf("Error connecting to database: %v", err)
		return err
	}
	defer db.Close()

	insertRequest := `INSERT INTO requests (BookID, Title, Request, Status, User_status, Date) VALUES (?, ?, ?, ?, ?, ?)`
	_, err = db.Exec(insertRequest, req.BookID, req.Title, req.Request, req.Status, req.User_status, req.Date)
	if err != nil {
		fmt.Printf("Error inserting into database: %v", err)
		return err
	}
	return nil
}

func UpdateRequest(Status, User_status string, ID int) error {
	db, err := connection()
	if err != nil {
		fmt.Printf("Error connecting to database: %v", err)
		return err
	}
	defer db.Close()

	updateRequest := `UPDATE requests SET Status = ?, User_status = ? WHERE ID = ?`
	_, err = db.Exec(updateRequest, Status, User_status, ID)
	if err != nil {
		fmt.Printf("Error updating database: %v", err)
		return err
	}
	return nil
}

func ExecuteRequest(ID int) error {
	db, err := connection()
	if err != nil {
		fmt.Printf("Error connecting to database: %v", err)
		return err
	}
	defer db.Close()

	request := FetchRequests("", "", "", "", ID, false)[0]
	if request.Status != "disapproved" {
		return nil
	}
	if request.Status == "approved" {
		if request.Title == "" {
			err = update_user(request.Username, "admin")
		} else {
			book := FetchBooks(request.Title, "", "", 0)[0]
			if request.Request == "checkout" {
				err = UpdateBook(book.Quantity-1, request.BookID)
				if err != nil {
					return err
				}
				err = InsertBorrowingHistory(types.BorrowingHistory{BookID: request.BookID, Title: request.Title, Username: request.Username, Borrowed_date: time.Now(), Returned_date: time.Time{}})
			} else if request.Request == "checkin" {
				err = UpdateBook(book.Quantity+1, request.BookID)
				if err != nil {
					return err
				}
				err = UpdateBorrowingHistory(request.BookID, time.Now().String())
			}
		}
		return err
	}
	return nil
}
