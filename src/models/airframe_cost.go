package models

import "gorm.io/gorm"

type AirframeCost struct {
	gorm.Model
	Cost      int `gorm:"index;size:3000;not null"`
	Airframes []Airframe
}
