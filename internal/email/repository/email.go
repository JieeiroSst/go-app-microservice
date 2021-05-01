package repository

import (
	"github.com/JIeeiroSst/go-app/internal/models/email"
	"github.com/JIeeiroSst/go-app/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type EmailRepository struct {
	db *gorm.DB
}

func NewEmailRepository(db *gorm.DB) *EmailRepository {
	return &EmailRepository{db:db}
}

func (e *EmailRepository) CreateEmail(el email.Email) (create email.Email,err error) {
	create = email.Email{
		EmailID:     uuid.UUID{},
		To:          el.To,
		From:        el.From,
		Body:        el.Body,
		Subject:     el.Subject,
		ContentType: el.ContentType,
		CreatedAt:   time.Now(),
	}
	err = e.db.Create(&create).Error
	return
}

func (e *EmailRepository) EmailById(id uuid.UUID) (email email.Email,err error){
	err = e.db.Where("id = ?",id).Find(&email).Error
	return
}

func (e *EmailRepository) EmailsByTo(to string) (emails []email.Email,err error){
	err = e.db.Where("to ilike ?",to).Order("created_at desc").Find(emails).Error
	return
}

func (e *EmailRepository) EmailAll(query *utils.PaginationQuery)(email.EmailsList,error){
	var count *int64
	e.db.Count(count).Find(&email.Email{})
	var emails []*email.Email
	e.db.Find(emails)
	var list= email.EmailsList{
		TotalCount: uint64(*count),
		TotalPages: utils.GetTotalPages(uint64(*count),query.GetSize()),
		Page:       query.GetPage(),
		Size:       query.GetSize(),
		HasMore:    utils.GetHasMore(query.GetPage(), uint64(*count), query.GetSize()),
		Emails:     emails,
	}
	return list,nil
}