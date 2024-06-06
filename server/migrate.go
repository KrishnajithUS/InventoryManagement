package server

import (
	"log"
	"InventoryManagement/database"
	"InventoryManagement/models"
	
)

func Migrate() {
	db := database.GetDB()
	err := db.AutoMigrate(
		&models.Branch{}, &models.Team{}, &models.Employee{},
		&models.ParkingSpot{},&models.User2{},&models.Company{},
		&models.CreditCard{},&models.User1{},&models.User{},&models.Attendence{},
	)
	if err != nil {
		log.Fatal("Unable to migrate", err)
	}
}
