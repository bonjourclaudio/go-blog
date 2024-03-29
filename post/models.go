package post

import (
	"github.com/jinzhu/gorm"
)

type Post struct {
	gorm.Model

	Title string `json:"title"`
	Content string `json:"content"`
	Tags []Tag `gorm:"many2many:post_tags" json:"-"`
	Likes []Like `gorm:"many2many:post_likes" json:"-"`

	UserID uint `json:"user_id"`
}

type Tag struct {
	gorm.Model

	Content string `json:"content"`
}

type Like struct {
	gorm.Model

	UserID uint `json:"user_id"`
}