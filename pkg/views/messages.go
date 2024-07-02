package views

import "html/template"

func Messages() *template.Template {
	temp := template.Must(template.ParseFiles("templates/user-navbar.html", "templates/messages.html"))
	return temp
}
