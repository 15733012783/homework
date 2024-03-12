package model

import (
	"fmt"
	"github.com/olivere/elastic/v7"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

var db *gorm.DB
var client *elastic.Client

func Init() {
	var err error
	viper.SetConfigName("userrpc")
	viper.AddConfigPath("./etc")
	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	mapString := viper.GetStringMapString("Mysql")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", mapString["username"], mapString["password"], mapString["host"], mapString["port"], mapString["mysqlbase"])
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		fmt.Println(err, "数据库链接失败")
		return
	}

	client, err = elastic.NewClient(elastic.SetURL("http://127.0.0.1:9201"), elastic.SetSniff(false))
	if err != nil {
		// Handle error
		panic(err)
	}

	err = db.AutoMigrate(new(User))
	if err != nil {
		log.Println("创建数据表失败")
		return
	}
}
