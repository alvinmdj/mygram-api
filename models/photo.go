package models

type Photo struct {
	Base
	Title    string    `gorm:"not null" json:"title" form:"title" valid:"required~title is required"`
	Caption  string    `gorm:"not null" json:"caption" form:"caption" valid:"required~caption is required"`
	PhotoURL string    `gorm:"not null" json:"photo_url" form:"photo_url" valid:"required~photo URL is required"`
	UserID   uint      `json:"user_id"`
	Comments []Comment `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"comments"`
}
