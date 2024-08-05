package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/v1bh475u/LibMan_MVC/pkg/controller"
	"github.com/v1bh475u/LibMan_MVC/pkg/middleware"

)

func StartApi() {
	router := mux.NewRouter()

	fs := http.FileServer(http.Dir("./static"))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	router.HandleFunc("/", controller.Home).Methods("GET")



	//Catalog routes
	GetBooks := http.HandlerFunc(controller.GetBooks)
	router.Handle("/books", middleware.AuthMiddleware(GetBooks)).Methods("GET")
	PostBooks := http.HandlerFunc(controller.PostBooks)
	router.Handle("/books", middleware.AuthMiddleware(PostBooks)).Methods("POST")

	//Login-Register routes
	LoginPage := http.HandlerFunc(controller.LoginPage)
	router.HandleFunc("/login", controller.Login).Methods("POST")
	router.Handle("/login", middleware.LogonMiddleware(LoginPage)).Methods("GET")
	router.HandleFunc("/register", controller.RegisterPage).Methods("GET")
	router.HandleFunc("/register", controller.Register).Methods("POST")
	router.HandleFunc("/logout", controller.Logout).Methods("GET")

	//Book details routes
	GetBook := http.HandlerFunc(controller.GetBook)
	router.Handle("/books/{id}", middleware.AuthMiddleware(GetBook)).Methods("GET")
	BookRequest := http.HandlerFunc(controller.BookRequest)
	router.Handle("/checkout", middleware.AuthMiddleware(BookRequest)).Methods("POST")
	router.Handle("/checkin", middleware.AuthMiddleware(BookRequest)).Methods("POST")

	Messages := http.HandlerFunc(controller.Messages)
	router.Handle("/messages", middleware.AuthMiddleware(Messages)).Methods("GET")

	BorrowingHistory := http.HandlerFunc(controller.BorrowingHistory)
	router.Handle("/borrowinghistory", middleware.AuthMiddleware(BorrowingHistory)).Methods("GET")

	AdminReq := http.HandlerFunc(controller.AdminRequest)
	router.Handle("/reqAdmin", middleware.AuthMiddleware(AdminReq)).Methods("POST")

	//Admin routes

	Requests := http.HandlerFunc(controller.Requests)
	router.Handle("/requests", middleware.AuthMiddleware(middleware.AdminMiddleware(Requests))).Methods("GET")
	PostRequests := http.HandlerFunc(controller.PostRequests)
	router.Handle("/apply-changes", middleware.AuthMiddleware(middleware.AdminMiddleware(PostRequests))).Methods("POST")

	BookManagement := http.HandlerFunc(controller.BookManagement)
	router.Handle("/book-management", middleware.AuthMiddleware(middleware.AdminMiddleware(BookManagement))).Methods("GET")
	AddBook := http.HandlerFunc(controller.AddBook)
	router.Handle("/addbook", middleware.AuthMiddleware(middleware.AdminMiddleware(AddBook))).Methods("POST")
	UpdateBook := http.HandlerFunc(controller.UpdateBook)
	router.Handle("/updatebook", middleware.AuthMiddleware(middleware.AdminMiddleware(UpdateBook))).Methods("POST")
	DeleteBook := http.HandlerFunc(controller.DeleteBook)
	router.Handle("/deletebook", middleware.AuthMiddleware(middleware.AdminMiddleware(DeleteBook))).Methods("POST")
	http.ListenAndServe(":8080", router)
}
