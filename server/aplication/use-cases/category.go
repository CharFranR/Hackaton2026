package usecases

import (
	"context"

	"github.com/CharFranR/Hackaton2026/aplication/dto"
	domain "github.com/CharFranR/Hackaton2026/domain/entities"
	"github.com/CharFranR/Hackaton2026/domain/port/primary"
	"github.com/CharFranR/Hackaton2026/domain/port/secondary"
)

type CategoryUseCaseImpl struct {
	categoryRepo port.CategoryRepository
}

func NewCategoryUseCase(categoryRepo port.CategoryRepository) *CategoryUseCaseImpl {
	return &CategoryUseCaseImpl{
		categoryRepo: categoryRepo,
	}
}

func (uc *CategoryUseCaseImpl) GetAll(ctx context.Context) ([]*dto.CategoryDTO, error) {
	categories, err := uc.categoryRepo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	dtos := make([]*dto.CategoryDTO, len(categories))
	for i := range categories {
		dtos[i] = categoryToDTO(&categories[i])
	}

	return dtos, nil
}

var _ primary.CategoryUseCase = (*CategoryUseCaseImpl)(nil)

func categoryToDTO(category *domain.Category) *dto.CategoryDTO {
	return &dto.CategoryDTO{
		ID:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}
}
