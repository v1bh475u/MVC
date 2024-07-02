package controller

import (
	"fmt"
	"net/http"
	"strconv"

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
	fmt.Printf("BookCatalog: %v\n", BookCatalog)
	t := views.BookCatalog()
	if err = t.Execute(w, BookCatalog); err != nil {
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
	t.Execute(w, BookCatalog)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	bookid, err := strconv.Atoi(r.URL.Query().Get("id"))
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
	t.Execute(w, types.DetailedBook{Book: book, Status: status, BorrowingHistory: borrowinghistory, Role: role})
}

func book_status(username, title string) string {
	borrowinghistory := models.FetchBorrowingHistory(username, title)
	if isBorrowed(borrowinghistory) {
		return "Borrowed"
	}
	requests := models.FetchRequests(username, "", title, "pending", 0, false)
	if len(requests) > 0 {
		return "Requested"
	}
	return "Available"
}

func isBorrowed(borrowing_history []types.BorrowingHistory) bool {
	for _, history := range borrowing_history {
		if history.Returned_date.IsZero() {
			return true
		}
	}
	return false
}

func isRequested(requests []types.Request) bool {
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
	requests := models.FetchRequests(username, action, book.Title, "pending", 0, false)
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
	request := types.Request{BookID: book.BookID, Title: book.Title, Request: action, Status: "pending", User_status: "unseen", Username: username}
	err = models.InsertRequest(request)
	if err != nil {
		SysMessages(types.Message{Message: err.Error(), Type: "Error"}, w, r)
		return
	}
	SysMessages(types.Message{Message: "Request submitted successfully", Type: "Success"}, w, r)
}

func prepareBookCatalog(username string, Books []types.Book) types.BookCatalog {
	BookList := []types.DisplayBook{}
	for _, book := range Books {
		BookList = append(BookList, types.DisplayBook{Book: book, Status: book_status(username, book.Title)})
	}
	genres := models.FetchUniqueitems("Genre")
	authors := models.FetchUniqueitems("Author")
	messages := len(models.FetchRequests(username, "", "", "", 0, true))
	user, _ := models.FetchUser(username)
	role := user.Role
	return types.BookCatalog{Books: BookList, Genres: genres, Authors: authors, Username: username, Role: role, Messages: messages}
}
