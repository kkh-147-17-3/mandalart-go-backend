package services

import (
	"context"
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
type Cell struct {
	Id          int32  `json:"id"`
	Color       string `json:"color"`
	Goal        string `json:"goal"`
	IsCompleted bool   `json:"isCompleted"`
}

type SheetWithMain struct {
	Id    int32  `json:"id"`
	Name  string `json:"name"`
	Cells []Cell `json:"cells"`
}

func (s *SheetService) GetSheetWithMainCellsById(ownerID int) (*SheetWithMain, error) {

	data, err := s.Queries.GetLatestSheetWithMainCellsByOwnerId(*s.Ctx, int32(ownerID))
	if err != nil {
		log.Println(err)
		return nil, err
	}
	if len(data) == 0 {
		return nil, nil
	}
	sheet := &SheetWithMain{Id: data[0].ID, Name: data[0].Name.String, Cells: []Cell{}}
	for _, el := range data {
		cell := Cell{el.CellID, el.Color.String, el.Goal.String, el.IsCompleted}
		sheet.Cells = append(sheet.Cells, cell)
	}

	return sheet, nil
}
