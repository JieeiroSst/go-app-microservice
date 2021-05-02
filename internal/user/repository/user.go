package repository

import (
	"github.com/JIeeiroSst/go-app/internal/models/user"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository{
	return &UserRepository{db:db}
}

func (d *UserRepository)CheckAccount(u user.Users) (bool, int, string) {
	var accounts []user.Users
	d.db.Find(&accounts)
	for _,account:=range accounts{
		if account.Username==u.Username{
			return true, account.Id, account.Password
		}
	}
	return false, 0, ""
}

func (d *UserRepository) CheckAccountExists(u user.Users) bool {
	var account [] user.Users
	d.db.Find(&account)
	for _,item:=range account{
		if item.Username == u.Username {
			return false
		}
	}
	return true
}

func (d *UserRepository) CreateAccount(user user.Users) error {
	return d.db.Create(&user).Error
}