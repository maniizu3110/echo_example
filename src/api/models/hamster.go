package models

import "github.com/jinzhu/gorm"

//Hamster struct
type Hamster struct {
	gorm.Model
	Name string `json:"name"`
	Type string `json:"type"`
}