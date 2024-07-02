package views

import "html/template"

func BookCatalog() *template.Template {
	temp := template.Must(template.ParseFiles("templates/user-navbar.html", "templates/admin-navbar.html", "templates/book-catalog.html"))
	return temp
}

func BookDetails() *template.Template {
	temp := template.Must(template.ParseFiles("templates/user-navbar.html", "templates/admin-navbar.html", "templates/book-details.html"))
	return temp
}
