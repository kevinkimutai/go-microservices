package db

import (
	"os"
	"simple_bank/model"

	"github.com/gofiber/fiber/v2/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func InitMigration() {
	ConnectionStr := os.Getenv("MYSQL_CONN_STRING")
	log.Info(ConnectionStr)

	DB, err = gorm.Open(mysql.Open(ConnectionStr), &gorm.Config{})
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}

	// Migrate the schema
	DB.AutoMigrate(&model.Account{}, &model.Entries{}, &model.Transfer{})
}
