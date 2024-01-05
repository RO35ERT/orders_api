package database

import (
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct{
	Db *gorm.DB
}

var Database DbInstance

func DOnnectDb(){
	// dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	db,err := gorm.Open(mysql.Open("api.db",&gorm.Config{}))

	if err != nil{
		log.Fatal("Failed to connect to the database \n", err.Error())
		os.Exit(2)
	}
	log.Println("Connected successfully!")

	db.Logger = logger.Default.LogMode(logger.Info)

	
}