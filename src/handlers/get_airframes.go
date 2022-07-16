package handlers

import (
	"net/http"

	"github.com/GIT_USER_ID/GIT_REPO_ID/src/requestSchema"
	"github.com/labstack/echo/v4"
)

// GetAirframes - 機体情報の一覧取得
func (c *Container) GetAirframes(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, requestSchema.HelloWorld{
		Message: "Hello World",
	})
}
