package db

import (
	"fmt"
	"os"

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
func InitDB() *gorm.DB {
	var err error
	config := getConfig()
	db, err = gorm.Open(config.Db())
	if err != nil {
		panic("failed to connect database.")
	}
	db.LogMode(true)
	return db
}

func getConfig() config {
	var config config
	//本番環境で環境変数を変える
	_, err := toml.DecodeFile("config.local.toml", &config)
	if err != nil {
		panic("unloaded config file")
	}
	return config
}
func (d dbConfig) DSN() string {
	    if os.Getenv("DB_ENV") == "production" {
        d.User = os.Getenv("DB_USER")
        d.Password = os.Getenv("DB_PASS")
		d.Server = os.Getenv("DB_ADDRESS")
		//TODO:間違っている可能性あるので確認
		return fmt.Sprintf(d.User+":"+d.Password+"@tcp("+d.Server+":3306)/"+d.Database)
    }
	return fmt.Sprintf("%s:%s@%s/%s?charset=%s&parseTime=%s", d.User, d.Password, d.Server, d.Database, d.Charset, d.ParseTime)
}

func (c config) Db() (string, string) {
	return c.Database.Driver, c.Database.DSN()
}
