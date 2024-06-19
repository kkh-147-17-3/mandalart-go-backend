package services

import (
	"context"
	repo "mandalart.com/repositories"
	"mandalart.com/utils"
)

type InsertTodosInput struct {
	OwnerID *int  `json:"ownerId"`
	CellID  *int  `json:"cellId"`
	Content *string `json:"content"`
}

type TodoService struct {
	q *repo.Queries
	cellService *CellService
}

func NewTodoService(q *repo.Queries, c *CellService) *TodoService {
	return &TodoService{q, c}
}

func (s *TodoService) InsertTodos(ctx context.Context, userID int, cellID int, todos []InsertTodosInput) (*CellWithTodos, error) {
	q := repo.New(utils.DBPool)
	var args []repo.InsertTodosByCellIDParams
	for _, todo := range todos {
		args = append(args, repo.InsertTodosByCellIDParams{OwnerID: *todo.OwnerID, CellID: *todo.CellID, Content: todo.Content})
	}
	_, err := q.InsertTodosByCellID(ctx, args)
	if err != nil {
		return nil, err
	}

	cell, err := s.cellService.GetCellWithTodosByID(ctx, cellID)

	if err != nil {
		return nil, err
	}

	return cell, err
}
