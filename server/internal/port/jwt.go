package port

import (
	"github.com/CharFranR/Hackaton2026/internal/domain"
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
