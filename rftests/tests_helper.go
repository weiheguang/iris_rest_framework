package rftests

import (
	"fmt"
	"log"
	"time"

	"github.com/weiheguang/iris_rest_framework/database"
)

// 初始换数据库

// 创建数据库
func CreateDatabase(db *database.Db, dbName string) {
	createSQL := fmt.Sprintf(
		"CREATE DATABASE IF NOT EXISTS `%s` CHARACTER SET utf8mb4;",
		dbName,
	)
	err := db.Exec(createSQL).Error
	if err != nil {
		log.Fatal(err)
	}
	// 切换数据库
	err = db.Exec("use " + dbName).Error
	if err != nil {
		log.Fatal(err)
	}
	log.Println("create database success:", dbName)
}

// 删除数据库
func DropDatabase(db *database.Db, dbName string) {
	dropSQL := fmt.Sprintf(
		"DROP DATABASE IF EXISTS `%s`;",
		dbName,
	)
	err := db.Exec(dropSQL).Error
	if err != nil {
		log.Fatal(err)
	}
	log.Println("drop database success:", dbName)
}

// 创建表
func CreateTables(db *database.Db, sql string) {
	result := db.Exec(sql)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
}

// 插入数据
func InsertData(db *database.Db, sql string, args ...interface{}) {
	result := db.Exec(sql, args...)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
}

// 根据当前时间戳生成测试数据库名字
func GetTestDbName() string {
	return fmt.Sprintf("test_%d", time.Now().UnixNano())
}
