package user

import "github.com/lukiriskigumilar/SupplyLedger-be/internal/common"

type UserService interface {
	GetDetailUser(id string) (*ResponseGetUserDTO, error)
}

type userService struct {
	userRepo UserRepository
}

func NewUserService(repo UserRepository) UserService {
	return &userService{repo}
}

// GET DETAIL USER SERVICE
func (s *userService) GetDetailUser(id string) (*ResponseGetUserDTO, error) {
	const messageError = "failed to get data"

	//Call repository
	user, err := s.userRepo.FindById(id)
	if user == nil {
		return nil, &common.AppError{
			StatusCode: 404,
			Message:    messageError,
			Reason:     "user not found",
		}
	}
	if err != nil {
		return nil, &common.AppError{
			StatusCode: 500,
			Message:    messageError,
			Reason:     err.Error(),
		}
	}
	response := &ResponseGetUserDTO{
		Username:  user.Username,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
	}
	return response, nil
}
