package usecases

import (
	"context"
	"time"

	"github.com/CharFranR/Hackaton2026/aplication/dto"
	"github.com/CharFranR/Hackaton2026/domain/port/primary"
	port "github.com/CharFranR/Hackaton2026/domain/port/secondary"
	"github.com/google/uuid"
)

type CachedReviewUseCase struct {
	next  primary.ReviewUseCase
	cache port.Cache
}

func NewCachedReviewUseCase(next primary.ReviewUseCase, cache port.Cache) *CachedReviewUseCase {
	return &CachedReviewUseCase{
		next:  next,
		cache: cache,
	}
}

func (uc *CachedReviewUseCase) CreateReview(ctx context.Context, req dto.CreateReviewRequest) (*dto.ReviewDTO, error) {
	return uc.next.CreateReview(ctx, req)
}

func (uc *CachedReviewUseCase) FindByUser(ctx context.Context, userID uuid.UUID) ([]*dto.ReviewDTO, error) {
	var reviews []*dto.ReviewDTO

	_, err := uc.cache.Remember(
		ctx,
		"reviews:byuser:"+userID.String(),
		time.Hour,
		&reviews,
		func() error {
			result, err := uc.next.FindByUser(ctx, userID)
			if err != nil {
				return err
			}

			reviews = result
			return nil
		},
	)

	return reviews, err
}

func (uc *CachedReviewUseCase) FindByCompany(ctx context.Context, companyID uuid.UUID) ([]*dto.ReviewDTO, error) {
	var reviews []*dto.ReviewDTO

	_, err := uc.cache.Remember(
		ctx,
		"reviews:bycompany:"+companyID.String(),
		time.Hour,
		&reviews,
		func() error {
			result, err := uc.next.FindByCompany(ctx, companyID)
			if err != nil {
				return err
			}

			reviews = result
			return nil
		},
	)

	return reviews, err
}

var _ primary.ReviewUseCase = (*CachedReviewUseCase)(nil)
