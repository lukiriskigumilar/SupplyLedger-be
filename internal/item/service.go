package item

import (
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/lukiriskigumilar/SupplyLedger-be/internal/common"
)

type ItemService interface {
	CreateItem(input CreateItemRequestDTO) (*ItemResponseDTO, error)
}

type itemService struct {
	itemRepo ItemRepository
}

func NewItemService(repo ItemRepository) ItemService {
	return &itemService{repo}
}

func normalizeName(name string) string {
	return strings.ToLower(strings.TrimSpace(name))
}

func (s *itemService) CreateItem(input CreateItemRequestDTO) (*ItemResponseDTO, error) {
	itemName := normalizeName(input.Name)
	const messageError = "failed to create item"

	//Check if name item already exists
	existing, err := s.itemRepo.SearchByName(itemName)
	if err != nil {
		return nil, err
	}
	if existing != nil {
		return nil, &common.AppError{
			StatusCode: 409,
			Message:    messageError,
			Reason:     "Item Name already in use",
		}
	}

	//validate price must be greater than 0
	if input.Price <= 0 {
		return nil, &common.AppError{
			StatusCode: 400,
			Message:    messageError,
			Reason:     "price must be greater than 0",
		}
	}

	//Create item model with generated UUID for uniqueness
	newItem := &Item{
		ID:        uuid.New(),
		Name:      itemName,
		Stock:     input.Stock,
		Price:     input.Price,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	//try save model new item
	fail := s.itemRepo.CreateItem(newItem)
	if fail != nil {
		return nil, &common.AppError{
			StatusCode: 500,
			Message:    messageError,
			Reason:     fail.Error(),
		}
	}

	//create dto response
	response := &ItemResponseDTO{
		ID:        newItem.ID.String(),
		Name:      newItem.Name,
		Stock:     newItem.Stock,
		Price:     newItem.Price,
		CreatedAt: newItem.CreatedAt,
	}

	return response, nil

}
