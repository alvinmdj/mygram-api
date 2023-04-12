package models

import (
	"github.com/alvinmdj/mygram-api/helpers"
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	Base
	Username     string        `gorm:"not null;uniqueIndex" valid:"required~username is required"`
	Email        string        `gorm:"not null;uniqueIndex" valid:"required~email is required,email~invalid email format"`
	Password     string        `gorm:"not null" valid:"required~password is required,minstringlength(6)~password must have a minimum length of 6 characters"`
	Age          int           `gorm:"not null" valid:"required~age is required,range(8|99)~user must be at least 8 years old"`
	Photos       []Photo       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Comments     []Comment     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	SocialMedias []SocialMedia `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	// validate input
	if _, err = govalidator.ValidateStruct(u); err != nil {
		return
	}

	// hash password
	u.Password, err = helpers.HashPassword(u.Password)
	return
}
