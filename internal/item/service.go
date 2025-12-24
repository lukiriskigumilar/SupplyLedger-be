package item

import (
	"math"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/lukiriskigumilar/SupplyLedger-be/internal/common"
)

type ItemService interface {
	CreateItem(input CreateItemRequestDTO) (*ItemResponseDTO, error)
	GetAll(page int, limit int) ([]ItemResponseDTO, common.Pagination, error)
	getItemById(id string) (*ItemResponseDTO, error)
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

// CREATE ITEM SERVICE
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

// CREATE GET ALL ITEMS SERVICE
func (s *itemService) GetAll(page int, limit int) ([]ItemResponseDTO, common.Pagination, error) {
	const errorMessage = "get data item failed"
	//guard and give default value
	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}

	//counting offset
	offset := (page - 1) * limit

	//call repository
	items, total, err := s.itemRepo.GetAllItem(
		common.PaginationQuery{
			Limit:  limit,
			Offset: offset,
		},
	)
	//check if error
	if err != nil {
		return nil, common.Pagination{}, &common.AppError{
			StatusCode: 500,
			Message:    errorMessage,
			Reason:     "internal server error",
		}
	}

	//mapping domain -> response dto
	itemResponses := make([]ItemResponseDTO, 0, len(items))
	for _, item := range items {
		itemResponses = append(itemResponses, ItemResponseDTO{
			ID:        item.ID.String(),
			Name:      item.Name,
			Stock:     item.Stock,
			Price:     item.Price,
			CreatedAt: item.CreatedAt,
			UpdatedAt: item.UpdatedAt,
		})
	}

	//count pagination metadata
	totalPages := 0
	if limit > 0 {
		totalPages = int(math.Ceil(float64(total) / float64(limit)))
	}
	hashNext := (page * limit) < int(total)

	pagination := common.Pagination{
		TotalData:   total,
		Limit:       limit,
		CurrentPage: page,
		TotalPages:  totalPages,
		HasPrev:     page > 1,
		HasNext:     hashNext,
	}

	//return final result
	return itemResponses, pagination, nil

}

func (s *itemService) getItemById(id string) (*ItemResponseDTO, error) {
	const messageError = "failed to get data"
	//call repository
	item, err := s.itemRepo.getItemById(id)
	if err != nil {
		return nil, &common.AppError{
			StatusCode: 500,
			Message:    messageError,
			Reason:     err.Error(),
		}
	}
	if item == nil {
		return nil, &common.AppError{
			StatusCode: 404,
			Message:    messageError,
			Reason:     "item not found",
		}
	}

	res := &ItemResponseDTO{
		ID:        item.ID.String(),
		Name:      item.Name,
		Stock:     item.Stock,
		Price:     item.Price,
		CreatedAt: item.CreatedAt,
		UpdatedAt: item.CreatedAt,
	}
	return res, nil
}
