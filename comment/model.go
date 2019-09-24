package comment

import (
	"github.com/jinzhu/gorm"
)

type Comment struct {
	gorm.Model

	Content string `json:"content"`

	AuthorID uint `json:"author_id"`
	PostID uint `json:"post_id"`
}