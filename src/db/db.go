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
	User      string
	Password  string
	Address   string
	DBName	  string
}

//InitDB start MysqlDB
func InitDB() *gorm.DB {
	var err error
	config := getConfig()
	db, err = gorm.Open(config.Db())
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database.")
	}
	db.LogMode(true)
	return db
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
// func (d dbConfig) DSN() string {
// 	return fmt.Sprintf("%s:%s@%s/%s?charset=%s&parseTime=%s", d.User, d.Password, d.Server, d.Database, d.Charset, d.ParseTime)
// }

func (c config) Db() (string, string) {
	SETTING := c.Database.User + ":" + c.Database.Password + "@tcp(" + c.Database.Address +":3306)/" +c.Database.DBName+ "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"
	return c.Database.Driver,SETTING
}
