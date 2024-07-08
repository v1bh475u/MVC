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
	n_requests := n_requestedbooks(book.BookID)
	if !isQuantityValid(quantity, book.Quantity, n_requests) {
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

func isQuantityValid(quantity, curr_quantity, n_requests int) bool {
	return quantity+curr_quantity-n_requests >= 0
}

func n_requestedbooks(BookID sql.NullInt64) int {
	return len(models.FetchRequests("", "", "", "pending", BookID, false))
}
