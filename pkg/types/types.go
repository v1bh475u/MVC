package types

import (
	"database/sql"
	"time"
)
const (
	ADMIN = "admin"
	USER  = "user"
	ADMINPRIVS = "adminPrivs"
	PENDING = "pending"
	APPROVED = "approved"
	DISAPPROVED = "disapproved"
	SEEN = "seen"
	UNSEEN = "unseen"
	CHECKOUT = "checkout"
	CHECKIN = "checkin"
	BORROWED = "borrowed"
	REQUESTED = "requested"
	AVAILABLE = "available"
)
type User struct {
	ID       int
	Username string
	Password string
	Role     string
}

type Book struct {
	BookID   sql.NullInt64
	Title    string
	Author   string
	Genre    string
	Quantity int
}

type DBorrowingHistory struct {
	ID            int
	BookID        sql.NullInt64
	Title         string
	Username      string
	Borrowed_date string
	Returned_date string
}

type BorrowingHistory struct {
	ID            int
	BookID        sql.NullInt64
	Title         string
	Username      string
	Borrowed_date time.Time
	Returned_date time.Time
}

type DRequest struct {
	ID          int
	Username    string
	BookID      sql.NullInt64
	Title       string
	Request     string
	Status      string
	User_status string
	Date        string
}
type Request struct {
	ID          int
	Username    string
	BookID      sql.NullInt64
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
	Catalog  bool
}

type DetailedBook struct {
	Book             Book
	Status           string
	BorrowingHistory []DBorrowingHistory
	Role             string
	Catalog          bool
}

type Message struct {
	Message string
	Type    string
}

type PageData struct {
	Message          string
	Messages         []DRequest
	Books            []Book
	BorrowingHistory []DBorrowingHistory
	Catalog          bool
}
