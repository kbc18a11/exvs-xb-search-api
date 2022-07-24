package repositories

import (
	"github.com/GIT_USER_ID/GIT_REPO_ID/src/config"
	"github.com/GIT_USER_ID/GIT_REPO_ID/src/models"
)

type PilotRepository interface {
	FindByName(name string) (*models.Pilot, error)
}

type PilotRepositoryImp struct {
	DbConfig config.DbConfig
}

/*
パイロット名からパイロット情報の取得
*/
func (repository *PilotRepositoryImp) FindByName(name string) (*models.Pilot, error) {
	db, DbConnectionInitErr := repository.DbConfig.DbConnectionInit()

	if DbConnectionInitErr != nil {
		return nil, DbConnectionInitErr
	}

	pilot := &models.Pilot{}
	db.Where("name = ?", name).Find(pilot)

	dbCloseerr := repository.DbConfig.DbClose(db)
	if dbCloseerr != nil {
		return nil, dbCloseerr
	}

	return pilot, nil
}
