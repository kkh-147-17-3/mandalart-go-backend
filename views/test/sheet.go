package test

import (
	"context"
	"github.com/go-chi/render"
	"github.com/jackc/pgx/v5"
	"mandalart.com/repositories"
	"mandalart.com/services"
	"net/http"
	"os"
)

func CreateSheet(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, err.Error())
		return
	}
	defer conn.Close(ctx)

	queries := repositories.New(conn)
	sheetService := services.SheetService{queries, &ctx}
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
