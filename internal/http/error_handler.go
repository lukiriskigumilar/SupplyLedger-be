package http

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/lukiriskigumilar/SupplyLedger-be/internal/common"
	"github.com/lukiriskigumilar/SupplyLedger-be/internal/utils"
)

func ErrorHandler(c *fiber.Ctx, err error) error {
	//domain error
	var appErr *common.AppError
	if errors.As(err, &appErr) {
		return utils.NewApiResponseError(
			c,
			appErr.Message,
			appErr.StatusCode,
			fiber.Map{
				"reason": appErr.Reason,
			},
		)
	}
	return utils.NewApiResponseError(
		c,
		"internal server error",
		500,
		nil,
	)
}
