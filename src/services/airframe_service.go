package services

import (
	"fmt"

	"github.com/GIT_USER_ID/GIT_REPO_ID/src/common"
)

type AirframeService struct {
	ScrapeLogics common.ScrapeLogics
}

/*
@wiki記載している全ての機体情報の保存
*/
func (arframeService *AirframeService) SaveAtWikiOnAirframes() {
	// 機体情報のURLを一覧取得
	airframeUrls := arframeService.ScrapeLogics.GetAirframeUrls()

	fmt.Println(len(airframeUrls))

	for i, airframeUrl := range airframeUrls {
		// time.Sleep(time.Second * 5)

		airframeInfo, err := arframeService.ScrapeLogics.GetAirframeInfo(airframeUrl)

		if err != nil {
			// プレイアブルキャラじゃない場合
			continue
		}

		fmt.Println(i, airframeInfo)
	}
}
