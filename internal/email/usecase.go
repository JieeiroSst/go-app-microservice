package email

import (
	"context"
	"github.com/JIeeiroSst/go-app/internal/models/email"
	"github.com/JIeeiroSst/go-app/utils"
	"github.com/google/uuid"
)

type EmailsUseCase interface {
	SendEmail(ctx context.Context, deliveryBody []byte) error
	PublishEmailToQueue(ctx context.Context, email *email.Email) error
	FindEmailById(ctx context.Context, emailID uuid.UUID) (*email.Email, error)
	FindEmailsByReceiver(ctx context.Context, query *utils.PaginationQuery) (*email.EmailsList, error)
}
