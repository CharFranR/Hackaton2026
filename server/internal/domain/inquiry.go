package domain

import (
	"time"

	"github.com/google/uuid"
)

type InquiryStatus int

const (
	InquiryPending InquiryStatus = iota
	InquiryRead
	InquiryReplied
	InquiryClosed
)

type Inquiry struct {
	ID         uuid.UUID
	UserID     uuid.UUID
	OfferingID uuid.UUID
	Message    string
	Status     InquiryStatus

	CreatedAt time.Time
}

// Builder

func NewInquiry(userID, offeringID uuid.UUID, message string) (*Inquiry, error) {
	if message == "" {
		return nil, ErrMessageRequired
	}

	return &Inquiry{
		ID:         uuid.New(),
		UserID:     userID,
		OfferingID: offeringID,
		Message:    message,
		Status:     InquiryPending,
		CreatedAt:  time.Now(),
	}, nil
}

// Set

func (i *Inquiry) MarkAsRead() {
	i.Status = InquiryRead
}

func (i *Inquiry) MarkAsReplied() {
	i.Status = InquiryReplied
}

func (i *Inquiry) Close() {
	i.Status = InquiryClosed
}

func (s InquiryStatus) String() string {
	switch s {
	case InquiryPending:
		return "pending"
	case InquiryRead:
		return "read"
	case InquiryReplied:
		return "replied"
	case InquiryClosed:
		return "closed"
	default:
		return "unknown"
	}
}
