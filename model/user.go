package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name      string `gorm:"type:varchar(20);not null"`
	Telephone string `gorm:"varchar(10);not null;unique"`
	Password  string `gorm:"size:255;not null"`
}