package dto

import (
	"time"

	"github.com/google/uuid"
)

type ReviewDTO struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	CompanyID uuid.UUID
	Rating    int
	Comment   string
	CreatedAt time.Time
}

type CreateReviewRequest struct {
	CompanyID uuid.UUID
	Rating    int
	Comment   string
}
