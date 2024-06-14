package services

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"mandalart.com/repositories"
)

type SheetService struct {
	queries *repositories.Queries
}

type Cell struct {
	Id          int32  `json:"id"`
	Color       *string `json:"color"`
	Goal        *string `json:"goal"`
	IsCompleted bool   `json:"isCompleted"`
}

type SheetWithMain struct {
	Id    int32  `json:"id"`
	Name  string `json:"name"`
	Cells []Cell `json:"cells"`
}

func NewSheetService(ctx context.Context) (*SheetService, error) {
	conn, ok := ctx.Value("db").(*pgxpool.Pool)
	if !ok {
		return nil, fmt.Errorf("database is not initialized")
	}
	return &SheetService{repositories.New(conn)}, nil
}

func (s *SheetService) GetSheetWithMainCellsById(ctx context.Context,ownerID int32) (*SheetWithMain, error) {
	
	data, err := s.queries.GetLatestSheetWithMainCellsByOwnerId(ctx, &ownerID)
	if err != nil {
		log.Println("Error fetching sheet data:", err)
		return nil, err
	}
	if len(data) == 0 {
		return nil, nil
	}
	sheet := &SheetWithMain{
		Id:    data[0].ID,
		Name:  *data[0].Name,
		Cells: make([]Cell, len(data)),
	}

	for i, el := range data {
		sheet.Cells[i] = Cell{
			Id:          el.CellID,
			Color:       el.Color,
			Goal:        el.Goal,
			IsCompleted: el.IsCompleted,
		}
	}


	return sheet, nil
}
