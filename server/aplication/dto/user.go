package dto

import (
	"time"

	domain "github.com/CharFranR/Hackaton2026/domain/entities"
	"github.com/google/uuid"
)

type UserDTO struct {
	ID          uuid.UUID
	Email       string
	FirstName   string
	LastName    string
	Address     string
	PhoneNumber string
	Role        domain.RoleOptions

	CreatedAt time.Time
	UpdatedAt time.Time
}

type RegisterUserRequest struct {
	Email           string
	FirstName       string
	LastName        string
	Address         string
	PhoneNumber     string
	Role            domain.RoleOptions
	Password        string
	ConfirmPassword string
}

type LoginRequest struct {
	Email    string
	Password string
}

type UpdateUserRequest struct {
	Email       *string
	FirstName   *string
	LastName    *string
	Address     *string
	PhoneNumber *string
}

type LoginResponse struct {
	AccessToken string
	ExpiresIn   int64
	User        UserDTO
}
