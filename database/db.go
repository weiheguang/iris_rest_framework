package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// 全局唯一的db对象
// var db *gorm.DB

type Db = gorm.DB

var db *Db

func Init(userName, password, host, port, dbName string, sqlDebug bool) {
	db = NewDb(userName, password, host, port, dbName, sqlDebug)
}

// InitDb init mysql database connection
func NewDb(userName, password, host, port, dbName string, sqlDebug bool) *Db {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true&loc=Local", userName, password, host, port, dbName)
	var mlogger logger.Interface
	if sqlDebug {
		mlogger = logger.Default.LogMode(logger.Info)
	} else {
		mlogger = logger.Default.LogMode(logger.Silent)
	}
	// var logger
	gdb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		//日志级别
		Logger: mlogger,
		NamingStrategy: schema.NamingStrategy{
			// TablePrefix: "t_",   // table name prefix, table for `User` would be `t_users`
			SingularTable: true, // use singular table name, table for `User` would be `user` with this option enabled
			// NoLowerCase: true, // skip the snake_casing of names
			// NameReplacer: strings.NewReplacer("CID", "Cid"), // use name replacer to change struct/field name before convert it to db name
		},
	})
	if err != nil {
		panic("failed to connect database")
	}
	return gdb
}

func GetDb() *Db {
	if db == nil {
		panic("db 没有初始化")
	}
	return db
}
