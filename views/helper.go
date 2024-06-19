package views

import (
	"net/http"

	"github.com/go-chi/render"
)

type ErrorResponse struct {
	Code int `json:"code"`
	Message  string `json:"message"`
}

func Respond(w http.ResponseWriter, r *http.Request, status int, data interface{}) {
	if err, ok := data.(error); ok {
		render.Status(r, status)
		render.JSON(w, r, ErrorResponse{Message: err.Error()})
		return
	}

	render.Status(r, status)
	render.JSON(w, r, data)
}
