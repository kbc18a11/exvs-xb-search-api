package main

import (
	"github.com/GIT_USER_ID/GIT_REPO_ID/src/common"
	"github.com/GIT_USER_ID/GIT_REPO_ID/src/config"
	"github.com/GIT_USER_ID/GIT_REPO_ID/src/repositories"
	"github.com/GIT_USER_ID/GIT_REPO_ID/src/services"
)

func main() {
	airframeService := services.AirframeService{
		ScrapeLogics:          &common.ScrapeLogicsImp{},
		TitleOfWorkRepository: &repositories.TitleOfWorkRepositoryImp{DbConfig: &config.DbConfigImp{}},
	}

	// 機体情報URL一覧取得
	airframeService.SaveAtWikiOnAirframes()
}
