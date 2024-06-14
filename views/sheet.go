package views

import (
	"fmt"
	"net/http"
	s "mandalart.com/services"
)

func GetLatestSheetWithMainCells(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID, ok := ctx.Value("userID").(int)
	if !ok {
		respond(w, r, http.StatusInternalServerError, fmt.Errorf("user id not found"))
		return
	}

	sheetService, err := s.NewSheetService(ctx)
	if err != nil {
		respond(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	sheet, err := sheetService.GetSheetWithMainCellsById(ctx, int32(userID))
	if err != nil {
		respond(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	respond(w, r, http.StatusOK, sheet)
}