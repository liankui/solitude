package config

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *DataBase

type DataBase struct {
	TestDate *gorm.DB
}

func openDB(user, pass, addr, dbname string) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pass, addr, dbname)
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{})
	if err != nil {
		log.Fatalf("Database:%s connection error: %v", dbname, err)
	}
	return db
}

func InitTestDB() *gorm.DB {
	return openDB(
		viper.GetString("MySQl.User"),
		viper.GetString("MySQl.Password"),
		viper.GetString("MySQl.Addr"),
		viper.GetString("MySQl.DBname"),
	)
}

func (db *DataBase) InitDB() {
	DB = &DataBase{
		TestDate: InitTestDB(),
	}
}
