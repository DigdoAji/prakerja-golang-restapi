package configs

import (
	"fmt"
	"os"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	comicdatabase "ujk-golang/models/comic/database"
)


var DB *gorm.DB

func InitDatabase(){
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local" , 
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))
    var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Database connection failed")
	}
	Migration()
}

func Migration(){
	DB.AutoMigrate(comicdatabase.Comic{})
}