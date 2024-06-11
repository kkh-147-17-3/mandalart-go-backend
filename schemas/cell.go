package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Cell struct {
	ID          int            `gorm:"primaryKey;autoIncrement"`
	SheetID     uint           `json:"sheetId"`
	Sheet       User           `json:"sheet"`
	Goal        string         `json:"goal"`
	Color       string         `json:"color"`
	Step        int            `json:"step"`
	Order       int            `json:"order"`
	ParentID    int            `json:"parentId"`
	Parent      *Cell          `json:"parent"`
	IsCompleted bool           `json:"isCompleted"`
	OwnerID     int            `json:"ownerId"`
	Owner       User           `json:"owner" gorm:"foreignKey:OwnerID"`
	Todos       []Todo         `json:"todos"`
	Children    []*Cell        `gorm:"foreignKey:ParentID" json:"children"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
