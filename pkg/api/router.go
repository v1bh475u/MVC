package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/v1bh475u/LibMan_MVC/pkg/controller"
)

func StartApi() {
	router := mux.NewRouter()

	fs := http.FileServer(http.Dir("./static"))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	router.HandleFunc("/", controller.Home).Methods("GET")

	//Login-Register routes

	router.HandleFunc("/login", controller.Login).Methods("POST")
	router.HandleFunc("/login", controller.LoginPage).Methods("GET")

	router.HandleFunc("/register", controller.RegisterPage).Methods("GET")
	router.HandleFunc("/register", controller.Register).Methods("POST")
	router.HandleFunc("/logout", controller.Logout).Methods("GET")

	//Catalog routes

	GetBooks := http.HandlerFunc(controller.GetBooks)
	router.Handle("/books", controller.AuthMiddleware(GetBooks)).Methods("GET")
	PostBooks := http.HandlerFunc(controller.PostBooks)
	router.Handle("/books", controller.AuthMiddleware(PostBooks)).Methods("POST")

	//Book details routes

	GetBook := http.HandlerFunc(controller.GetBook)
	router.Handle("/books/{id}", controller.AuthMiddleware(GetBook)).Methods("GET")
	BookRequest := http.HandlerFunc(controller.BookRequest)
	router.Handle("/checkout", controller.AuthMiddleware(BookRequest)).Methods("POST")
	router.Handle("/checkin", controller.AuthMiddleware(BookRequest)).Methods("POST")

	Messages := http.HandlerFunc(controller.Messages)
	router.Handle("/messages", controller.AuthMiddleware(Messages)).Methods("GET")

	BorrowingHistory := http.HandlerFunc(controller.BorrowingHistory)
	router.Handle("/borrowinghistory", controller.AuthMiddleware(BorrowingHistory)).Methods("GET")

	AdminReq := http.HandlerFunc(controller.AdminRequest)
	router.Handle("/reqAdmin", controller.AuthMiddleware(AdminReq)).Methods("POST")

	//Admin routes

	Requests := http.HandlerFunc(controller.Requests)
	router.Handle("/requests", controller.AuthMiddleware(controller.AdminMiddleware(Requests))).Methods("GET")
	PostRequests := http.HandlerFunc(controller.PostRequests)
	router.Handle("/apply-changes", controller.AuthMiddleware(controller.AdminMiddleware(PostRequests))).Methods("POST")

	BookManagement := http.HandlerFunc(controller.BookManagement)
	router.Handle("/book-management", controller.AuthMiddleware(controller.AdminMiddleware(BookManagement))).Methods("GET")
	AddBook := http.HandlerFunc(controller.AddBook)
	router.Handle("/addbook", controller.AuthMiddleware(controller.AdminMiddleware(AddBook))).Methods("POST")
	UpdateBook := http.HandlerFunc(controller.UpdateBook)
	router.Handle("/updatebook", controller.AuthMiddleware(controller.AdminMiddleware(UpdateBook))).Methods("POST")

	http.ListenAndServe(":8080", router)
}
