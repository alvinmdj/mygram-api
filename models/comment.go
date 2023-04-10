package models

type Comment struct {
	Base
	Message string `gorm:"not null" json:"message" form:"message" valid:"required~message is required"`
	UserID  uint   `json:"user_id"`
	PhotoID uint   `json:"photo_id"`
}
