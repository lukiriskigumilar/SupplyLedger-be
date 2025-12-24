package item

import "gorm.io/gorm"

type ItemModule struct {
	ItemRepository ItemRepository
	ItemService    ItemService
	ItemHandler    *ItemHandler
}

func InitItemModule(db *gorm.DB) *ItemModule {
	repo := NewItemRepository(db)
	service := NewItemService(repo)
	handler := NewItemHandler(service)

	return &ItemModule{
		ItemRepository: repo,
		ItemService:    service,
		ItemHandler:    handler,
	}
}
