package auth

import (
	"errors"
	"regexp"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// DTO REGISTER
type RegisterRequestDTO struct {
	Username        string `json:"username"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

func (r RegisterRequestDTO) ValidateRegisterDTO() error {
	return validation.ValidateStruct(
		&r,
		validation.Field(&r.Username,
			validation.Required.Error("Username is required"),
			validation.Match(regexp.MustCompile(`^[^\s]+$`)).Error("username must not contain space"),
			validation.Length(4, 100).Error(`Username must be between 4 and 100 characters`),
		),
		validation.Field(&r.Password,
			validation.Required.Error("Password is required"),
			validation.Length(6, 0).Error("Password must contain at least 6 characters"),
		),
		validation.Field(&r.ConfirmPassword,
			validation.Required.Error("confirm_password is required"),
			validation.By(
				func(value interface{}) error {
					if value.(string) != r.Password {
						return errors.New("password do not match")
					}
					return nil
				},
			),
		),
	)
}

type RegisterResponseDTO struct {
	Id        string    `json:"id"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
}

//DTO LOGIN

type LoginRequestDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (r LoginRequestDTO) ValidateRegisterDTO() error {
	return validation.ValidateStruct(
		&r,
		validation.Field(&r.Username,
			validation.Required.Error("Username is required"),
			validation.Match(regexp.MustCompile(`^[^\s]+$`)).Error("username must not contain space"),
			validation.Length(4, 100).Error(`Username must be between 4 and 100 characters`),
		),
		validation.Field(&r.Password,
			validation.Required.Error("Password is required"),
			validation.Length(6, 0).Error("Password must contain at least 6 characters"),
		),
	)
}

type LoginResponseDTO struct {
	Username  string    `json:"username"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
}

type LoginResult struct {
	userID    string
	Username  string
	Role      string
	Token     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
