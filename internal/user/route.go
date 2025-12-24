package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lukiriskigumilar/SupplyLedger-be/internal/middleware"
)

func UserRoute(r fiber.Router, m *UserModule) {
	user := r.Group("/user")

	user.Get("/detail", middleware.AuthMiddleware("user", "admin"), m.UserHandler.GetDetailUserHandler)
}
