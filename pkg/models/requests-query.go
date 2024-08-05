package models

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/v1bh475u/LibMan_MVC/pkg/types"
)

func FetchRequests(username, request, title, status string, ID sql.NullInt64, User bool) []types.DRequest {
	db, err := connection()
	if err != nil {
		fmt.Printf("Error connecting to database: %v", err)
		return nil
	}
	defer db.Close()

	query := `SELECT * FROM requests`
	conditions := []string{}
	params := []interface{}{}
	if ID.Valid && ID.Int64 != 0 {
		conditions = append(conditions, `ID = ?`)
		params = append(params, ID.Int64)
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
		params = append(params, types.UNSEEN)
	}
	if conditions != nil {
		query += ` WHERE ` + strings.Join(conditions, ` AND `)
	}
	result, err := db.Query(query, params...)
	if err != nil {
		fmt.Printf("Error querying database: %v\n", err)
		return nil
	}
	var requests []types.DRequest
	for result.Next() {
		var req types.DRequest
		var dateBytes []byte
		var title sql.NullString
		err := result.Scan(&req.ID, &req.Username, &req.BookID, &title, &req.Request, &req.Status, &req.User_status, &dateBytes)
		if err != nil {
			fmt.Printf("Error scanning database: %v\n", err)
			return nil
		}
		req.Title = title.String
		date, err := time.Parse("2006-01-02", string(dateBytes))
		if err != nil {
			fmt.Printf("Error parsing date: %v\n", err)
			return nil
		}
		req.Date = date.Format("Mon Jan _2 15:04:05 2006")
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
	conditions := []string{}
	params := []interface{}{}
	insertRequest := `INSERT INTO requests`
	if req.BookID.Valid && req.BookID.Int64 != 0 {
		conditions = append(conditions, `BookID`)
		params = append(params, req.BookID)
	}
	if req.Title != "" {
		conditions = append(conditions, `Title`)
		params = append(params, req.Title)
	}
	if req.Request != "" {
		conditions = append(conditions, `Request`)
		params = append(params, req.Request)
	}
	if req.Status != "" {
		conditions = append(conditions, `Status`)
		params = append(params, req.Status)
	}
	if req.User_status != "" {
		conditions = append(conditions, `User_status`)
		params = append(params, req.User_status)
	}
	if req.Username != "" {
		conditions = append(conditions, `Username`)
		params = append(params, req.Username)
	}
	if !req.Date.IsZero() {
		conditions = append(conditions, `Date`)
		params = append(params, req.Date)
	}
	if len(conditions) > 0 {
		insertRequest += ` (` + strings.Join(conditions, `,`) + `)`
	}
	insertRequest += ` VALUES`
	if len(conditions) > 0 {
		insertRequest += ` (`
		for i := 0; i < len(conditions); i++ {
			insertRequest += `?`
			if i < len(conditions)-1 {
				insertRequest += `,`
			}
		}
		insertRequest += `)`
	}
	_, err = db.Exec(insertRequest, params...)
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
	conditions := []string{}
	params := []interface{}{}
	if Status != "" {
		conditions = append(conditions, `Status = ?`)
		params = append(params, Status)
	}
	if User_status != "" {
		conditions = append(conditions, `User_status = ?`)
		params = append(params, User_status)
	}
	params = append(params, ID)
	updateRequest := `UPDATE requests SET ` + strings.Join(conditions, `,`) + ` WHERE ID = ?`
	_, err = db.Exec(updateRequest, params...)
	if err != nil {
		fmt.Printf("Error updating database: %v", err)
		return err
	}
	return nil
}

func ExecuteRequest(ID sql.NullInt64) error {
	db, err := connection()
	if err != nil {
		fmt.Printf("Error connecting to database: %v", err)
		return err
	}
	defer db.Close()

	request := FetchRequests("", "", "", "", ID, false)[0]
	if request.Status == types.DISAPPROVED {
		return nil
	}
	if request.Status == types.APPROVED {
		if request.Title == "" {
			err = update_user(request.Username, types.ADMIN)
		} else {
			book := FetchBooks(request.Title, "", "", 0)[0]
			if request.Request == types.CHECKOUT {
				if book.Quantity == 0 {
					UpdateRequest(types.DISAPPROVED, types.UNSEEN, int(ID.Int64))
					return fmt.Errorf("book not available")
				}
				err = UpdateBook(book.Quantity-1, request.BookID)
				if err != nil {
					return err
				}
				err = InsertBorrowingHistory(types.BorrowingHistory{BookID: request.BookID, Title: request.Title, Username: request.Username, Borrowed_date: time.Now()})
			} else if request.Request == types.CHECKIN {
				err = UpdateBook(book.Quantity+1, request.BookID)
				if err != nil {
					return err
				}
				err = UpdateBorrowingHistory(request.BookID, time.Now(), request.Username)
			}
		}
		return err
	}
	return nil
}
