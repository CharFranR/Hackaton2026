package primary

import (
	"context"

	"github.com/CharFranR/Hackaton2026/aplication/dto"
	"github.com/google/uuid"
)

type ReviewUseCase interface {
	CreateReview(ctx context.Context, req dto.CreateReviewRequest) (*dto.ReviewDTO, error)
	GetByUser(ctx context.Context, UserId uuid.UUID) ([]*dto.ReviewDTO, error)
	GetByCompany(ctx context.Context, CompanyId uuid.UUID) ([]*dto.ReviewDTO, error)
}
