package types

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
	Borrowed_date string
	Returned_date string
}

type Request struct {
	ID          int
	Username    string
	BookID      int
	Title       string
	Request     string
	Status      string
	User_status string
	Date        string
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
	BorrowingHistory []BorrowingHistory
	Role             string
	Catalog          bool
}

type Message struct {
	Message string
	Type    string
}

type PageData struct {
	Message          string
	Messages         []Request
	Books            []Book
	BorrowingHistory []BorrowingHistory
	Catalog          bool
}
