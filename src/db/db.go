package db

import (
	"fmt"
	"myapp/src/api/models"

	"github.com/BurntSushi/toml"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

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

func InitDB() {
	var err error
	config := getConfig()
	db, err = gorm.Open(config.Db())
	if err != nil {
		panic(err.Error())
	}
	db.AutoMigrate(&models.Dog{})
}

func getConfig() Config {
	var config Config
	_, err := toml.DecodeFile("config.local.toml", &config)
	if err != nil {
		fmt.Println(err)
		panic("unloaded config file")
	}
	return config
}
func (d DbConfig) DSN() string {
	return fmt.Sprintf("%s:%s@%s/%s?charset=%s&parseTime=%s", d.User, d.Password, d.Server, d.Database, d.Charset, d.ParseTime)
}

func (c Config) Db() (string, string) {
	return c.Database.Driver, c.Database.DSN()
}
