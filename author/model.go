package author

import (
	"github.com/jinzhu/gorm"
)

type Author struct {
	gorm.Model
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
	Email string `json:"email"`
	AvatarUrl string `json:"avatarUrl"`
}