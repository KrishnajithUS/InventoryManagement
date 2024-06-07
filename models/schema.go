package models

import (
	"time"

	"gorm.io/datatypes"
)

type Base struct {
	ID        int `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Branch struct {
	Base
	Name         string
	Location     string
	Address      string
	ParkingSpots []ParkingSpot 
	Teams        []Team 
}



type Team struct {
	Base
	Name      string
	Employees []Employee
	BranchId uint
}

type Employee struct {
	Base
	PersonalInfo datatypes.JSON
	Salary       float64
	TeamId uint
	Username string `gorm:"size:255;not null;unique"`
	Password string `gorm:"size:255;not null;"`
	Attendences []Attendence
}

type ParkingSpot struct {
	Base
	SpotName string
	BranchId uint
}

type Attendence struct{
	Date time.Time `gorm:"index:idx_name,unique"`
	EmployeeId uint `gorm:"index:idx_name,unique"`
}