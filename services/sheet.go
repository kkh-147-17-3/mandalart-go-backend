package services

import (
	"context"
	"log"
	repo "mandalart.com/repositories"
	"mandalart.com/utils"
)

const CellNums = 8

type Todo struct {
	OwnerID int32
	CellID  int32
	Content *string
}

type SheetService struct {
	q *repo.Queries
}

type Cell struct {
	ID          int   `json:"id"`
	Color       *string `json:"color"`
	Goal        *string `json:"goal"`
	IsCompleted bool    `json:"isCompleted"`
}

type SheetWithMain struct {
	ID    int   `json:"id"`
	Name  *string `json:"name"`
	Cells []Cell  `json:"cells"`
}

func NewSheetService(q *repo.Queries) *SheetService {
	return &SheetService{q}
}

func (s *SheetService) GetSheetWithMainCellsById(ctx context.Context, ownerID int) (*SheetWithMain, error) {
	data, err := s.q.GetLatestSheetWithMainCellsByOwnerId(ctx, ownerID)
	if err != nil {
		log.Println("Error fetching sheet data:", err)
		return nil, err
	}
	if len(data) == 0 {
		return nil, nil
	}

	sheet := &SheetWithMain{
		ID:    data[0].ID,
		Name:  data[0].Name,
		Cells: make([]Cell, len(data)),
	}

	for i, el := range data {
		sheet.Cells[i] = Cell{
			ID:          el.CellID,
			Color:       el.Color,
			Goal:        el.Goal,
			IsCompleted: el.IsCompleted,
		}
	}

	return sheet, nil
}

func (s *SheetService) CreateNewSheet(ctx context.Context, ownerID int) (*SheetWithMain, error) {
	tx, err := utils.DBPool.Begin(ctx)
	if err != nil {
		return nil, err
	}

	sheetID, err := s.q.WithTx(tx).CreateSheet(ctx, repo.CreateSheetParams{
		OwnerID: ownerID,
	})
	if err != nil {
		tx.Rollback(ctx)
		return nil, err
	}

	parentCellID, err := s.q.WithTx(tx).CreateCell(ctx, repo.CreateCellParams{
		OwnerID:     ownerID,
		SheetID:     sheetID,
		Step:        1,
		Order:       0,
		IsCompleted: false,
	})
	if err != nil {
		tx.Rollback(ctx)
		return nil, err
	}

	parentCellIDs := make([]int, CellNums)
	for i := 0; i < CellNums; i++ {
		cellID, err := s.q.WithTx(tx).CreateCell(ctx, repo.CreateCellParams{
			OwnerID:     ownerID,
			SheetID:     sheetID,
			Step:        2,
			Order:       i,
			ParentID:    parentCellID,
			IsCompleted: false,
		})
		if err != nil {
			return nil, err
		}
		parentCellIDs[i] = cellID
	}

	for i := 0; i < CellNums; i++ {
		for j := i + 1; j < CellNums; j++ {
			if _, err := s.q.WithTx(tx).CreateCell(ctx, repo.CreateCellParams{
				OwnerID:     ownerID,
				SheetID:     sheetID,
				Step:        3,
				Order:       j,
				ParentID:    parentCellIDs[i],
				IsCompleted: false,
			}); err != nil {
				return nil, err
			}
		}
	}

	if err := tx.Commit(ctx); err != nil {
		_ = tx.Rollback(ctx)
		return nil, err
	}

	return s.GetSheetWithMainCellsById(ctx, ownerID)
}
