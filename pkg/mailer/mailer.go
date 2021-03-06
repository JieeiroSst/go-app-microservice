package mailer

import (
	"gopkg.in/gomail.v2"

	"github.com/JIeeiroSst/go-app/config"
)

// New Mail dialer
func NewMailDialer(cfg *config.Config) *gomail.Dialer {
	return gomail.NewDialer(cfg.Smtp.Host, cfg.Smtp.Port, cfg.Smtp.User, cfg.Smtp.Password)
}
