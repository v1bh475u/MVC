package middleware

import (
	"net/http"
	"strings"
	"fmt"
	"github.com/v1bh475u/LibMan_MVC/pkg/models"
	"github.com/v1bh475u/LibMan_MVC/pkg/utils"
	"github.com/v1bh475u/LibMan_MVC/pkg/types"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Cookie") == "" {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		token := strings.Split(r.Header.Get("Cookie"), "=")[1]
		if token == "" {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		username, role, err := utils.VerifyToken(token)
		if username == "" || role == "" || err != nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func AdminMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Cookie") == "" {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		token := strings.Split(r.Header.Get("Cookie"), "=")[1]
		if token == "" {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		username, role, err := utils.VerifyToken(token)
		if username == "" || role == "" || err != nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		user, err := models.FetchUser(username)
		if err != nil {
			return
		}
		if user.Role != types.ADMIN {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func LogonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		fmt.Printf("%v\n",r.Header.Get("Cookie"))
		if r.Header.Get("Cookie") == "" {
			next.ServeHTTP(w, r)			
			return
		}
		token := strings.Split(r.Header.Get("Cookie"), "=")[1]
		fmt.Printf("%v\n",token)
		if token == "" {
			next.ServeHTTP(w, r)
			return
		}
		username, _, err := utils.VerifyToken(token)
		fmt.Printf("%v\n",username)
		if username == "" || err != nil {
			next.ServeHTTP(w, r)
			return
		}
		http.Redirect(w, r, "/books", http.StatusSeeOther)
	})
}