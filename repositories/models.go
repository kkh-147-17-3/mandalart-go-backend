// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package repositories

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Cell struct {
	ID          int              `json:"id"`
	SheetID     int              `json:"sheetId"`
	Goal        *string          `json:"goal"`
	Color       *string          `json:"color"`
	Step        int              `json:"step"`
	Order       int              `json:"order"`
	ParentID    int              `json:"parentId"`
	IsCompleted bool             `json:"isCompleted"`
	CreatedAt   pgtype.Timestamp `json:"createdAt"`
	ModifiedAt  pgtype.Timestamp `json:"modifiedAt"`
	OwnerID     int              `json:"ownerId"`
}

type Sheet struct {
	ID         int              `json:"id"`
	OwnerID    int              `json:"ownerId"`
	Name       *string          `json:"name"`
	CreatedAt  pgtype.Timestamp `json:"createdAt"`
	ModifiedAt pgtype.Timestamp `json:"modifiedAt"`
}

type Todo struct {
	ID         int              `json:"id"`
	OwnerID    int              `json:"ownerId"`
	CellID     int              `json:"cellId"`
	Content    *string          `json:"content"`
	CreatedAt  pgtype.Timestamp `json:"createdAt"`
	ModifiedAt pgtype.Timestamp `json:"modifiedAt"`
}

type User struct {
	ID             int     `json:"id"`
	SocialID       *string `json:"socialId"`
	SocialProvider *string `json:"socialProvider"`
}
