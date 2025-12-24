package item

import (
	"regexp"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// dto create item
type CreateItemRequestDTO struct {
	Name  string `json:"name_item"`
	Stock int    `json:"stock"`
	Price int    `json:"price"`
}

func (c CreateItemRequestDTO) ValidateCreateItemRequestDTO() error {
	return validation.ValidateStruct(
		&c,
		validation.Field(&c.Name,
			validation.Required.Error("name item is required"),
			validation.Length(3, 30).Error("name item must be between 3 and 30 characters"),
			validation.Match(
				regexp.MustCompile(`^[A-Za-z]+( [A-Za-z]+)*$`),
			).Error("name item must contain only letters and single spaces between words"),
		),
		validation.Field(&c.Stock,
			validation.Required.Error("stock is required"),
		),
		validation.Field(&c.Price,
			validation.Required.Error("price is required"),
			validation.Min(0).Error("price must be greater than or equal to 0"),
		),
	)
}

// DTo response item
type ItemResponseDTO struct {
	ID        string
	Name      string    `json:"name"`
	Stock     int       `json:"stock"`
	Price     int       `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
