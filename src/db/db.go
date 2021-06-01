package db

import (
	"fmt"

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

//InitDB start MysqlDB
func InitDB() {
	var err error
	config := getConfig()
	db, err = gorm.Open(config.Db())
	defer db.Close()

	if err != nil {
		panic(err.Error())
	}
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

func OpenDB() *gorm.DB {
	var err error
	config := getConfig()
	db, err = gorm.Open(config.Db())
	if err != nil {
        panic("failed to connect database.")
    }
	db.LogMode(true)
	return db
}
