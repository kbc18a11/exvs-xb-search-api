package main

import (
	"log"

	"github.com/GIT_USER_ID/GIT_REPO_ID/src/common"
	"github.com/GIT_USER_ID/GIT_REPO_ID/src/config"
	"github.com/GIT_USER_ID/GIT_REPO_ID/src/repositories"
	"github.com/GIT_USER_ID/GIT_REPO_ID/src/services"
)

func main() {
	airframeService := services.AirframeService{
		ScrapeLogics:           &common.ScrapeLogicsImp{},
		FileLogics:             &common.FileLogicsImp{},
		AirframeRepository:     &repositories.AirframeRepositoryImp{DbConfig: &config.DbConfigImp{}},
		TitleOfWorkRepository:  &repositories.TitleOfWorkRepositoryImp{DbConfig: &config.DbConfigImp{}},
		PilotRepository:        &repositories.PilotRepositoryImp{DbConfig: &config.DbConfigImp{}},
		AwakenTypeRepository:   &repositories.AwakenTypeRepositoryImp{DbConfig: &config.DbConfigImp{}},
		AirframeCostRepository: &repositories.AirframeCostRepositoryImp{DbConfig: &config.DbConfigImp{}},
	}

	// 機体情報URL一覧取得
	err := airframeService.SaveAtWikiOnAirframes()

	if err != nil {
		log.Fatal(err)
	}
}
