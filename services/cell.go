package services

import (
	"context"
	"fmt"
	repo "mandalart.com/repositories"
	"mandalart.com/utils"
)

type CellWithTodos struct {
	repo.Cell
	Todos []repo.Todo
}

type CellService struct {
	q repo.Querier
}

func NewCellService(q repo.Querier) *CellService {
	return &CellService{q}
}

func (c *CellService) GetChildrenCellsByParentID(ctx context.Context, userID int, parentID int) ([]Cell, error) {

	parentCell, err := c.q.GetCellById(ctx, parentID)
	if err != nil {
		return nil, err
	}

	if parentCell.OwnerID != userID {
		return nil, fmt.Errorf("not authorized")
	}

	children, err := c.q.GetChildrenCellsByParentId(ctx, parentID)
	if err != nil {
		return nil, err
	}

	cells := make([]Cell, len(children))

	for i, child := range children {
		cells[i] = Cell{
			ID:          child.ID,
			Color:       child.Color,
			Goal:        child.Goal,
			IsCompleted: child.IsCompleted,
		}
	}

	return cells, nil
}

func (c *CellService) GetCellWithTodosByID(ctx context.Context, cellID int) (*CellWithTodos, error) {
	q := repo.New(utils.DBPool)
	result, err := q.GetCellById(ctx, cellID)
	if err != nil {
		return nil, err
	}

	todos, err := q.GetTodosByCellID(ctx, result.ID)
	if err != nil {
		return nil, err
	}
	return &CellWithTodos{Cell: result, Todos: todos}, nil
}
