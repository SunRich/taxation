package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"fmt"
)

var DB *gorm.DB

func InitMysql() {
	fmt.Println(os.Getenv("mysqlconn"))
	DB, _ = gorm.Open("mysql", os.Getenv("mysqlconn"))
	DB.DB().SetMaxIdleConns(10)
	DB.DB().SetMaxOpenConns(100)
	DB.LogMode(true)
}
