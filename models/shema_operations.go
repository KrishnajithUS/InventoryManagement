package models

import (
	"InventoryManagement/database"
	"fmt"

	"gorm.io/gorm"
)

func (b Branch) Create() (Branch, error) {
	db := database.GetDB()
	res := db.Create(&b)
	return b, res.Error
}

func (b Branch) FindById(pk int) (Branch, error) {
	db := database.GetDB()
	res := db.First(&b, pk)
	return b, res.Error

}

func (b Branch) Update(pk int, branchUpdated Branch) (Branch, error) {
	var res *gorm.DB
	db := database.GetDB()
	b, err := b.FindById(pk)
	if err != nil {
		return b, err
	}
	res = db.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&branchUpdated)
	fmt.Println(res,"res")
	fmt.Println(branchUpdated.ParkingSpots)
	return branchUpdated, res.Error
}

func (b Branch) Delete(pk int) (error) {
	var res *gorm.DB
	db := database.GetDB()
	b, err := b.FindById(pk)
	if err != nil {
		return err
	}
	res = db.Select("Teams","ParkingSpots").Delete(&b)
	return  res.Error
}

func GetAllBranches() ([]Branch, error) {
	var bRes []Branch
	db := database.GetDB()
	result := db.Preload("ParkingSpots").Preload("Teams").Preload("Teams.Employees").Omit("Teams.Employee.Password").Find(&bRes)
	fmt.Println("result", result)
	err := result.Error
	return bRes, err
}

func (t Team) Create() (Team, error) {
	db := database.GetDB()
	res := db.Create(&t)
	return t, res.Error
}

func (t ParkingSpot) Create() (ParkingSpot, error) {
	db := database.GetDB()
	res := db.Create(&t)
	fmt.Println((res))
	return t, res.Error

}
