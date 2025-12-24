package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lukiriskigumilar/SupplyLedger-be/internal/auth"
	"github.com/lukiriskigumilar/SupplyLedger-be/internal/item"
	"github.com/lukiriskigumilar/SupplyLedger-be/internal/user"
)

func GlobalRoutes(api fiber.Router,
	authModule *auth.AuthModule,
	userModule *user.UserModule,
	itemModule *item.ItemModule,
) {
	auth.AuthRoute(api, authModule)
	user.UserRoute(api, userModule)
	item.ItemsRoute(api, itemModule)
}
