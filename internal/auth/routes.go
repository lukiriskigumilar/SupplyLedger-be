package auth

import (
	"github.com/gofiber/fiber/v2"
)

func AuthRoute(r fiber.Router, m *AuthModule) {
	authRoutes := r.Group("/auth")
	authRoutes.Post("/register", m.AuthHandler.RegisterHandler)

}
