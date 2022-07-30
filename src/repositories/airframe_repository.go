package repositories

import (
	"github.com/GIT_USER_ID/GIT_REPO_ID/src/config"
	"github.com/GIT_USER_ID/GIT_REPO_ID/src/models"
	openApiModels "github.com/GIT_USER_ID/GIT_REPO_ID/src/models/open_api_schema"
)

type AirframeRepository interface {
	/*
		機体情報レコードの作成
	*/
	Create(*models.Airframe) error

	/*
		機体情報名から機体情報の取得
	*/
	FindByName(name string) (*models.Airframe, error)

	/*
		機体情報一覧取得
	*/
	GetAirframes(
		offset int,
		limit int,
		airframeName string,
		pilotName string,
		costValue int,
		titleOfWorkName string,
		awakenTypeName string,
	) ([]openApiModels.AirframeInfo, error)
}

type AirframeRepositoryImp struct {
	DbConfig config.DbConfig
}

/*
機体情報レコードの作成
*/
func (repository AirframeRepositoryImp) Create(airframe *models.Airframe) error {
	db, err := repository.DbConfig.DbConnectionInit()

	if err != nil {
		return err
	}

	db.Create(airframe)

	err = repository.DbConfig.DbClose(db)
	if err != nil {
		return err
	}

	return nil
}

/*
機体情報名から機体情報の取得
*/
func (repository AirframeRepositoryImp) FindByName(name string) (*models.Airframe, error) {
	db, err := repository.DbConfig.DbConnectionInit()

	if err != nil {
		return nil, err
	}

	airframe := &models.Airframe{}

	// 機体情報名から機体情報の検索
	db.Where("name LIKE ?", name).Find(airframe)

	err = repository.DbConfig.DbClose(db)
	if err != nil {
		return nil, err
	}

	if airframe.Name == "" {
		// 機体情報が存在しない場合
		return nil, nil
	}

	return airframe, nil
}

/*
機体情報一覧取得
*/
func (repository AirframeRepositoryImp) GetAirframes(
	offset int,
	limit int,
	airframeName string,
	pilotName string,
	costValue int,
	titleOfWorkName string,
	awakenTypeName string,
) ([]openApiModels.AirframeInfo, error) {
	db, err := repository.DbConfig.DbConnectionInit()

	if err != nil {
		return nil, err
	}

	db = db.Table("airframes").
		Select("airframes.id,airframes.name,airframes.hp,airframes.airframe_info_url,airframes.thumbnail_image_file_path," +
			"airframe_costs.cost,awaken_types.name,pilots.name,title_of_works.name").
		Joins("JOIN airframe_costs ON airframes.airframe_cost_id = airframe_costs.id").
		Joins("JOIN awaken_types ON airframes.awaken_type_id = awaken_types.id").
		Joins("JOIN pilots ON airframes.pilot_id = pilots.id").
		Joins("JOIN title_of_works ON airframes.title_of_work_id = title_of_works.id")

	if airframeName != "" {
		// 機体名の検索
		db = db.Where("airframes.name LIKE ?", "%"+airframeName+"%")
	}

	if costValue > 0 {
		// コストの検索
		db = db.Where("airframe_costs.cost = ?", costValue)
	}

	if awakenTypeName != "" {
		// 作品タイトル名の検索
		db = db.Where("awaken_types.name LIKE ?", awakenTypeName)
	}

	if pilotName != "" {
		// パイロット名の検索
		db = db.Where("pilots.name LIKE ?", pilotName)
	}

	rows, err := db.Debug().Limit(limit).Offset(offset).Rows()
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	airframeInfos := []openApiModels.AirframeInfo{}
	for rows.Next() {
		// レコードの取得
		airframeInfo := openApiModels.AirframeInfo{}
		err = rows.Scan(
			&airframeInfo.Id,
			&airframeInfo.Name,
			&airframeInfo.Hp,
			&airframeInfo.AirframeInfoUrl,
			&airframeInfo.ThumbnailImageUrl,
			&airframeInfo.Cost,
			&airframeInfo.AwakenType,
			&airframeInfo.Pilot,
			&airframeInfo.TitleOfWork,
		)

		if err != nil {
			return nil, err
		}

		airframeInfos = append(airframeInfos, airframeInfo)
	}

	err = repository.DbConfig.DbClose(db)
	if err != nil {
		return nil, err
	}

	return airframeInfos, nil
}
