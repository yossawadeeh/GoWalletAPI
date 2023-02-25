package orm

import (
	"my-go-wallet/model"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// for external file call (first character is upper case)
var Db *gorm.DB
var err error

func InitDB() {
	dsn := os.Getenv("MYSQL_DNS")
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	Db.AutoMigrate(&model.User{})
}
