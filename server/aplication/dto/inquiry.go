package dto

import (
	"time"

	domain "github.com/CharFranR/Hackaton2026/domain/entities"
	"github.com/google/uuid"
)

type InquiryDTO struct {
	ID         uuid.UUID
	UserID     uuid.UUID
	OfferingID uuid.UUID
	Message    string
	Status     domain.InquiryStatus

	CreatedAt time.Time
	UpdatedAt time.Time
}

type CreateInquiryRequest struct {
	OfferingID uuid.UUID
	Message    string
}

type UpdateInquiryRequest struct {
	Status *domain.InquiryStatus
}
