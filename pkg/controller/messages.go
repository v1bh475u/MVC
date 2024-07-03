package controller

import (
	"database/sql"
	"net/http"

	"github.com/v1bh475u/LibMan_MVC/pkg/models"
	"github.com/v1bh475u/LibMan_MVC/pkg/types"
	"github.com/v1bh475u/LibMan_MVC/pkg/utils"
	"github.com/v1bh475u/LibMan_MVC/pkg/views"
)

func Messages(w http.ResponseWriter, r *http.Request) {
	username, _, err := utils.VerifyToken(r.Cookies()[0].Value)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	messages := models.FetchRequests(username, "", "", "approved", sql.NullInt64{}, true)
	messages = append(messages, models.FetchRequests(username, "", "", "disapproved", sql.NullInt64{}, true)...)
	updateMessages(messages)
	t := views.Messages()
	t.ExecuteTemplate(w, "messages", types.PageData{Messages: messages, Catalog: false})
}

func updateMessages(messages []types.DRequest) {
	for _, message := range messages {
		models.UpdateRequest("", "seen", message.ID)
	}
}
