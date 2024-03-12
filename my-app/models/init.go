package models

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	db  *gorm.DB
	err error
)

func init() {
	viper.SetConfigName("app")
	viper.AddConfigPath("./")
	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	mapString := viper.GetStringMapString("MysqlK")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", mapString["username"], mapString["password"], mapString["host"], mapString["port"], mapString["mysql"])
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	err = db.AutoMigrate()
	if err != nil {
		return
	}
}
