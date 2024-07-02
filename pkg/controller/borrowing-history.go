package controller

import (
	"net/http"

	"github.com/v1bh475u/LibMan_MVC/pkg/models"
	"github.com/v1bh475u/LibMan_MVC/pkg/utils"
	"github.com/v1bh475u/LibMan_MVC/pkg/views"
)

func BorrowingHistory(w http.ResponseWriter, r *http.Request) {
	username, _, err := utils.VerifyToken(r.Cookies()[0].Value)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	BorrowingHistory := models.FetchBorrowingHistory(username, "")
	t := views.BorrowingHistory()
	t.Execute(w, BorrowingHistory)
}
