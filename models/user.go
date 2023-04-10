package models

type User struct {
	Base
	Username     string        `gorm:"not null;uniqueIndex" json:"username" form:"username" valid:"required~username is required"`
	Email        string        `gorm:"not null;uniqueIndex" json:"email" form:"email" valid:"required~email is required,email~invalid email format"`
	Password     string        `gorm:"not null" json:"password" form:"password" valid:"required~password is required,minstringlength(6)~password must have a minimum length of 6 characters"`
	Age          int           `gorm:"not null" json:"age" form:"age" valid:"required~age is required,min=8~user must be at least 8 years old"`
	Photos       []Photo       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"photos"`
	Comments     []Comment     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"comments"`
	SocialMedias []SocialMedia `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"social_medias"`
}
