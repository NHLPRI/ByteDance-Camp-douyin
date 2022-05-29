package common

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"

	_ "github.com/go-sql-driver/mysql"
)

var DB *gorm.DB

//获取数据库对象
func GetDB() *gorm.DB {
	return DB
}

//读取配置文件
func initConfig() {
	log.Println(">>> Reading configuration file ......")
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")

	err := viper.ReadInConfig()
	if err != nil {
		log.Println("[InitConfig error]>>>", err)
	}
}

//建立数据库连接
func InitDbConnection() *gorm.DB {
	//读取application.yml配置文件
	initConfig()

	//配置数据库连接信息
	driverName := viper.GetString("datasource.driverName")
	host := viper.GetString("datasource.host")
	port := viper.GetString("datasource.port")
	database := viper.GetString("datasource.database")
	username := viper.GetString("datasource.username")
	password := viper.GetString("datasource.password")
	charset := viper.GetString("datasource.charset")

	//loc := viper.GetString("datasource.loc")  //时区设置，用于json时间格式
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset)

	//打开数据库连接
	db, err := gorm.Open(driverName, args)
	if err != nil {
		panic("failed to connect database: err>>>" + err.Error())
	}

	DB = db
	log.Println(">>> Database connection established successfully !!!")
	return db
}
