package schemas

import (
	"gorm.io/gorm"
	"mandalart.com/types"
	"time"
)

type User struct {
	ID             int                  `gorm:"primaryKey;autoIncrement"`
	SocialID       string               `json:"socialId"`
	SocialProvider types.SocialProvider `json:"socialProvider"`
	CreatedAt      time.Time            `json:"createdAt"`
	UpdatedAt      time.Time            `json:"updatedAt"`
	DeletedAt      gorm.DeletedAt       `gorm:"index"`
	Sheets         []Sheet              `gorm:"foreignKey:OwnerID"`
}
