package types

import "time"

type User struct {
	ID       int
	Username string
	Password string
	Role     string
}

type Book struct {
	BookID   int
	Title    string
	Author   string
	Genre    string
	Quantity int
}

type BorrowingHistory struct {
	ID            int
	BookID        int
	Title         string
	Username      string
	Borrowed_date time.Time
	Returned_date time.Time
}

type Request struct {
	ID          int
	Username    string
	BookID      int
	Title       string
	Request     string
	Status      string
	User_status string
	Date        time.Time
}

type DisplayBook struct {
	Book   Book
	Status string
}

type BookCatalog struct {
	Books    []DisplayBook
	Genres   []string
	Authors  []string
	Username string
	Role     string
	Messages int
}

type DetailedBook struct {
	Book             Book
	Status           string
	BorrowingHistory []BorrowingHistory
	Role             string
}

type Message struct {
	Message string
	Type    string
}

type PageData struct {
	Message string
}
