package models

import (
	"gorm.io/gorm"
)

type Airframe struct {
	gorm.Model
	TitleOfWorkId          int
	PilotId                int
	AirframeCostId         int
	AwakenTypeId           int
	Name                   string `gorm:"size:60;unique;not null"`
	Hp                     int    `gorm:"not null"`
	AirframeInfoUrl        string `gorm:"size:2083;not null"`
	ThumbnailImageFilePath string `gorm:"size:2083;not null"`
	IsTransformation       bool   `gorm:"not null"`
	IsDeformation          bool   `gorm:"not null"`
}
