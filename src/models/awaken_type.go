package models

import "gorm.io/gorm"

type AwakenType struct {
	gorm.Model
	Name      string `gorm:"size:15;unique;not null"`
	Airframes []Airframe
}
