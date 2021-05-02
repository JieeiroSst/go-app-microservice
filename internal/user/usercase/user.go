package usercase

import (
	"github.com/JIeeiroSst/go-app/config"
	userModel "github.com/JIeeiroSst/go-app/internal/models/user"
	"github.com/JIeeiroSst/go-app/internal/user"
	"github.com/JIeeiroSst/go-app/pkg/jwt"
	"github.com/JIeeiroSst/go-app/utils"
	"log"
)

type UserUseCase struct {
	userRepo user.UserRepository
	hash utils.Hash
	jwt jwt.TokenUser
	conf config.Config
}

func NewUserCase(userRepo user.UserRepository, hash utils.Hash, jwt jwt.TokenUser, conf config.Config) *UserUseCase{
	return &UserUseCase{
		userRepo:userRepo,
		hash:hash,
		jwt:jwt,
		conf:conf,
	}
}

func (u *UserUseCase) Login(user userModel.Users) string{
	check, _, hashPassword := u.userRepo.CheckAccount(user)
	if check == false {
		return "User does not exist"
	}
	if checkPass := u.hash.CheckPassowrd(user.Password, hashPassword); checkPass != nil {
		return "password entered incorrectly"
	}
	token, _ := u.jwt.GenerateToken()
	return token
}

func (u *UserUseCase) SignUp(user userModel.Users) string{
	check := u.userRepo.CheckAccountExists(user)
	if check == false {
		return "user already exists"
	}
	hashPassword, err := u.hash.HashPassword(user.Password)
	if err != nil {
		log.Println("error server", err)
	}
	account:=userModel.Users{
		Username:     user.Username,
		Password:     hashPassword,
	}
	_ = u.userRepo.CreateAccount(account)
	return user.Username+":"+hashPassword
}