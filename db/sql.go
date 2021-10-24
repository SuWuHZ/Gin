package db

import (
	"Gin/utils"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DBInnoDB *gorm.DB

func init() {
	DBInnoDB = Connect()
}

func Connect() *gorm.DB {
	conn := utils.ServerConf.Database.SqlConn
	fmt.Println("conn",conn)
	db, err := gorm.Open(mysql.Open(conn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	return db
}

func GetDB() *gorm.DB {
	if DBInnoDB != nil {
		return DBInnoDB
	} else {
		return Connect()
	}
}
