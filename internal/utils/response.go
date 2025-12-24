package utils

import (
	"github.com/gofiber/fiber/v2"
)

type SuccessResponse struct {
	Success    bool        `json:"success"`
	Message    string      `json:"message"`
	StatusCode int         `json:"status_code"`
	Data       interface{} `json:"data,omitempty"`
	Pagination interface{} `json:"pagination,omitempty"`
}

type ErrorResponse struct {
	Success    bool        `json:"success"`
	Message    string      `json:"message"`
	StatusCode int         `json:"status_code"`
	Errors     interface{} `json:"errors,omitempty"`
}

func NewApiResponseSuccess(c *fiber.Ctx,
	message string, statusCode int, data interface{}, pagination interface{},
) error {
	response := SuccessResponse{
		Success:    true,
		Message:    message,
		StatusCode: statusCode,
		Data:       data,
	}
	if pagination != nil {
		response.Pagination = pagination
	}

	return c.Status(statusCode).JSON(response)
}

func NewApiResponseError(c *fiber.Ctx,
	message string, statusCode int, reason interface{},
) error {
	response := ErrorResponse{
		Success:    false,
		Message:    message,
		StatusCode: statusCode,
		Errors:     reason,
	}
	return c.Status(statusCode).JSON(response)
}
