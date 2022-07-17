package main

import (
	"github.com/GIT_USER_ID/GIT_REPO_ID/src/common"
	"github.com/GIT_USER_ID/GIT_REPO_ID/src/services"
)

func main() {
	airframeService := services.AirframeService{ScrapeLogics: &common.ScrapeLogicsImp{}}

	// 機体情報URL一覧取得
	airframeService.SaveAtWikiOnAirframes()
}
