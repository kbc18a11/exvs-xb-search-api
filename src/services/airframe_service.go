package services

import (
	"fmt"

	"github.com/GIT_USER_ID/GIT_REPO_ID/src/common"
	"github.com/GIT_USER_ID/GIT_REPO_ID/src/repositories"
)

type AirframeService struct {
	ScrapeLogics          common.ScrapeLogics
	TitleOfWorkRepository repositories.TitleOfWorkRepository
}

/*
@wiki記載している全ての機体情報の保存
*/
func (arframeService *AirframeService) SaveAtWikiOnAirframes() {
	// 機体情報のURLを一覧取得
	airframeUrls := arframeService.ScrapeLogics.GetAirframeUrls()

	for _, airframeUrl := range airframeUrls {
		// time.Sleep(time.Second * 5)

		airframeInfo, err := arframeService.ScrapeLogics.GetAirframeInfo(airframeUrl)

		if err != nil {
			// プレイアブルキャラじゃない場合
			continue
		}

		fmt.Println(airframeInfo.Name, airframeInfo.TitleOfWork)

		arframeService.TitleOfWorkRepository.FindByName(airframeInfo.TitleOfWork)
	}
}
