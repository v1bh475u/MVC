package controller

import (
	"net/http"

	"github.com/v1bh475u/LibMan_MVC/pkg/types"
	"github.com/v1bh475u/LibMan_MVC/pkg/views"
)

func SysMessages(message types.Message, w http.ResponseWriter, r *http.Request) {
	t := views.Sysmessages()
	t.Execute(w, message)
}
