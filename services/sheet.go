package services

import (
	"context"
	"github.com/jackc/pgx/v5/pgtype"
	"log"
	"mandalart.com/repositories"
)

//func CreateSheetByOwnerId(sheetID string) []schemas.Cell {
//
//	db := repo.New()
//	db.GetLatestSheetByOwnerId()
//
//	return cells
//}

type SheetService struct {
	Queries *repositories.Queries
	Ctx     *context.Context
}

type SheetWithMain struct {
	Sheet repositories.Sheet  `json:"sheet"`
	Cells []repositories.Cell `json:"cells"`
}

func (s *SheetService) GetSheetWithMainCellsById(ownerID int) (*SheetWithMain, error) {
	var (
		pOwnerID pgtype.Int4
		pSheetID pgtype.Int4
	)
	pOwnerID.Int32 = int32(ownerID)
	pOwnerID.Valid = true
	sheet, err := s.Queries.GetLatestSheetByOwnerId(*s.Ctx, pOwnerID)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	pSheetID.Int32 = int32(sheet.ID)
	pSheetID.Valid = true
	cells, err := s.Queries.GetMainCellsBySheetId(*s.Ctx, pSheetID)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &SheetWithMain{sheet, cells}, nil
}
