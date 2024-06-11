package schemas

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Sheet struct {
	ID        uuid.UUID      `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	OwnerID   int            `json:"ownerId"`
	Owner     User           `json:"owner"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
