package usecases

import (
	"context"
	"time"

	"github.com/CharFranR/Hackaton2026/aplication/dto"
	"github.com/CharFranR/Hackaton2026/domain/port/primary"
	port "github.com/CharFranR/Hackaton2026/domain/port/secondary"
	"github.com/google/uuid"
)

type CachedCompanyUseCase struct {
	next  primary.CompanyUseCase
	cache port.Cache
}

func NewCachedCompanyUseCase(next primary.CompanyUseCase, cache port.Cache) *CachedCompanyUseCase {
	return &CachedCompanyUseCase{
		next:  next,
		cache: cache,
	}
}

func (uc *CachedCompanyUseCase) GetByID(ctx context.Context, id uuid.UUID) (*dto.CompanyDTO, error) {

	var company *dto.CompanyDTO

	_, err := uc.cache.Remember(
		ctx,
		"company:"+id.String(),
		time.Hour,
		&company,
		func() error {
			result, err := uc.next.GetByID(ctx, id)

			if err != nil {
				return err
			}

			company = result
			return nil
		},
	)

	return company, err
}

func (uc *CachedCompanyUseCase) GetByOwner(ctx context.Context, ownerID uuid.UUID) ([]*dto.CompanyDTO, error) {

	var company []*dto.CompanyDTO

	_, err := uc.cache.Remember(
		ctx,
		"company:byowner:"+ownerID.String(),
		time.Hour,
		&company,
		func() error {
			result, err := uc.next.GetByOwner(ctx, ownerID)

			if err != nil {
				return err
			}

			company = result
			return nil
		},
	)

	return company, err
}

func (uc *CachedCompanyUseCase) CreateCompany(ctx context.Context, req dto.RegisterCompanyRequest) (*dto.CompanyDTO, error) {
	return uc.next.CreateCompany(ctx, req)
}

func (uc *CachedCompanyUseCase) UpdateCompany(ctx context.Context, id uuid.UUID, req dto.UpdateCompanyRequest) error {
	return uc.next.UpdateCompany(ctx, id, req)
}

var _ primary.CompanyUseCase = (*CachedCompanyUseCase)(nil)
