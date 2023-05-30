package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// 全局唯一的db对象
var db *Db

// 封装gorm.DB
type Db struct {
	*gorm.DB
	// logger DBLogger
}

// InitDb init mysql database connection
func Init(userName, password, host, port, dbName string) *Db {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true&loc=Local", userName, password, host, port, dbName)
	// log.Printf("current mysql connection dsn: %s\n", dsn)
	// var err error
	gdb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		//日志级别
		Logger: logger.Default.LogMode(logger.Silent),
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
	db = &Db{gdb}
	return db
}

// 初始化全局mock db对象
// func InitMockDb(modk_db *sql.DB) {
// 	db, _ = gorm.Open(mysql.New(mysql.Config{
// 		Conn:                      modk_db,
// 		SkipInitializeWithVersion: true,
// 	}), &gorm.Config{})
// }

// GetDb return instance of database
func GetDb() *Db {
	if db != nil {
		return db
	}
	panic("db 没有初始化")
}
