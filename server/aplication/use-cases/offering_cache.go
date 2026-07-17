package usecases

import (
	"context"
	"time"

	"github.com/CharFranR/Hackaton2026/aplication/dto"
	"github.com/CharFranR/Hackaton2026/domain/port/primary"
	port "github.com/CharFranR/Hackaton2026/domain/port/secondary"
	"github.com/google/uuid"
)

type CachedOfferingUseCase struct {
	next  primary.OfferingUseCase
	cache port.Cache
}

func NewCacheCategoryUseCase(next primary.OfferingUseCase, cache port.Cache) *CachedOfferingUseCase {
	return &CachedOfferingUseCase{
		next:  next,
		cache: cache,
	}
}

func (uc CachedOfferingUseCase) GetByID(ctx context.Context, id uuid.UUID) (*dto.OfferingDTO, error) {
	var offering *dto.OfferingDTO

	err := uc.cache.Remember(
		ctx,
		"offering:byid",
		time.Hour,
		&offering,
		func() error {
			result, err := uc.next.GetByID(ctx, id)
			if err != nil {
				return err
			}
			offering = result
			return nil
		},
	)

	return offering, err
}

func (uc CachedOfferingUseCase) GetByCompany(ctx context.Context, companyID uuid.UUID) ([]*dto.OfferingDTO, error) {
	var offering []*dto.OfferingDTO

	err := uc.cache.Remember(
		ctx,
		"offering:byid",
		time.Hour,
		&offering,
		func() error {
			result, err := uc.next.GetByCompany(ctx, companyID)
			if err != nil {
				return err
			}
			offering = result
			return nil
		},
	)

	return offering, err
}
