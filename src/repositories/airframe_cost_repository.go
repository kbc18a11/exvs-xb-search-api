package repositories

import (
	"github.com/GIT_USER_ID/GIT_REPO_ID/src/config"
	"github.com/GIT_USER_ID/GIT_REPO_ID/src/models"
)

type AirframeCostRepository interface {
	FindByCostValue(cost int) (*models.AirframeCost, error)
}

type AirframeCostRepositoryImp struct {
	DbConfig config.DbConfig
}

/*
パイロット名からパイロット情報の取得
*/
func (repository *AirframeCostRepositoryImp) FindByCostValue(cost int) (*models.AirframeCost, error) {
	db, DbConnectionInitErr := repository.DbConfig.DbConnectionInit()

	if DbConnectionInitErr != nil {
		return nil, DbConnectionInitErr
	}

	airframeCost := &models.AirframeCost{}
	db.Where("cost = ?", cost).Find(airframeCost)

	dbCloseerr := repository.DbConfig.DbClose(db)
	if dbCloseerr != nil {
		return nil, dbCloseerr
	}

	return airframeCost, nil
}
