package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Comment struct {
	Base
	Message string `gorm:"not null" json:"message" form:"message" valid:"required~message is required"`
	UserID  uint   `json:"user_id"`
	PhotoID uint   `json:"photo_id"`
}

func (c *Comment) BeforeCreate(tx *gorm.DB) (err error) {
	// validate input
	_, err = govalidator.ValidateStruct(c)
	return
}

func (c *Comment) BeforeUpdate(tx *gorm.DB) (err error) {
	// validate input
	_, err = govalidator.ValidateStruct(c)
	return
}
