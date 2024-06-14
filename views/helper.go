package views

import (
	"net/http"

	"github.com/go-chi/render"
)

func respond(w http.ResponseWriter, r *http.Request, status int, data interface{}) {
	render.Status(r, status)
	render.JSON(w, r, data)
}