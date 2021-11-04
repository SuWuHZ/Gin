package db

import (
	"Gin/models"
	"Gin/utils"

	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var MysqlDB *gorm.DB

func init() {
	MysqlDB = Connect()
	MysqlDB.Set("gorm:table_options", "CHARSET=utf8").AutoMigrate(models.User{})
}

func Connect() *gorm.DB {
	conn := utils.ServerConf.Database.SqlConn
	fmt.Println("conn", conn)
	db, err := gorm.Open(mysql.Open(conn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	mySQL, err := db.DB()
	if err != nil {
		fmt.Println(err)
	}
	mySQL.SetMaxIdleConns(10)
	return db
}

func GetDB() *gorm.DB {
	if MysqlDB != nil {
		return MysqlDB
	} else {
		return Connect()
	}
}
