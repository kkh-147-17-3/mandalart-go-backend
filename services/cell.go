package services

import (
	"context"
	"fmt"
)

type CellService struct {
	*BaseService
}

func NewCellService(ctx context.Context) (*CellService, error){
	base,err := NewBaseService(ctx)
	if err != nil {
		return nil, err
	}
	return &CellService{base}, nil
}


func (c *CellService) GetChildrenCellsByParentID(ctx context.Context, userID int32, parentID int32) ([]Cell, error) {

	parentCell, err := c.Queries.GetCellById(ctx, parentID)
	if err != nil {
		return nil, err
	}

	if *parentCell.OwnerID != userID {
		return nil, fmt.Errorf("not authorized")
	}

	children, err := c.Queries.GetChildrenCellsByParentId(ctx, &parentID)
	if err != nil {
		return nil, err
	}

	cells := make([]Cell, len(children))

	for i, child := range children {
		cells[i] = Cell{
			Id:          child.ID,
			Color:       child.Color,
			Goal:        child.Goal,
			IsCompleted: child.IsCompleted,
		}
	}

	return cells, nil
}
