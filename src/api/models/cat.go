package models

import "github.com/jinzhu/gorm"


type Cat struct {
	gorm.Model
	Name string `json:"name"`
	Type string `json:"type"`
}