package models

import (
	"InventoryManagement/database"
	"InventoryManagement/utils/token"
	"fmt"
	"html"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

func (u *User) CreateUser() (*User, error){
	db := database.GetDB()
	res := db.Create(&u)
	return u, res.Error
}


func (e *Employee) CreateUser() (*Employee, error){
	db := database.GetDB()
	res := db.Create(&e)
	return e, res.Error
}

func (u User) GetUserById(pk int) (User, error) {
	db := database.GetDB()
	res := db.First(&u, pk)
	u.RemoveHashedPassword()
	return u, res.Error

}

func (u *User) RemoveHashedPassword(){
	u.Password = ""
}

func (u *Employee) BeforeSave(*gorm.DB) error {
	fmt.Println("Before save")
	//turn password into hash
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password),bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)

	//remove spaces in username 
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))

	return nil

}

func (u *User) BeforeSave(*gorm.DB) error {
	fmt.Println("Before save")
	//turn password into hash
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password),bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)

	//remove spaces in username 
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))

	return nil

}

func CheckPassword(password ,hashedPassword string) error{
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword),[]byte(password))
}
// TODO : pluck
func (u User) LoginUser(username string, password string) (string, error){
	db := database.GetDB()
	err := db.Where("username = ?",username).Find(&u).Error
	if err != nil {
		return "", err
	}
	err = CheckPassword(password, u.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}
	token,err := token.GenerateToken(uint(u.ID))

	if err != nil {
		return "",err
	}

	return token,nil
}

func (a Attendence) MarkAttendence(emp_id uint) error{
	db := database.GetDB()
	date  := time.Now().Format("2006-01-02")
	a.Date, _ = time.Parse("2006-01-02", date)
	a.EmployeeId = emp_id
	res := db.Create(&a)
	return res.Error
}


func (u Employee) LoginUser(username string, password string) (string, error){
	db := database.GetDB()
	err := db.Where("username = ?",username).Find(&u).Error
	if err != nil {
		return "", err
	}
	err = CheckPassword(password, u.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}
	token,err := token.GenerateToken(uint(u.ID))

	if err != nil {
		return "",err
	}
	var attendence Attendence
	err = attendence.MarkAttendence(uint(u.ID))
	if err != nil{
		return "", err
	}
	return token,nil
}