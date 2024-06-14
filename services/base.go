package services

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	repo "mandalart.com/repositories"
)

type BaseService struct {
	Queries *repo.Queries
}

func NewBaseService(ctx context.Context) (*BaseService, error){
	conn, ok := ctx.Value("db").(*pgxpool.Pool)
	if !ok {
		return nil, fmt.Errorf("database is not initialized")
	}
	Queries := repo.New(conn)
	return &BaseService{Queries: Queries}, nil
}