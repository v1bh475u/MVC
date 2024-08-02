package controller

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/v1bh475u/LibMan_MVC/pkg/models"
	"github.com/v1bh475u/LibMan_MVC/pkg/types"
	"github.com/v1bh475u/LibMan_MVC/pkg/utils"
	"github.com/v1bh475u/LibMan_MVC/pkg/views"
)

func Requests(w http.ResponseWriter, r *http.Request) {
	Requests := models.FetchRequests("", "", "", types.PENDING, sql.NullInt64{}, false)
	t := views.Requests()
	t.ExecuteTemplate(w, "requests", types.PageData{Messages: Requests, Catalog: false})
}

func PostRequests(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		SysMessages(types.Message{Message: "Error parsing form", Type: "Error"}, w, r)
		return
	}
	body := make(map[int]string)
	for key, value := range r.PostForm {
		fmt.Printf("key: %v, value: %v\n", key, value)
		k, err := strconv.Atoi(key)
		if err != nil {
			SysMessages(types.Message{Message: "Error converting key to int", Type: "Error"}, w, r)
			return
		}
		body[k] = value[0]
	}
	var err error
	for key, value := range body {
		err = models.UpdateRequest(value, types.UNSEEN, key)
		if err != nil {
			SysMessages(types.Message{Message: "Error updating request", Type: "Error"}, w, r)
			return
		}
		if value == types.APPROVED {
			k := sql.NullInt64{Int64: int64(key), Valid: true}
			err = models.ExecuteRequest(k)
		}
		if err != nil {
			if err == fmt.Errorf("book not available") {
				models.UpdateRequest(types.DISAPPROVED, types.UNSEEN, key)
				SysMessages(types.Message{Message: "Book not available", Type: "Warning"}, w, r)
				return
			}
			SysMessages(types.Message{Message: "Error executing request", Type: "Error"}, w, r)
			return
		}
	}
	SysMessages(types.Message{Message: "Request(s) updated successfully", Type: "Info"}, w, r)
}

func AdminRequest(w http.ResponseWriter, r *http.Request) {
	username, _, err := utils.VerifyToken(r.Cookies()[0].Value)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	borrowing_history:=models.FetchBorrowingHistory(username, "")

	if anyBookBorrowed(borrowing_history) {
		SysMessages(types.Message{Message: "You have borrowed books. Return them before requesting for admin privileges", Type: "Warning"}, w, r)
		return
	}
	Requests := models.FetchRequests(username, types.ADMINPRIVS, "", types.PENDING, sql.NullInt64{}, false)
	if isRequested(Requests) {
		SysMessages(types.Message{Message: "You have already requested for admin privileges", Type: "Warning"}, w, r)
		return
	}
	request := types.Request{Request: types.ADMINPRIVS, Status: types.PENDING, User_status: types.UNSEEN, Username: username, Date: time.Now()}
	err = models.InsertRequest(request)
	if err != nil {
		SysMessages(types.Message{Message: "Error submitting request", Type: "Error"}, w, r)
		return
	}
	SysMessages(types.Message{Message: "Request submitted successfully", Type: "Info"}, w, r)
}

func anyBookBorrowed(borrowing_history []types.DBorrowingHistory) bool {
	for _, history := range borrowing_history {
		if history.Returned_date == "Mon Jan  1 00:00:00 0001" {
			return true
		}
	}
	return false
}