package email

import (
	"context"
	"github.com/JIeeiroSst/go-app/internal/models/email"
)

type Mailer interface {
	Send(ctx context.Context, email *email.Email) error
}
