package models

import "gorm.io/gorm"

type Pilot struct {
	gorm.Model
	Name      string `gorm:"size:20;not null"`
	Airframes []Airframe
}
