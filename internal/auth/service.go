package auth

import (
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/lukiriskigumilar/SupplyLedger-be/internal/common"
	"github.com/lukiriskigumilar/SupplyLedger-be/internal/user"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	RegisterService(input RegisterRequestDTO) (*RegisterResponseDTO, error)
	LoginService(input LoginRequestDTO) (*LoginResult, error)
}

type authService struct {
	repo user.UserRepository
}

func NewAuthService(repo user.UserRepository) AuthService {
	return &authService{repo}
}

// REGISTER SERVICE
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

// LOGIN SERVICE
func (s *authService) LoginService(input LoginRequestDTO) (
	*LoginResult, error) {
	username := strings.ToLower(strings.ReplaceAll(input.Username, " ", ""))

	//find user
	user, err := s.repo.FindByUsername(username)
	if err != nil {
		return nil, &common.AppError{
			StatusCode: 400,
			Message:    "Login failed",
			Reason:     "invalid email or password",
		}
	}

	//compare password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		return nil, &common.AppError{
			StatusCode: 400,
			Message:    "Login failed",
			Reason:     "invalid email or password",
		}
	}

	createToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  user.ID,
		"role":     user.Role,
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	key := os.Getenv("JWT_SECRET")

	token, err := createToken.SignedString([]byte(key))
	if err != nil {
		return nil, &common.AppError{}
	}

	response := &LoginResult{
		userID:    user.ID.String(),
		Username:  user.Username,
		Role:      user.Role,
		Token:     token,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return response, nil

}
