package dto

import (
	"time"

	domain "github.com/CharFranR/Hackaton2026/domain/entities"
	"github.com/google/uuid"
)

type OfferingDTO struct {
	ID          uuid.UUID
	CompanyID   uuid.UUID
	Type        domain.OfferingType
	Name        string
	Description string
	Price       float64
	ImageURL    string

	CreatedAt time.Time
	UpdatedAt time.Time
}

type CreateOfferingRequest struct {
	Type        domain.OfferingType
	Name        string
	Description string
	Price       float64
	ImageURL    string
}

type UpdateOfferingRequest struct {
	Type        *domain.OfferingType
	Name        *string
	Description *string
	Price       *float64
	ImageURL    *string
}
