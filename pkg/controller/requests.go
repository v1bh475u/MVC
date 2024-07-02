package controller

import (
	"encoding/json"
	"net/http"

	"github.com/v1bh475u/LibMan_MVC/pkg/models"
	"github.com/v1bh475u/LibMan_MVC/pkg/types"
	"github.com/v1bh475u/LibMan_MVC/pkg/views"
)

func Requests(w http.ResponseWriter, r *http.Request) {
	Requests := models.FetchRequests("", "", "", "pending", 0, false)
	t := views.Requests()
	t.Execute(w, Requests)
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
