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

	for _, v := range airframeUrls {
		fmt.Printf("%s ", v) // a bc def
	}
}
