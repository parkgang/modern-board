package mysql

import (
	"fmt"
	"log"

	"github.com/parkgang/modern-board/model"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	Client *gorm.DB
)

func init() {
	mysqlHost := viper.GetString("database.mysql.host")
	mysqlPort := viper.GetString("database.mysql.port")
	mysqlUser := viper.GetString("database.mysql.user")
	mysqlPass := viper.GetString("database.mysql.pass")
	mysqlDbname := viper.GetString("database.mysql.dbname")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", mysqlUser, mysqlPass, mysqlHost, mysqlPort, mysqlDbname)

	log.Println("mysql connection...")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("DB 연결에 실패하였습니다.\n\t" + err.Error())
	}

	log.Println("mysql migration...")
	err = db.AutoMigrate(&model.User{})
	if err != nil {
		panic("DB 마이그레이션에 실패하였습니다.\n\t" + err.Error())
	}

	log.Printf("mysql starting at %s\n", dsn)
	Client = db
}
