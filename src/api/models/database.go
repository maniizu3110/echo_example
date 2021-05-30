package models

import (
	"gorm.io/gorm"
)

type Config struct {
	Database DbConfig
}

type DbConfig struct {
	Driver    string
	Server    string
	User      string
	Password  string
	Database  string
	Charset   string
	ParseTime string
}

type Dog struct {
	gorm.Model
	Name string `json:"name"`
	Type string `json:"type"`
}
