package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lukiriskigumilar/SupplyLedger-be/internal/auth"
	"github.com/lukiriskigumilar/SupplyLedger-be/internal/config"
	httpx "github.com/lukiriskigumilar/SupplyLedger-be/internal/http"
	"github.com/lukiriskigumilar/SupplyLedger-be/internal/routes"
	"github.com/lukiriskigumilar/SupplyLedger-be/internal/user"
)

func main() {

	app := fiber.New(
		fiber.Config{
			ErrorHandler: httpx.ErrorHandler,
		},
	)

	//init db
	config.ConnectDatabase()
	db := config.DB

	// init modules
	userModule := user.InitUserModule(db)
	authModule := auth.InitAuthModule(userModule)

	api := app.Group("/api/v1")
	routes.GlobalRoutes(api, authModule)

	//ini routing

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"message": "success",
		})
	})

	app.Listen(":3000")
}
