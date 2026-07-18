package usecases

import (
	"context"
	"time"

	"github.com/CharFranR/Hackaton2026/aplication/dto"
	"github.com/CharFranR/Hackaton2026/domain/port/primary"
	port "github.com/CharFranR/Hackaton2026/domain/port/secondary"
	"github.com/google/uuid"
)

type CachedInquiryUseCase struct {
	next  primary.InquiryUseCase
	cache port.Cache
}

func NewCachedInquiryUseCase(next primary.InquiryUseCase, cache port.Cache) *CachedInquiryUseCase {
	return &CachedInquiryUseCase{
		next:  next,
		cache: cache,
	}
}

func (uc *CachedInquiryUseCase) CreateInquiry(ctx context.Context, req dto.CreateInquiryRequest) (*dto.InquiryDTO, error) {
	return uc.next.CreateInquiry(ctx, req)
}

func (uc *CachedInquiryUseCase) GetByID(ctx context.Context, id uuid.UUID) (*dto.InquiryDTO, error) {
	var inquiry *dto.InquiryDTO

	err := uc.cache.Remember(
		ctx,
		"inquiry:"+id.String(),
		time.Hour,
		&inquiry,
		func() error {
			result, err := uc.next.GetByID(ctx, id)
			if err != nil {
				return err
			}

			inquiry = result
			return nil
		},
	)

	return inquiry, err
}

func (uc *CachedInquiryUseCase) GetByUser(ctx context.Context, userID uuid.UUID) ([]*dto.InquiryDTO, error) {
	var inquiries []*dto.InquiryDTO

	err := uc.cache.Remember(
		ctx,
		"inquiries:byuser:"+userID.String(),
		time.Hour,
		&inquiries,
		func() error {
			result, err := uc.next.GetByUser(ctx, userID)
			if err != nil {
				return err
			}

			inquiries = result
			return nil
		},
	)

	return inquiries, err
}

func (uc *CachedInquiryUseCase) UpdateInquiry(ctx context.Context, id uuid.UUID, req dto.UpdateInquiryRequest) error {
	return uc.next.UpdateInquiry(ctx, id, req)
}

var _ primary.InquiryUseCase = (*CachedInquiryUseCase)(nil)
