package port

import (
	"github.com/CharFranR/Hackaton2026/internal/domain"
	"github.com/google/uuid"
)

type UserRepository interface {
	FindByID(id uuid.UUID) (*domain.User, error)
	FindByEmail(email string) (*domain.User, error)
	ExistsByEmail(email string) (bool, error)
	Save(user *domain.User) error
	Update(user *domain.User) error
}

type CompanyRepository interface {
	FindByID(id uuid.UUID) (*domain.Company, error)
	FindByOwner(ownerID uuid.UUID) ([]domain.Company, error)
	Save(company *domain.Company) error
	Update(company *domain.Company) error
}

type OfferingRepository interface {
	FindByID(id uuid.UUID) (*domain.Offering, error)
	FindByCompany(companyID uuid.UUID) ([]domain.Offering, error)
	Save(offering *domain.Offering) error
	Update(offering *domain.Offering) error
	Delete(id uuid.UUID) error
}

type ReviewRepository interface {
	FindByCompany(companyID uuid.UUID) ([]domain.Review, error)
	FindByUser(userID uuid.UUID) ([]domain.Review, error)
	Save(review *domain.Review) error
}

type CategoryRepository interface {
	FindAll() ([]domain.Category, error)
	FindByID(id uuid.UUID) (*domain.Category, error)
}

type InquiryRepository interface {
	FindByID(id uuid.UUID) (*domain.Inquiry, error)
	FindByUser(userID uuid.UUID) ([]domain.Inquiry, error)
	Save(inquiry *domain.Inquiry) error
	Update(inquiry *domain.Inquiry) error
}
