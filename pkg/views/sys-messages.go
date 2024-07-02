package views

import "text/template"

func Sysmessages() *template.Template {
	temp := template.Must(template.ParseFiles("templates/sys-messages.html"))
	return temp
}
