package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type SocialMedia struct {
	Base
	Name           string `gorm:"not null" json:"name" form:"name" valid:"required~name is required"`
	SocialMediaURL string `gorm:"not null" json:"social_media_url" form:"social_media_url" valid:"required~social media URL is required"`
	UserID         uint   `json:"user_id"`
}

func (s *SocialMedia) BeforeCreate(tx *gorm.DB) (err error) {
	// validate input
	_, err = govalidator.ValidateStruct(s)
	return
}

func (s *SocialMedia) BeforeUpdate(tx *gorm.DB) (err error) {
	// validate input
	_, err = govalidator.ValidateStruct(s)
	return
}
