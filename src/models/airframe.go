package models

import "time"

type Airframe struct {
	Id               int
	AirframeCostId   int
	TitleOfWorkId    int
	PilotId          int
	AwakeningId      int
	Name             string
	Hp               int
	ThumbnailUrl     string
	IsTransformation bool
	IsDeformation    bool
	CreateAt         time.Time
	UpdateAt         time.Time
}
