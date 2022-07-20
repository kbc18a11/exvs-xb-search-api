package repositories

import (
	"fmt"

	"github.com/GIT_USER_ID/GIT_REPO_ID/src/config"
	"github.com/GIT_USER_ID/GIT_REPO_ID/src/models"
)

type TitleOfWorkRepository interface {
	FindByName(name string) (*models.TitleOfWork, error)
}

type TitleOfWorkRepositoryImp struct {
	DbConfig config.DbConfig
}

/*
作品タイトル名から作品タイトル情報の取得
*/
func (repository *TitleOfWorkRepositoryImp) FindByName(name string) (*models.TitleOfWork, error) {
	db, err := repository.DbConfig.DbConnectionInit()

	if err != nil {
		return nil, err
	}

	// 作品タイトル名から作品タイトル情報の検索
	resultTitleOfWork := &models.TitleOfWork{}
	db.Where("name = ?", name).Find(resultTitleOfWork)

	dbCloseerr := repository.DbConfig.DbClose(db)
	if dbCloseerr != nil {
		return nil, err
	}

	fmt.Println(resultTitleOfWork.Name)

	return resultTitleOfWork, nil
}
