package views

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
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
		respond(w, r, http.StatusInternalServerError, fmt.Errorf("user id not found"))
		return
	}

	sheet, err := v.sheetService.GetSheetWithMainCellsById(ctx, int32(userID))
	if err != nil {
		respond(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	respond(w, r, http.StatusOK, sheet)
}

func (v *SheetView) GetChildrenCells(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID, ok := ctx.Value("userID").(int)

	if !ok {
		respond(w, r, http.StatusInternalServerError, fmt.Errorf("user id not found"))
		return
	}

	cellID, err := strconv.Atoi(chi.URLParam(r, "cellID"))
	if err != nil {
		respond(w, r, http.StatusBadRequest, err.Error())
		return
	}

	cells, err := v.cellService.GetChildrenCellsByParentID(ctx, int32(userID), int32(cellID))
	if err != nil {
		respond(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	respond(w, r, http.StatusOK, cells)
}

func (v *SheetView) CreateSheet(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID, ok := ctx.Value("userID").(int)
	if !ok {
		respond(w, r, http.StatusUnauthorized, fmt.Errorf("user id not found"))
	}

	sheet, err := v.sheetService.CreateNewSheet(ctx, int32(userID))
	if err != nil {
		respond(w, r, http.StatusInternalServerError, err.Error())
	}
	respond(w, r, http.StatusOK, sheet)
}
