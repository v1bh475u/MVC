package controller

import (
	"database/sql"
	"net/http"
	"strconv"
	"github.com/v1bh475u/LibMan_MVC/pkg/models"
	"github.com/v1bh475u/LibMan_MVC/pkg/types"
	"github.com/v1bh475u/LibMan_MVC/pkg/views"
)

func BookManagement(w http.ResponseWriter, r *http.Request) {
	books := models.FetchBooks("", "", "", 0)
	t := views.BookManagement()
	t.ExecuteTemplate(w, "book-management", types.PageData{Books: books, Catalog: false})
}

func AddBook(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	author := r.FormValue("author")
	genre := r.FormValue("genre")
	quantity, err := strconv.Atoi(r.FormValue("quantity"))
	if err != nil {
		SysMessages(types.Message{Message: "Invalid quantity", Type: "Error"}, w, r)
		return
	}
	book := types.Book{Title: title, Author: author, Genre: genre, Quantity: quantity}
	if l := len(models.FetchBooks(title, author, genre, 0)); l > 0 {
		SysMessages(types.Message{Message: "Book already exists", Type: "Warning"}, w, r)
		return
	}
	err = models.InsertBook(book)
	if err != nil {
		SysMessages(types.Message{Message: "Error adding book", Type: "Error"}, w, r)
		return
	}
	SysMessages(types.Message{Message: "Book added successfully", Type: "Info"}, w, r)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	quantity, err := strconv.Atoi(r.FormValue("quantity"))
	if err != nil {
		SysMessages(types.Message{Message: "Invalid quantity", Type: "Error"}, w, r)
		return
	}
	book := models.FetchBooks(title, "", "", 0)[0]
	if !isQuantityValid(quantity, book.Quantity) {
		SysMessages(types.Message{Message: "Invalid quantity", Type: "Error"}, w, r)
		return
	}
	err = models.UpdateBook(book.Quantity+quantity, book.BookID)
	if err != nil {
		SysMessages(types.Message{Message: "Error updating book", Type: "Error"}, w, r)
		return
	}
	SysMessages(types.Message{Message: "Book updated successfully", Type: "Info"}, w, r)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	book := models.FetchBooks(title, "", "", 0)[0]
	book_history:=models.FetchBorrowingHistory("",title)
	n_borrowedbooks:=0
	for _, bh := range book_history {
		if bh.Returned_date == "Mon Jan  1 00:00:00 0001" {
			n_borrowedbooks++
		}
	}
	if n_borrowedbooks > 0 {
		SysMessages(types.Message{Message: "Book is borrowed. Cannot delete", Type: "Error"}, w, r)
		return
	}
	requests:=models.FetchRequests("", "", title, types.PENDING, sql.NullInt64{}, false)
	for _, request := range requests {
		models.UpdateRequest(types.DISAPPROVED, types.UNSEEN, request.ID)
	}
	err := models.DeleteBook(book.BookID)
	if err != nil {
		SysMessages(types.Message{Message: "Error deleting book", Type: "Error"}, w, r)
		return
	}
	SysMessages(types.Message{Message: "Book deleted successfully", Type: "Info"}, w, r)
}

func isQuantityValid(quantity, curr_quantity int) bool {
	return quantity+curr_quantity >= 0
}

func n_requestedbooks(title string) int {
	return len(models.FetchRequests("", "", title, types.PENDING, sql.NullInt64{}, false))
}
