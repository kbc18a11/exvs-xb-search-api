package repositories

import (
	"github.com/GIT_USER_ID/GIT_REPO_ID/src/config"
	"github.com/GIT_USER_ID/GIT_REPO_ID/src/models"
)

type AwakenTypeRepository interface {
	FindByName(name string) (*models.AwakenType, error)
}

type AwakenTypeRepositoryImp struct {
	DbConfig config.DbConfig
}

/*
覚醒タイプ名から覚醒タイプ情報の取得
*/
func (repository *AwakenTypeRepositoryImp) FindByName(name string) (*models.AwakenType, error) {
	db, DbConnectionInitErr := repository.DbConfig.DbConnectionInit()

	if DbConnectionInitErr != nil {
		return nil, DbConnectionInitErr
	}

	awakenType := &models.AwakenType{}
	db.Where("name = ?", name).Find(awakenType)

	dbCloseerr := repository.DbConfig.DbClose(db)
	if dbCloseerr != nil {
		return nil, dbCloseerr
	}

	return awakenType, nil
}
