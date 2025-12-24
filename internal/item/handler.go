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

// Get all item handler
func (h *ItemHandler) GetAllHandler(c *fiber.Ctx) error {

	//Get query param
	page := c.QueryInt("page", 0)
	limit := c.QueryInt("limit", 0)

	items, pagination, err := h.itemService.GetAll(page, limit)
	if err != nil {
		return err
	}
	return utils.NewApiResponseSuccess(
		c,
		"get items successfully",
		200,
		items,
		pagination,
	)

}

// Get item by id handler
func (h *ItemHandler) GetItemByIdHandler(c *fiber.Ctx) error {

	//get id from params
	item_id := c.Params("id")

	//call service
	result, err := h.itemService.GetItemById(item_id)
	if err != nil {
		return err
	}

	//send result
	return utils.NewApiResponseSuccess(c, "get item successfully", 200, result, nil)

}
