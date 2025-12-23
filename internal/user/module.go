package user

import "gorm.io/gorm"

type UserModule struct {
	UserRepo    UserRepository
	UserService UserService
	UserHandler *UserHandler
}

func InitUserModule(db *gorm.DB) *UserModule {
	repo := NewUserRepository(db)
	service := NewUserService(repo)
	handler := NewUserHandler(service)

	return &UserModule{
		UserRepo:    repo,
		UserService: service,
		UserHandler: handler,
	}
}
