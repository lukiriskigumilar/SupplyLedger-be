package item

import (
	"errors"

	"gorm.io/gorm"
)

type ItemRepository interface {
	CreateItem(item *Item) error
	SearchByName(name string) (*Item, error)
}

type itemRepository struct {
	db *gorm.DB
}

func NewItemRepository(db *gorm.DB) ItemRepository {
	return &itemRepository{db}
}

func (r *itemRepository) CreateItem(item *Item) error {
	return r.db.Create(item).Error
}

func (r *itemRepository) SearchByName(name string) (*Item, error) {
	var item Item
	err := r.db.Where("name = ?", name).First(&item).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &item, nil
}
