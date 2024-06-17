package services

import (
	"context"
	repo "mandalart.com/repositories"
	"mandalart.com/utils"
)

type InsertTodosInput struct {
	OwnerID *int32  `json:"ownerId"`
	CellID  *int32  `json:"cellId"`
	Content *string `json:"content"`
}

type TodoService struct {
	q *repo.Queries
}

func NewTodoService() *TodoService {
	q := repo.New(utils.DBPool)
	return &TodoService{q}
}

func (s *TodoService) InsertTodos(ctx context.Context, userID int32, cellID int32, todos []InsertTodosInput) (*CellWithTodos, error) {
	q := repo.New(utils.DBPool)
	var args []repo.InsertTodosByCellIDParams
	for _, todo := range todos {
		args = append(args, repo.InsertTodosByCellIDParams{OwnerID: todo.OwnerID, CellID: todo.CellID, Content: todo.Content})
	}
	_, err := q.InsertTodosByCellID(ctx, args)
	if err != nil {
		return nil, err
	}

}
