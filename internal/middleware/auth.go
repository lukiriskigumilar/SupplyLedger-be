package middleware

import (
	"fmt"
	"os"
	"slices"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/lukiriskigumilar/SupplyLedger-be/internal/utils"
)

type CustomClaims struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

func AuthMiddleware(requiredRole ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		{

			const messageError = "Can't access this resource"

			//GET TOKEN FROM COOKIES
			tokenString := c.Cookies("access-token")
			if tokenString == "" {
				return utils.NewApiResponseError(
					c,
					"Unauthorized",
					401,
					fiber.Map{
						"reason": "missing cookies",
					},
				)
			}

			//GET AND CHECK env files
			key := os.Getenv("JWT_SECRET")
			if key == "" {
				return utils.NewApiResponseError(
					c,
					"Server error",
					500,
					fiber.Map{"reason": "JWT secret not configured"},
				)
			}

			claims := &CustomClaims{}

			//parsed token
			parsedToken, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
				if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method")
				}
				return []byte(key), nil
			})

			//check validity token
			if err != nil || !parsedToken.Valid {
				return utils.NewApiResponseError(
					c,
					messageError,
					401,
					fiber.Map{
						"reason": "invalid or expired token",
					},
				)
			}

			//CHECK VALIDATION STRUCT CLAIM
			if claims.UserID == "" || claims.Role == "" {
				return utils.NewApiResponseError(
					c,
					messageError,
					401,
					fiber.Map{"reason": "invalid claims structure"},
				)
			}

			//CHECK TYPE OF ROLE
			if len(requiredRole) > 0 && !slices.Contains(requiredRole, claims.Role) {
				return utils.NewApiResponseError(
					c,
					messageError,
					403,
					fiber.Map{"reason": "insufficient role permissions"},
				)
			}

			//ADD data to context
			c.Locals("user_id", claims.UserID)
			c.Locals("role", claims.Role)
			c.Locals("username", claims.Username)

			return c.Next()
		}

	}
}
