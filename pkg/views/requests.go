package views

import "html/template"

func Requests() *template.Template {
	temp := template.Must(template.ParseFiles("templates/admin-navbar.html", "templates/requests.html"))
	return temp
}
