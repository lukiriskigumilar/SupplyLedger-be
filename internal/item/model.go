package item

import (
	"time"

	"github.com/google/uuid"
)

type Item struct {
	ID        uuid.UUID
	Name      string
	Stock     int
	Price     int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
