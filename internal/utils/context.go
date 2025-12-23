package utils

import "github.com/gofiber/fiber/v2"

func GetUserIDfromContext(c *fiber.Ctx) string {
	userID, ok := c.Locals("user_id").(string)
	if !ok {
		return ""
	}
	return userID
}

func GetUserRoleFromContext(c *fiber.Ctx) string {
	role, ok := c.Locals("role").(string)

	if !ok {
		return ""
	}
	return role
}

func GetUserUsernameFromContext(c *fiber.Ctx) string {
	username, ok := c.Locals("username").(string)
	if !ok {
		return ""
	}
	return username
}
