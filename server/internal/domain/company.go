package domain

import (
	"time"

	"github.com/google/uuid"
)

type Company struct {
	ID          uuid.UUID
	Name        string
	Category    []Category
	Owner       User
	Address     Address
	Description string
	PhoneNumber string
	Email       string
	Website     string
	Verified    bool

	CreatedAt time.Time
	UpdatedAt time.Time
}

// Builder

func NewCompany(owner User) *Company {
	now := time.Now()

	return &Company{
		ID:        uuid.New(),
		Owner:     owner,
		Verified:  false,
		CreatedAt: now,
		UpdatedAt: now,
	}
}

// Get

func (c Company) IsVerified() bool {
	return c.Verified
}

// Set

func (c *Company) Verify() {
	c.Verified = true
}

func (c *Company) Touch() {
	c.UpdatedAt = time.Now()
}

func (c *Company) ChangeOwner(new_owner User) {
	c.Owner = new_owner
}
