package views

import (
	"fmt"
	"net/http"
	s "mandalart.com/services"
)

type SheetView struct {
	sheetService *s.SheetService
	cellService  *s.CellService
}

func NewSheetView(sheetService *s.SheetService, cellService *s.CellService) *SheetView {
	return &SheetView{sheetService, cellService}
}

func (v *SheetView) GetLatestSheetWithMainCells(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID, ok := ctx.Value("userID").(int)
	if !ok {
		Respond(w, r, http.StatusInternalServerError, fmt.Errorf("user id not found"))
		return
	}

	sheet, err := v.sheetService.GetSheetWithMainCellsById(ctx, userID)
	if err != nil {
		Respond(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	Respond(w, r, http.StatusOK, sheet)
}

func (v *SheetView) GetChildrenCells(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID, ok := ctx.Value("userID").(int)

	if !ok {
		Respond(w, r, http.StatusInternalServerError, fmt.Errorf("user id not found"))
		return
	}

	cellID, ok := ctx.Value("cellID").(int)
	if !ok {
		Respond(w, r, http.StatusBadRequest, fmt.Errorf("cell id not found"))
	}


	cells, err := v.cellService.GetChildrenCellsByParentID(ctx, userID, cellID)
	if err != nil {
		Respond(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	Respond(w, r, http.StatusOK, cells)
}

func (v *SheetView) CreateSheet(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID, ok := ctx.Value("userID").(int)
	if !ok {
		Respond(w, r, http.StatusUnauthorized, fmt.Errorf("user id not found"))
	}

	sheet, err := v.sheetService.CreateNewSheet(ctx, userID)
	if err != nil {
		Respond(w, r, http.StatusInternalServerError, err.Error())
	}
	Respond(w, r, http.StatusOK, sheet)
}
