package controller

import (
	"fmt"
	"net/http"
	"time"

	"github.com/v1bh475u/LibMan_MVC/pkg/models"
	"github.com/v1bh475u/LibMan_MVC/pkg/types"
	"github.com/v1bh475u/LibMan_MVC/pkg/utils"
	"github.com/v1bh475u/LibMan_MVC/pkg/views"
)

func LoginPage(w http.ResponseWriter, r *http.Request) {
	t := views.LoginPage()
	t.Execute(w, nil)
}

func RegisterPage(w http.ResponseWriter, r *http.Request) {
	t := views.RegisterPage()
	t.Execute(w, nil)
}

func Login(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	user, _ := models.FetchUser(username)
	if user.Username != username {
		t := views.LoginPage()
		fmt.Println("Invalid Username")
		message := "Invalid Username"
		t.Execute(w, types.PageData{Message: message})
		return
	}
	if !utils.CheckPassword(password, user.Password) {
		t := views.LoginPage()
		message := "Invalid Password"
		t.Execute(w, types.PageData{Message: message})
		return
	}
	token, err := utils.CreateToken(user)
	if err != nil {
		t := views.LoginPage()
		message := "Error creating token"
		t.Execute(w, types.PageData{Message: message})
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    token,
		Path:     "/",
		Domain:   "",
		Expires:  time.Now().Add(time.Hour * 24),
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	})
	http.Redirect(w, r, "/books", http.StatusSeeOther)
	fmt.Println("Cookie set")
}

func Register(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	user, _ := models.FetchUser(username)
	if user.Username == username {
		t := views.RegisterPage()
		message := "Username already exists"
		t.Execute(w, types.PageData{Message: message})
		return
	}
	password := r.FormValue("password")
	role := "user"
	hashedPassword := utils.HashPassword(password)
	user = types.User{
		Username: username,
		Password: hashedPassword,
		Role:     role,
	}
	err := models.InsertUser(user)
	if err != nil {
		t := views.RegisterPage()
		message := "Error inserting user"
		t.Execute(w, types.PageData{Message: message})
		return
	}
	fmt.Print("User inserted")
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
