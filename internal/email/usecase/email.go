package usecase

import (
	"context"
	"encoding/json"
	"github.com/JIeeiroSst/go-app/config"
	"github.com/JIeeiroSst/go-app/internal/email"
	emailModel"github.com/JIeeiroSst/go-app/internal/models/email"
	"github.com/JIeeiroSst/go-app/pkg/logger"
	"github.com/JIeeiroSst/go-app/utils"
	"github.com/google/uuid"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
	"github.com/pkg/errors"
)

type EmailUseCase struct {
	mailer     email.Mailer
	emailsRepo email.EmailsRepository
	logger     logger.Logger
	cfg        *config.Config
	publisher  email.EmailsPublisher
}

func NewEmailUseCase(emailsRepo email.EmailsRepository, logger logger.Logger, mailer email.Mailer, cfg *config.Config, publisher email.EmailsPublisher) *EmailUseCase {
	return &EmailUseCase{emailsRepo: emailsRepo, logger: logger, mailer: mailer, cfg: cfg, publisher: publisher}
}

func (e *EmailUseCase) SendEmail(ctx context.Context, deliveryBody []byte) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "EmailUseCase.SendEmail")
	defer span.Finish()

	mail := &emailModel.Email{}
	if err := json.Unmarshal(deliveryBody, mail); err != nil {
		return errors.Wrap(err, "json.Unmarshal")
	}

	mail.Body = utils.SanitizeString(mail.Body)

	mail.From = e.cfg.Smtp.User
	if err := utils.ValidateStruct(ctx, mail); err != nil {
		return errors.Wrap(err, "ValidateStruct")
	}

	if err := e.mailer.Send(ctx, mail); err != nil {
		return errors.Wrap(err, "mailer.Send")
	}

	createdEmail, err := e.emailsRepo.CreateEmail(*mail)
	if err != nil {
		return errors.Wrap(err, "emailsRepo.CreateEmail")
	}

	span.LogFields(log.String("emailID", createdEmail.EmailID.String()))
	e.logger.Infof("Success send email: %v", createdEmail.EmailID)
	return nil
}

func (e *EmailUseCase) PublishEmailToQueue(ctx context.Context, email *emailModel.Email) error {
	span, _ := opentracing.StartSpanFromContext(ctx, "EmailUseCase.PublishEmailToQueue")
	defer span.Finish()

	mailBytes, err := json.Marshal(email)
	if err != nil {
		return errors.Wrap(err, "json.Marshal")
	}

	return e.publisher.Publish(mailBytes, email.ContentType)
}

func (e *EmailUseCase) FindEmailById(ctx context.Context, emailID uuid.UUID) (*emailModel.Email, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "EmailUseCase.FindEmailById")
	defer span.Finish()
	email,err:=e.emailsRepo.EmailById(emailID)
	return &email, err
}

func (e *EmailUseCase) FindEmailsByReceiver(ctx context.Context, query *utils.PaginationQuery) (*emailModel.EmailsList, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "EmailUseCase.FindEmailsByReceiver")
	defer span.Finish()
	email,err:=e.emailsRepo.EmailAll(query)
	return &email,err
}