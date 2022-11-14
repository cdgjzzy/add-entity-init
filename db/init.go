package db

import (
	"fmt"

	"gopkg.in/ini.v1"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
	DBConfig `ini:"db"`
}
type DBConfig struct {
	Host     string `ini:"host"`
	Port     string `ini:"port"`
	Username string `ini:"username"`
	Password string `ini:"password"`
	DataBase string `ini:"database"`
}

var (
	DB *gorm.DB
)

func init() {
	config := &Config{}
	err := ini.MapTo(config, "./config.ini")
	if err != nil {
		panic(err)
	}
	dbConfig := config.DBConfig
	fmt.Printf("config:%v\n", config)
	DSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True",
		dbConfig.Username, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.DataBase)
	fmt.Println("DSN:", DSN)
	db, err := gorm.Open(mysql.Open(DSN), &gorm.Config{
		PrepareStmt:            true,
		SkipDefaultTransaction: true, // 禁用默认事务
	})
	if err != nil {
		panic(err)
	}
	DB = db
}
