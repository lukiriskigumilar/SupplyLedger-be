package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lukiriskigumilar/SupplyLedger-be/internal/auth"
)

func GlobalRoutes(api fiber.Router,
	authModule *auth.AuthModule,
) {
	auth.AuthRoute(api, authModule)
}
