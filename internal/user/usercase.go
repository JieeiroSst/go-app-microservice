package user

import (
	"github.com/JIeeiroSst/go-app/internal/models/user"
)

type UserUserCase interface {
	Login(user user.Users) string
	SignUp(user user.Users) string
}