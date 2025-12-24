package item

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lukiriskigumilar/SupplyLedger-be/internal/middleware"
)

func ItemsRoute(r fiber.Router, m *ItemModule) {
	items := r.Group("/item")
	items.Use(middleware.AuthMiddleware("user", "admin"))

	items.Post("", m.ItemHandler.CreateItemHandler)
	items.Get("", m.ItemHandler.GetAllHandler)
	items.Get("/search", m.ItemHandler.SearchItemByName)
	items.Get("/:id", m.ItemHandler.GetItemByIdHandler)
	items.Delete("/:id", m.ItemHandler.DeleteItemHandler)
}
