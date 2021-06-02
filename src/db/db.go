package db

import (
	"fmt"
	"myapp/src/api/models"

	"github.com/BurntSushi/toml"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type config struct {
	Database dbConfig
}

type dbConfig struct {
	Driver    string
	Server    string
	User      string
	Password  string
	Database  string
	Charset   string
	ParseTime string
}

func OpenDB() *gorm.DB {
	var err error
	config := getConfig()
	db, err = gorm.Open(config.Db())
	if err != nil {
		panic("failed to connect database.")
	}
	autoAllMigrate(db)
	db.LogMode(true)
	return db
}

func autoAllMigrate(db *gorm.DB) error {
	db.AutoMigrate(&models.User{})
	return nil
}

func getConfig() config {
	var config config
	_, err := toml.DecodeFile("config.local.toml", &config)
	if err != nil {
		fmt.Println(err)
		panic("unloaded config file")
	}
	return config
}
func (d dbConfig) DSN() string {
	return fmt.Sprintf("%s:%s@%s/%s?charset=%s&parseTime=%s", d.User, d.Password, d.Server, d.Database, d.Charset, d.ParseTime)
}

func (c config) Db() (string, string) {
	return c.Database.Driver, c.Database.DSN()
}
