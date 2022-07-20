package repositories

import "github.com/GIT_USER_ID/GIT_REPO_ID/src/models"

type AirframeCostRepository interface {
	Find(id int, cost int) models.AirframeCost
}
