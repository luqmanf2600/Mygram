package models

type Comment struct {
	GormModel
	UserId  uint
	User    *User
	Photo   *Photo
	Content string `gorm:"not null" json:"content" form:"content" valid:"required~content is required"`
}
