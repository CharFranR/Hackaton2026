package dto

import (
	"time"

	"github.com/google/uuid"
)

type CompanyDTO struct {
	ID          uuid.UUID
	Name        string
	CategoryID  uuid.UUID
	OwnerID     uuid.UUID
	Address     string
	Description string
	PhoneNumber string
	Email       string
	Website     string
	Verified    bool

	CreatedAt time.Time
	UpdatedAt time.Time
}

type RegisterCompanyRequest struct {
	Name        string
	CategoryID  uuid.UUID
	Address     string
	Description string
	PhoneNumber string
	Email       string
	Website     string
}

type UpdateCompanyRequest struct {
	Name        *string
	Category    *string
	Address     *string
	Description *string
	PhoneNumber *string
	Email       *string
	Website     *string
}
