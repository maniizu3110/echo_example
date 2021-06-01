package models

import "github.com/jinzhu/gorm"

//Dog struct
type Dog struct {
	gorm.Model
	Name string `json:"name"`
	Type string `json:"type"`
}