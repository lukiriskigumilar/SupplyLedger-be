package auth

import (
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/lukiriskigumilar/SupplyLedger-be/internal/common"
	"github.com/lukiriskigumilar/SupplyLedger-be/internal/user"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	RegisterService(input RegisterRequestDTO) (*RegisterResponseDTO, error)
}

type authService struct {
	repo user.UserRepository
}

func NewAuthService(repo user.UserRepository) AuthService {
	return &authService{repo}
}

func (s *authService) RegisterService(input RegisterRequestDTO) (*RegisterResponseDTO, error) {

	//checkUsername
	username := strings.ToLower(strings.ReplaceAll(input.Username, " ", ""))

	//check if username already exist
	existing, _ := s.repo.FindByUsername(username)
	if existing.ID != uuid.Nil {
		return nil, &common.AppError{
			StatusCode: 409,
			Message:    "Create user failed",
			Reason:     "Username already in use",
		}
	}

	//hash password
	hashed, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	//create users
	const DefaultRole = "user"
	newUser := &user.User{
		ID:        uuid.New(),
		Username:  username,
		Role:      DefaultRole,
		Password:  string(hashed),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err = s.repo.Create(newUser)
	if err != nil {
		return nil, err
	}

	Response := &RegisterResponseDTO{
		Id:        newUser.ID.String(),
		Username:  newUser.Username,
		CreatedAt: newUser.CreatedAt,
	}

	return Response, nil
}
