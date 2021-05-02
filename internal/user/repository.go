package user

import (
	"github.com/JIeeiroSst/go-app/internal/models/user"
)

type UserRepository interface {
	CheckAccount(u user.Users) (bool, int, string)
	CheckAccountExists(u user.Users) bool
	CreateAccount(user user.Users) error
}