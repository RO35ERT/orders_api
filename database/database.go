package database

import (
	"log"
	"os"

	"github.com/ro35ert/orders_api/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct{
	Db *gorm.DB
}

var Database DbInstance

func DOnnectDb(){
	
	dsn := os.Getenv("DB_USER")+":"+os.Getenv("DB_PASS")+"@tcp(127.0.0.1:3306)/order_api?charset=utf8mb4&parseTime=True&loc=Local"
	db,err := gorm.Open(mysql.Open(dsn),&gorm.Config{})

	if err != nil{
		log.Fatal("Failed to connect to the database \n", err.Error())
		os.Exit(2)
	}
	log.Println("Connected successfully!")

	db.Logger = logger.Default.LogMode(logger.Info)

	//TO Do: Migrations
	db.AutoMigrate(&models.User{},&models.Product{},&models.Order{})

	Database = DbInstance{Db: db}
	
}