package primary

import (
	"context"

	"github.com/CharFranR/Hackaton2026/aplication/dto"
)

type CategoryUseCase interface {
	GetAll(ctx context.Context) ([]*dto.CategoryDTO, error)
}
