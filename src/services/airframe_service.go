package services

import (
	"log"

	"github.com/GIT_USER_ID/GIT_REPO_ID/src/common"
	"github.com/GIT_USER_ID/GIT_REPO_ID/src/models"
	"github.com/GIT_USER_ID/GIT_REPO_ID/src/repositories"
)

type AirframeService struct {
	ScrapeLogics           common.ScrapeLogics
	FileLogics             common.FileLogics
	AirframeRepository     repositories.AirframeRepository
	TitleOfWorkRepository  repositories.TitleOfWorkRepository
	PilotRepository        repositories.PilotRepository
	AwakenTypeRepository   repositories.AwakenTypeRepository
	AirframeCostRepository repositories.AirframeCostRepository
}

/*
@wiki記載している全ての機体情報の保存
*/
func (airframeService *AirframeService) SaveAtWikiOnAirframes() error {
	// 機体情報のURLを一覧取得
	airframeUrls := airframeService.ScrapeLogics.GetAirframeUrls()

	for _, airframeUrl := range airframeUrls {
		// time.Sleep(time.Second * 5)

		// 機体情報の取得
		airframeInfo, err := airframeService.ScrapeLogics.GetAirframeInfo(airframeUrl)
		if err != nil {
			// プレイアブルキャラじゃない場合
			continue
		}

		airframe := &models.Airframe{}
		airframe.Name = airframeInfo.Name
		airframe.AirframeInfoUrl = airframeInfo.AirframeInfoUrl

		// 作品タイトル情報の取得
		titleOfWork, err := airframeService.TitleOfWorkRepository.FindByName(airframeInfo.TitleOfWorkName)
		if err != nil {
			return err
		}
		airframe.TitleOfWorkId = int(titleOfWork.ID)

		// パイロット情報の取得
		pilot, err := airframeService.PilotRepository.FindByName(airframeInfo.PilotName)
		if err != nil {
			return err
		}
		airframe.PilotId = int(pilot.ID)

		// 覚醒タイプ情報の取得
		awakenType, err := airframeService.AwakenTypeRepository.FindByName(airframeInfo.AwakenTypeName)
		if err != nil {
			return err
		}
		airframe.AwakenTypeId = int(awakenType.ID)

		// コスト情報の取得
		costType, airframeCostRepositoryFindByCostValue := airframeService.AirframeCostRepository.FindByCostValue(airframeInfo.AirframeCostValue)
		if airframeCostRepositoryFindByCostValue != nil {
			log.Fatal(airframeCostRepositoryFindByCostValue)
		}
		airframe.AirframeCostId = int(costType.ID)

		// サムネイル画像の保存
		thumbnailImage, err := airframeService.FileLogics.DownloadSaveImage(airframeInfo.ThumbnailUrl, "src/images/airframe_thumbnails/")
		if err != nil {
			return err
		}
		airframe.ThumbnailImageFilePath = thumbnailImage

		err = airframeService.AirframeRepository.Create(airframe)
		if err != nil {
			return err
		}
	}

	return nil
}
