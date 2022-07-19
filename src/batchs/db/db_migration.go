package main

import (
	"log"

	"github.com/GIT_USER_ID/GIT_REPO_ID/src/common"
	"github.com/GIT_USER_ID/GIT_REPO_ID/src/models"
)

func main() {
	db, err := common.DbConnectionInit()

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(
		&models.TitleOfWork{},
		&models.Pilot{},
		&models.AirframeCost{},
		&models.AwakenType{},
		&models.Airframe{},
	)

	AddMasterData(db)
}
