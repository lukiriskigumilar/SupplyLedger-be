package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lukiriskigumilar/SupplyLedger-be/internal/utils"
)

type AuthHandler struct {
	service AuthService
}

func NewAuthHandler(service AuthService) *AuthHandler {
	return &AuthHandler{service}
}

func (h *AuthHandler) RegisterHandler(c *fiber.Ctx) error {
	var req RegisterRequestDTO

	//validate request body
	if err := c.BodyParser(&req); err != nil {
		return utils.NewApiResponseError(c, "invalid request body", 400, fiber.Map{
			"reason": err.Error(),
		})

	}

	//validate request field
	if err := req.ValidateRegisterDTO(); err != nil {
		errors := utils.ParseValidationError(err)
		return utils.NewApiResponseError(
			c,
			"validation error",
			400,
			errors,
		)
	}

	//call service
	user, err := h.service.RegisterService(req)
	if err != nil {
		return err
	}

	//return value
	return utils.NewApiResponseSuccess(
		c,
		"User registered successfully",
		201,
		user,
		nil,
	)

}
