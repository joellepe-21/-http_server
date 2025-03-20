package models

//import "gorm.io/gorm"

type User struct{
	ID uint `gorm:"primaryKey;autoIncrement"`
	Login string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
}