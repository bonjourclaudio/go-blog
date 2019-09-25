package auth

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model

	Firstname     	string
	Lastname		string
	Email    		string `gorm:"type:varchar(100);unique_index"`
	Password		string `json:"password"`
}