package models

type SocialMediaGetOutput struct {
	Base
	Name           string             `json:"name" form:"name"`
	SocialMediaURL string             `json:"social_media_url" form:"social_media_url"`
	User           UserRegisterOutput `json:"user"`
}

type SocialMediaCreateInput struct {
	Name           string `json:"name" form:"name" valid:"required~name is required"`
	SocialMediaURL string `json:"social_media_url" form:"social_media_url" valid:"required~social media URL is required"`
	UserID         uint   `valid:"required~user ID is required"`
}

type SocialMediaCreateOutput struct {
	Base
	Name           string `json:"name" form:"name"`
	SocialMediaURL string `json:"social_media_url" form:"social_media_url"`
	UserID         uint   `json:"user_id"`
}