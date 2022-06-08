package models

type Socialmedia struct {
	GormModel
	Name             string `gorm:"not null" json:"name" form:"name" valid:"required~name is required"`
	Social_media_url string `gorm:"not null" json:"social_media_url" form:"social_media_url" valid:"required~social_media_url is required"`
	UserId           uint
	User             *User
}
