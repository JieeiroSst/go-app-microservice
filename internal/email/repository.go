package email

import (
	"github.com/JIeeiroSst/go-app/internal/models/email"
	"github.com/JIeeiroSst/go-app/utils"
	"github.com/google/uuid"
)

type EmailsRepository interface {
	CreateEmail(el email.Email) (create email.Email,err error)
	EmailById(id uuid.UUID) (email email.Email,err error)
	EmailsByTo(to string) (emails []email.Email,err error)
	EmailAll(query *utils.PaginationQuery)(email.EmailsList,error)
}
