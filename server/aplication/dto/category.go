package dto

import "github.com/google/uuid"

type CategoryDTO struct {
	ID          uuid.UUID
	Name        string
	Description string
}
