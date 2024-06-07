package models


type User struct{
	Base `gorm:"embedded"`
	Username string `gorm:"size:255;not null;unique;false;<-:create"`
	Password string `gorm:"size:255;not null;"`
}