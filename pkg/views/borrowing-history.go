package views

import "html/template"

func BorrowingHistory() *template.Template {
	temp := template.Must(template.ParseFiles("templates/user-navbar.html", "templates/borrowing-history.html"))
	return temp
}
