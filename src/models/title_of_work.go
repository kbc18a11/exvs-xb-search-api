package models

import "gorm.io/gorm"

type TitleOfWork struct {
	gorm.Model
	Name      string `gorm:"size:60;unique;not null"`
	Airframes []Airframe
}
