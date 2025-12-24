package item

import (
	"errors"

	"github.com/lukiriskigumilar/SupplyLedger-be/internal/common"
	"gorm.io/gorm"
)

type ItemRepository interface {
	CreateItem(item *Item) error
	SearchByName(name string) (*Item, error)
	GetAllItem(query common.PaginationQuery) ([]Item, int, error)
	getItemById(id string) (*Item, error)
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

func (r *itemRepository) getItemById(id string) (*Item, error) {
	var item Item
	err := r.db.Where("id=?", id).First(&item).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *itemRepository) GetAllItem(query common.PaginationQuery) ([]Item, int, error) {
	var items []Item
	var total int64

	if err := r.db.Model(&Item{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := r.db.Limit(query.Limit).Offset(query.Offset).Find(&items).Error; err != nil {
		return nil, 0, err
	}

	return items, int(total), nil

}
