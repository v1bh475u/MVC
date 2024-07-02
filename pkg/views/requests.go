package views

import "html/template"

func Requests() *template.Template {
	temp := template.Must(template.ParseFiles("templates/requests.html"))
	return temp
}
