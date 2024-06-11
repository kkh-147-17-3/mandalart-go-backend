package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	ID        int  `gorm:"primaryKey;autoIncrement"`
	OwnerID   uint `json:"ownerId"`
	Owner     User `json:"owner" gorm:"foreignKey:OwnerID"`
	CellID    int  `json:"cellId"`
	Cell      Cell `json:"cell"`
	content   string
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
