package item

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Item struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name      string
	Stock     int
	Price     int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
