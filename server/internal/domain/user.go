package domain

import (
	"time"

	"github.com/google/uuid"

	"golang.org/x/crypto/bcrypt"
)

type RoleOptions int

const (
	RolePending RoleOptions = iota
	RoleMIPYME
	RoleProvider
	RoleAdmin
)

type User struct {
	ID         uuid.UUID
	FirstName  string
	LastName   string
	Role       RoleOptions
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Department string

	Email       string
	PhoneNumber string

	PasswordHash string
}

// Builder
func NewUser() *User {
	now := time.Now()

	return &User{
		ID:        uuid.New(),
		Role:      RolePending,
		CreatedAt: now,
		UpdatedAt: now,
	}
}

// Get

func (u User) FullName() string {
	return u.FirstName + " " + u.LastName
}

func (u User) IsAdmin() bool {
	return u.Role == RoleAdmin
}

func (u User) HasRole(role RoleOptions) bool {
	return u.Role == role
}

func (u User) CheckPassword(password string) error {

	return bcrypt.CompareHashAndPassword(
		[]byte(u.PasswordHash),
		[]byte(password),
	)

}

func (r RoleOptions) String() string {
	switch r {
	case RolePending:
		return "pending"
	case RoleMIPYME:
		return "mipyme"
	case RoleProvider:
		return "provider"
	case RoleAdmin:
		return "admin"
	default:
		return "unknown"
	}
}

// Set
func (u *User) SetPassword(password string) error {

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return err
	}

	u.PasswordHash = string(hashedPassword)
	return nil

}

func (u *User) Touch() {
	u.UpdatedAt = time.Now()
}
