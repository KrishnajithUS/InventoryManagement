package controllers

import "gorm.io/datatypes"

type RegisterInput struct{
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginInput struct{
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type EmployeeLoginInput struct{
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type EmployeeCreateInput struct{
	PersonalInfo datatypes.JSON
	Salary       float64
	TeamId uint
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}