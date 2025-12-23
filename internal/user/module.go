package user

import "gorm.io/gorm"

type UserModule struct {
	UserRepo UserRepository
}

func InitUserModule(db *gorm.DB) *UserModule {
	repo := NewUserRepository(db)

	return &UserModule{
		UserRepo: repo,
	}
}
