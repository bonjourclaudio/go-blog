package post

import (
	"github.com/jinzhu/gorm"
)

type Post struct {
	gorm.Model

	Title string `json:"title"`
	Content string `json:"content"`
	Tags []Tag `gorm:"many2many:post_tags" json:"-"`
	Likes int `json:"likes"`

	AuthorID uint `json:"author_id"`
}

type Tag struct {
	gorm.Model

	Content string `json:"content"`
}