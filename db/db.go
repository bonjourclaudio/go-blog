package db

import (
	"github.com/claudioontheweb/go-blog/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
)

var DB *gorm.DB
func ConnectDB() *gorm.DB {

	config.GetConfig()

	// Get DB config from config.json
	var dbHost = viper.GetString("DB_HOST")
	var dbUser = viper.GetString("DB_USERNAME")
	var dbName = viper.GetString("DB_NAME")
	var dbPassword = viper.GetString("DB_PASSWORD")

	// Connect to DB
	db, dbError2 := gorm.Open("mysql", dbUser+":"+ dbPassword +"@tcp(" + dbHost+ ":3306)/" + dbName + "?charset=utf8&parseTime=True&loc=Local")
	if dbError2 != nil {
		panic(dbError2)
	}

	db.DB().SetMaxIdleConns(0)

	return db

}