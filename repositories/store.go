package repositories

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type Store struct {
	Queries
	DB *pgx.Conn
}

func (s *Store) Begin(ctx context.Context) (pgx.Tx, error){
	return s.DB.Begin(ctx)
}