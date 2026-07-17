package usecases

import (
	"context"
	"time"

	"github.com/CharFranR/Hackaton2026/aplication/dto"
	"github.com/CharFranR/Hackaton2026/domain/port/primary"
	port "github.com/CharFranR/Hackaton2026/domain/port/secondary"
)

type CachedCategoryUseCase struct {
	next  primary.CategoryUseCase
	cache port.Cache
}

func NewCachedCategoryUseCase(next primary.CategoryUseCase, cache port.Cache) *CachedCategoryUseCase {
	return &CachedCategoryUseCase{
		next:  next,
		cache: cache,
	}
}

func (uc *CachedCategoryUseCase) GetAll(ctx context.Context) ([]*dto.CategoryDTO, error) {

	var categories []*dto.CategoryDTO

	err := uc.cache.Remember(
		ctx,
		"categories:all",
		time.Hour,
		&categories,
		func() error {
			result, err := uc.next.GetAll(ctx)
			if err != nil {
				return err
			}

			categories = result
			return nil
		},
	)

	return categories, err
}

var _ primary.CategoryUseCase = (*CachedCategoryUseCase)(nil)
