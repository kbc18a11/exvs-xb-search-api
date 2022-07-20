package models

import "gorm.io/gorm"

type Pilot struct {
	gorm.Model
	Name      string `gorm:"size:30;unique;not null"`
	Airframes []Airframe
}
