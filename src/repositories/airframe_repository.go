package repositories

import "github.com/GIT_USER_ID/GIT_REPO_ID/src/models"

type AirframeRepository interface {
	Create(models.Airframe)
}
