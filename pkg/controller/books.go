package controller

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/v1bh475u/LibMan_MVC/pkg/models"
	"github.com/v1bh475u/LibMan_MVC/pkg/types"
	"github.com/v1bh475u/LibMan_MVC/pkg/utils"
	"github.com/v1bh475u/LibMan_MVC/pkg/views"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
	Books := models.FetchBooks("", "", "", 0)
	username, _, err := utils.VerifyToken(r.Cookies()[0].Value)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	BookCatalog := prepareBookCatalog(username, Books)
	t := views.BookCatalog()
	if err = t.ExecuteTemplate(w, "book-catalog", BookCatalog); err != nil {
		fmt.Println("Error executing template: ", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

func PostBooks(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	author := r.FormValue("author")
	genre := r.FormValue("genre")
	books := models.FetchBooks(title, author, genre, 0)
	username, _, err := utils.VerifyToken(r.Cookies()[0].Value)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	BookCatalog := prepareBookCatalog(username, books)
	t := views.BookCatalog()
	if err = t.ExecuteTemplate(w, "book-catalog", BookCatalog); err != nil {
		fmt.Println("Error executing template: ", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookid, err := strconv.Atoi(vars["id"])
	fmt.Printf("Book ID: %d\n", bookid)
	if err != nil {
		http.Redirect(w, r, "/books", http.StatusSeeOther)
		return
	}
	book := models.FetchBooks("", "", "", bookid)[0]
	username, role, err := utils.VerifyToken(r.Cookies()[0].Value)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	status := book_status(username, book.Title)
	borrowinghistory := models.FetchBorrowingHistory("", book.Title)
	t := views.BookDetails()
	if err = t.ExecuteTemplate(w, "book-details", types.DetailedBook{Book: book, Status: status, BorrowingHistory: borrowinghistory, Role: role, Catalog: false}); err != nil {
		fmt.Println("Error executing template: ", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

func book_status(username, title string) string {
	requests := models.FetchRequests(username, "", title, "pending", sql.NullInt64{Int64: 0}, false)
	if len(requests) > 0 {
		return "Requested"
	}
	borrowinghistory := models.FetchBorrowingHistory(username, title)
	if isBorrowed(borrowinghistory) {
		return "Borrowed"
	}
	return "Available"
}

func isBorrowed(borrowing_history []types.DBorrowingHistory) bool {
	for _, history := range borrowing_history {
		if history.Returned_date == "Mon Jan  1 00:00:00 0001" {
			return true
		}
	}
	return false
}

func isRequested(requests []types.DRequest) bool {
	for _, request := range requests {
		if request.Status == "pending" {
			return true
		}
	}
	return false
}

func BookRequest(w http.ResponseWriter, r *http.Request) {
	username, _, err := utils.VerifyToken(r.Cookies()[0].Value)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	bookid, err := strconv.Atoi(r.FormValue("bookId"))
	if err != nil {
		http.Redirect(w, r, "/books", http.StatusSeeOther)
		return
	}
	action := r.FormValue("action")
	book := models.FetchBooks("", "", "", bookid)[0]
	requests := models.FetchRequests(username, action, book.Title, "pending", sql.NullInt64{}, false)
	if isRequested(requests) {
		SysMessages(types.Message{Message: "You have already requested this book", Type: "Warning"}, w, r)
		return
	}
	borrowinghistory := models.FetchBorrowingHistory(username, book.Title)
	if action == "checkout" {
		if isBorrowed(borrowinghistory) {
			SysMessages(types.Message{Message: "You have already borrowed this book", Type: "Warning"}, w, r)
			return
		}
	} else if action == "checkin" {
		if !isBorrowed(borrowinghistory) {
			SysMessages(types.Message{Message: "You have not borrowed this book", Type: "Warning"}, w, r)
			return
		}
	}
	request := types.Request{BookID: book.BookID, Title: book.Title, Request: action, Status: "pending", User_status: "unseen", Username: username, Date: time.Now()}
	err = models.InsertRequest(request)
	if err != nil {
		SysMessages(types.Message{Message: err.Error(), Type: "Error"}, w, r)
		return
	}
	SysMessages(types.Message{Message: "Request submitted successfully", Type: "Info"}, w, r)
}

func prepareBookCatalog(username string, Books []types.Book) types.BookCatalog {
	BookList := []types.DisplayBook{}
	for _, book := range Books {
		BookList = append(BookList, types.DisplayBook{Book: book, Status: book_status(username, book.Title)})
	}
	genres := models.FetchUniqueitems("Genre")
	authors := models.FetchUniqueitems("Author")
	messages := models.FetchRequests(username, "", "", "approved", sql.NullInt64{}, true)
	messages = append(messages, models.FetchRequests(username, "", "", "disapproved", sql.NullInt64{}, true)...)
	n_messages := len(messages)
	user, _ := models.FetchUser(username)
	role := user.Role
	return types.BookCatalog{Books: BookList, Genres: genres, Authors: authors, Username: username, Role: role, Messages: n_messages, Catalog: true}
}
