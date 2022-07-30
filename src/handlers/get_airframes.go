package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/GIT_USER_ID/GIT_REPO_ID/src/common"
	"github.com/GIT_USER_ID/GIT_REPO_ID/src/config"
	models "github.com/GIT_USER_ID/GIT_REPO_ID/src/models/open_api_schema"
	"github.com/GIT_USER_ID/GIT_REPO_ID/src/repositories"
	"github.com/GIT_USER_ID/GIT_REPO_ID/src/services"
	"github.com/labstack/echo/v4"
)

/*
機体情報の一覧取得
*/
func (c *Container) GetAirframes(ctx echo.Context) error {
	airframeService := services.AirframeService{
		ScrapeLogics:           &common.ScrapeLogicsImp{},
		FileLogics:             &common.FileLogicsImp{},
		AirframeRepository:     &repositories.AirframeRepositoryImp{DbConfig: &config.DbConfigImp{}},
		TitleOfWorkRepository:  &repositories.TitleOfWorkRepositoryImp{DbConfig: &config.DbConfigImp{}},
		PilotRepository:        &repositories.PilotRepositoryImp{DbConfig: &config.DbConfigImp{}},
		AwakenTypeRepository:   &repositories.AwakenTypeRepositoryImp{DbConfig: &config.DbConfigImp{}},
		AirframeCostRepository: &repositories.AirframeCostRepositoryImp{DbConfig: &config.DbConfigImp{}},
	}

	// 機体情報一覧取得
	offset, _ := strconv.Atoi(ctx.QueryParam("offset"))
	limit, _ := strconv.Atoi(ctx.QueryParam("limit"))
	cost, _ := strconv.Atoi(ctx.QueryParam("cost"))
	airframeInfos, err := airframeService.GetAirframes(
		offset,
		limit,
		ctx.QueryParam("airframeName"),
		ctx.QueryParam("pilot"),
		cost,
		ctx.QueryParam("titleOfWork"),
		ctx.QueryParam("awakenType"),
	)

	if err != nil {
		log.Fatal(err)

		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusOK, models.GetAirframesResponse{
		Total:     int32(len(airframeInfos)),
		Airframes: airframeInfos,
	})
}
