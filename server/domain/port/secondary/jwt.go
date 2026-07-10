package port

import (
	domain "github.com/CharFranR/Hackaton2026/domain/entities"
	"github.com/google/uuid"
)

type JWTProvider interface {
	GenerateToken(userID uuid.UUID, role domain.RoleOptions) (string, error)
	ValidateToken(token string) (*JWTClaims, error)
}

type JWTClaims struct {
	UserID uuid.UUID
	Role   domain.RoleOptions
}
