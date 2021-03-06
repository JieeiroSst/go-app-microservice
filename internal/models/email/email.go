package email

import (
	"github.com/google/uuid"
	"time"
)

type Email struct {
	EmailID     uuid.UUID `json:"emailId" db:"email_id" validate:"omitempty"`
	To          string  `json:"to" db:"to" validate:"required"`
	From        string    `json:"from,omitempty" db:"from" validate:"required,email"`
	Body        string    `json:"body" db:"body" validate:"required"`
	Subject     string    `json:"subject" db:"subject" validate:"required,lte=250"`
	ContentType string    `json:"contentType,omitempty" db:"content_type" validate:"lte=250"`
	CreatedAt   time.Time `json:"created_at,omitempty" db:"created_at"`
}

type EmailsList struct {
	TotalCount uint64   `json:"total_count"`
	TotalPages uint64   `json:"total_pages"`
	Page       uint64   `json:"page"`
	Size       uint64   `json:"size"`
	HasMore    bool     `json:"has_more"`
	Emails     []*Email `json:"emails"`
}
