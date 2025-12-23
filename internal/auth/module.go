package auth

import "github.com/lukiriskigumilar/SupplyLedger-be/internal/user"

type AuthModule struct {
	AuthHandler *AuthHandler
	AuthService AuthService
}

func InitAuthModule(userModule *user.UserModule) *AuthModule {
	service := NewAuthService(userModule.UserRepo)
	handler := NewAuthHandler(service)

	return &AuthModule{
		AuthService: service,
		AuthHandler: handler,
	}

}
