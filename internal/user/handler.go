package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lukiriskigumilar/SupplyLedger-be/internal/utils"
)

type UserHandler struct {
	userService UserService
}

func NewUserHandler(s UserService) *UserHandler {
	return &UserHandler{s}
}
func (h *UserHandler) GetDetailUserHandler(c *fiber.Ctx) error {

	//Get context
	userId := utils.GetUserIDfromContext(c)

	//call service
	result, err := h.userService.GetDetailUser(userId)

	if err != nil {
		return err
	}
	return utils.NewApiResponseSuccess(c, "Get data successfully", 200, result, nil)
}
