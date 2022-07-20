package main

import (
	"log"

	"github.com/GIT_USER_ID/GIT_REPO_ID/src/config"
	"github.com/GIT_USER_ID/GIT_REPO_ID/src/models"
)

func main() {
	dbConfig := &config.DbConfigImp{}
	db, err := dbConfig.DbConnectionInit()

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
