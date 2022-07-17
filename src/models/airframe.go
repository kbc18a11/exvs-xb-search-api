package models

import (
	"gorm.io/gorm"
)

type Airframe struct {
	gorm.Model
	TitleOfWorkId    int
	TitleOfWork      TitleOfWork
	PilotId          int
	Pilot            Pilot
	AirframeCostId   int
	AirframeCost     AirframeCost
	AwakenTypeId     int
	AwakenType       AwakenType
	Name             string `gorm:"size:30;not null"`
	Hp               int    `gorm:"not null"`
	ThumbnailUrl     string `gorm:"size:2083;not null"`
	IsTransformation bool   `gorm:"not null"`
	IsDeformation    bool   `gorm:"not null"`
}
