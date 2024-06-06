package models

import "gorm.io/gorm"

// `User` belongs to `Company`, `CompanyID` is the foreign key
type User1 struct {
	gorm.Model
	Name      string
	CompanyID int
	Company   Company
  }
  
  type Company struct {
	ID   int
	Name string
  }

  type User2 struct {
	gorm.Model
	CreditCard CreditCard
  }
  
  type CreditCard struct {
	gorm.Model
	Number string
	User2ID uint
  }