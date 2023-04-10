package models

type SocialMedia struct {
	Base
	Name           string `gorm:"not null" json:"name" form:"name" valid:"required~name is required"`
	SocialMediaURL string `gorm:"not null" json:"social_media_url" form:"social_media_url" valid:"required~social media URL is required"`
	UserID         uint   `json:"user_id"`
}
