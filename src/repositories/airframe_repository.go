package repositories

import (
	"github.com/GIT_USER_ID/GIT_REPO_ID/src/config"
	"github.com/GIT_USER_ID/GIT_REPO_ID/src/models"
)

type AirframeRepository interface {
	Create(*models.Airframe) error
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

	return nil
}
