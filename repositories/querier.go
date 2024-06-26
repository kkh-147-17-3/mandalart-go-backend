// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package repositories

import (
	"context"
)

type Querier interface {
	CreateCell(ctx context.Context, arg CreateCellParams) (int, error)
	CreateSheet(ctx context.Context, arg CreateSheetParams) (int, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (int, error)
	DeleteTodosByCellID(ctx context.Context, cellID int) error
	GetCellById(ctx context.Context, id int) (Cell, error)
	GetChildrenCellsByParentId(ctx context.Context, parentID int) ([]Cell, error)
	GetLatestSheetByOwnerId(ctx context.Context, ownerID int) (Sheet, error)
	GetLatestSheetWithMainCellsByOwnerId(ctx context.Context, ownerID int) ([]GetLatestSheetWithMainCellsByOwnerIdRow, error)
	GetMainCellsBySheetId(ctx context.Context, sheetID int) ([]Cell, error)
	GetTodosByCellID(ctx context.Context, cellID int) ([]Todo, error)
	GetTodosByCellId(ctx context.Context, cellID int) ([]Todo, error)
	GetUserBySocialProviderInfo(ctx context.Context, arg GetUserBySocialProviderInfoParams) (User, error)
	InsertTodosByCellID(ctx context.Context, arg []InsertTodosByCellIDParams) (int64, error)
	UpdateCell(ctx context.Context, arg UpdateCellParams) error
}

var _ Querier = (*Queries)(nil)
