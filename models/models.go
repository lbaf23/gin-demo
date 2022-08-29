package models

import (
	"database/sql"
	"fmt"
	"gin-demo/conf"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var SQLDB *sql.DB

func Init() {
	var err error

	connStr := fmt.Sprintf(
		"%s:%s@(%s:%s)/%s?charset=utf8",
		conf.Config.DBUser,
		conf.Config.DBPassword,
		conf.Config.DBHost,
		conf.Config.DBPort,
		conf.Config.DBName,
	)

	db, err = gorm.Open(mysql.Open(connStr))

	if err != nil {
		panic(err)
	}
	SQLDB, err = db.DB()
	if err != nil {
		panic(err)
	}
	SQLDB.SetConnMaxLifetime(time.Hour)
	SQLDB.SetMaxIdleConns(5)
	SQLDB.SetMaxOpenConns(100)

	db.AutoMigrate(&User{})

}
