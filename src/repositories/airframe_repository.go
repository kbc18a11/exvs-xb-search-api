package repositories

import (
	"github.com/GIT_USER_ID/GIT_REPO_ID/src/config"
	"github.com/GIT_USER_ID/GIT_REPO_ID/src/models"
)

type AirframeRepository interface {
	Create(*models.Airframe) error
	FindByName(name string) (*models.Airframe, error)
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
	db, DbConnectionInitErr := repository.DbConfig.DbConnectionInit()

	if DbConnectionInitErr != nil {
		return nil, DbConnectionInitErr
	}

	airframe := &models.Airframe{}

	// 機体情報名から機体情報の検索
	db.Where("name LIKE ?", name).Find(airframe)

	dbCloseerr := repository.DbConfig.DbClose(db)
	if dbCloseerr != nil {
		return nil, dbCloseerr
	}

	if airframe.Name == "" {
		// 機体情報が存在しない場合
		return nil, nil
	}

	return airframe, nil
}
