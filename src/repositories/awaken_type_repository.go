package repositories

import "github.com/GIT_USER_ID/GIT_REPO_ID/src/models"

type AwakenTypeRepository interface {
	Find(id int, name string) models.AwakenType
}
