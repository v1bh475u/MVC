package views

import "html/template"

func BookManagement() *template.Template {
	temp := template.Must(template.ParseFiles("templates/book-management.html"))
	return temp
}
