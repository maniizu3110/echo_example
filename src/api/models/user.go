package models

import "github.com/jinzhu/gorm"

//Dog struct
type User struct {
	gorm.Model
	Name  string `json:"name" gorm:"not null"`
	Email string `json:"email" gorm:"not null"`
}
