package test

import (
	"net/http"

	"github.com/go-chi/render"
	s "mandalart.com/services"
)

func CreateSheet(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	sheetService, err := s.NewSheetService(ctx)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, err)
		return
	}

	sheet, err := sheetService.GetSheetWithMainCellsById(49)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, err)
		return
	} else {
		render.JSON(w, r, sheet)
		return
	}
}
