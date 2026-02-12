package usecase

import (
	"test_task/internal/repo"

	"gorm.io/gorm"
)

type IUpdateSubscription interface {
	repo.IUpdateSubscription
}

func UpdateSubscription(
	db *gorm.DB,
	params IUpdateSubscription,
) (err error) {

	// Место для бизнес правил

	_, err = repo.UpdateSubscription(db, params)
	return nil
}
