package primary

import (
	"context"

	"github.com/google/uuid"

	"github.com/CharFranR/Hackaton2026/aplication/dto"
)

type OfferingUseCase interface {
	CreateOffering(ctx context.Context, req dto.CreateOfferingRequest) (*dto.OfferingDTO, error)
	GetByID(ctx context.Context, id uuid.UUID) (*dto.OfferingDTO, error)
	GetByCompany(ctx context.Context, CompanyId uuid.UUID) ([]*dto.OfferingDTO, error)
	UpdateOffering(ctx context.Context, id uuid.UUID, req dto.UpdateOfferingRequest) error
}
