package controller

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/v1bh475u/LibMan_MVC/pkg/models"
	"github.com/v1bh475u/LibMan_MVC/pkg/types"
	"github.com/v1bh475u/LibMan_MVC/pkg/utils"
	"github.com/v1bh475u/LibMan_MVC/pkg/views"
)

func Requests(w http.ResponseWriter, r *http.Request) {
	Requests := models.FetchRequests("", "", "", "pending", 0, false)
	t := views.Requests()
	t.ExecuteTemplate(w, "requests", types.PageData{Messages: Requests, Catalog: false})
}

func PostRequests(w http.ResponseWriter, r *http.Request) {
	var body map[int]string
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		SysMessages(types.Message{Message: "Invalid request", Type: "Error"}, w, r)
		return
	}
	for key, value := range body {
		if value == "approved" {
			err = models.ExecuteRequest(key)
		}
		if err != nil {
			SysMessages(types.Message{Message: "Error executing request", Type: "Error"}, w, r)
			return
		}
		err = models.UpdateRequest(value, "unseen", key)
		if err != nil {
			SysMessages(types.Message{Message: "Error updating request", Type: "Error"}, w, r)
			return
		}
	}
	SysMessages(types.Message{Message: "Request(s) updated successfully", Type: "Info"}, w, r)
}

func AdminRequest(w http.ResponseWriter, r *http.Request) {
	username, _, err := utils.VerifyToken(r.Cookies()[0].Value)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	Requests := models.FetchRequests(username, "adminPrivs", "", "pending", 0, false)
	if isRequested(Requests) {
		SysMessages(types.Message{Message: "You have already requested for admin privileges", Type: "Warning"}, w, r)
		return
	}
	request := types.Request{Request: "adminPrivs", Status: "pending", User_status: "unseen", Username: username, Date: time.Now().Format("Mon Jan _2 15:04:05 2006")}
	err = models.InsertRequest(request)
	if err != nil {
		SysMessages(types.Message{Message: "Error submitting request", Type: "Error"}, w, r)
		return
	}
	SysMessages(types.Message{Message: "Request submitted successfully", Type: "Info"}, w, r)
}
