module myapp

go 1.16

require (
	// firebase.google.com/go v3.13.0+incompatible
	firebase.google.com/go/v4 v4.5.0 // indirect
	github.com/BurntSushi/toml v0.3.1
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/jinzhu/gorm v1.9.16
	github.com/labstack/echo/v4 v4.3.0
// gorm.io/gorm v1.21.10 // indirect　//jinzhu/gormを呼び出して欲しいのでコメントアウト
)
