package controller

import (
	"net/http"

	"github.com/v1bh475u/LibMan_MVC/pkg/types"
	"github.com/v1bh475u/LibMan_MVC/pkg/views"
)

func Home(w http.ResponseWriter, r *http.Request) {
	t := views.StartPage()
	t.Execute(w, types.PageData{Message: ""})
}

func Logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:   "token",
		Value:  "",
		MaxAge: -1,
	})
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
