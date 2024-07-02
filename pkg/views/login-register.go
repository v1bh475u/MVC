package views

import "html/template"

func LoginPage() *template.Template {
	temp := template.Must(template.ParseFiles("templates/login.html"))
	return temp
}
func RegisterPage() *template.Template {
	temp := template.Must(template.ParseFiles("templates/register.html"))
	return temp
}
