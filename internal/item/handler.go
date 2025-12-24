package item

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lukiriskigumilar/SupplyLedger-be/internal/utils"
)

type ItemHandler struct {
	itemService ItemService
}

func NewItemHandler(s ItemService) *ItemHandler {
	return &ItemHandler{s}
}

// Crate item handler
func (h *ItemHandler) CreateItemHandler(c *fiber.Ctx) error {
	var req CreateItemRequestDTO

	//validate request body
	if err := c.BodyParser(&req); err != nil {
		return utils.NewApiResponseError(
			c,
			"invalid request",
			400,
			fiber.Map{
				"reason": err.Error(),
			},
		)
	}

	// validate request field
	if err := req.ValidateCreateItemRequestDTO(); err != nil {
		errors := utils.ParseValidationError(err)
		return utils.NewApiResponseError(
			c,
			"Validation error",
			400,
			errors,
		)
	}

	//call service
	result, err := h.itemService.CreateItem(req)
	if err != nil {
		return err
	}

	return utils.NewApiResponseSuccess(
		c,
		"create Items Successfully", 201, result, nil,
	)
}
