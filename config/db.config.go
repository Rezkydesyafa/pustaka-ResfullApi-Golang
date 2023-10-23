package config

var DB_DRIVER = "mysql"

var DB_HOST = "127.0.0.1"

var DB_PORT = "3306"

var DB_NAME = "pustaka_api"

var DB_USER = "root"

var DB_PASSWORD = ""

// import (
// 	"gorm.io/driver/mysql"
// 	"gorm.io/gorm"
//   )

//   func main() {
// 	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
// 	dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
//   }