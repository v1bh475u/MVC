package views

import (
	"html/template"
)

func StartPage() *template.Template {
	temp := template.Must(template.ParseFiles("templates/index.html"))
	return temp
}
