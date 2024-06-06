package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	
)


var db *gorm.DB

func Init(){
	var err error
	// dbURL := "postgres://krishnajith:krishna123%40@localhost:5432/rssagg?sslmode=disable"
	dsn := "host=localhost user=krishnajith password=krishna123@ dbname=Inventory port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{}); 

	if err != nil{
		log.Fatal("Unable to connect to db ",err)
	}

}


func GetDB() *gorm.DB{
	return db
}