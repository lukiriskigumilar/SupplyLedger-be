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

func (h *AuthHandler) LoginHandler(c *fiber.Ctx) error {
	var req LoginRequestDTO
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
	if err := req.ValidateRegisterDTO(); err != nil {
		errors := utils.ParseValidationError(err)
		return utils.NewApiResponseError(
			c,
			"Validation error",
			400,
			errors,
		)
	}

	// call service
	result, err := h.service.LoginService(req)
	if err != nil {
		return err
	}

	// SET COOKIES
	c.Cookie(&fiber.Cookie{
		Name:     "access-token",
		Value:    result.Token,
		Path:     "/",
		HTTPOnly: true,
		Secure:   true,
		SameSite: fiber.CookieSameSiteLaxMode,
	})

	//RESPONSE
	response := &LoginResponseDTO{
		Username:  result.Username,
		Role:      result.Role,
		CreatedAt: result.CreatedAt,
	}

	return utils.NewApiResponseSuccess(
		c, "Login Successfully", 200, response, nil,
	)
}
