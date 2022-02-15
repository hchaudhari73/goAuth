package database

import (
	"fmt"

	"github.com/hchaudhari73/goAuth/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	connStr, err := config.GetMysqlConnString()
	if err != nil {
		panic(fmt.Sprintf("Error while connecting to MySQL database: %s\n", err))
	}

	DB, err = gorm.Open(mysql.Open(*connStr), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("Error while connecting to MySQL database: %s\n", err))
	}

	fmt.Printf("Connected to MySQL database: %s\n", DB.Name())
}
