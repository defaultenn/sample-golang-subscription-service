package usecase

import (
	"test_task/internal/repo"

	"gorm.io/gorm"
)

type ICreateSubscriptionParams interface {
	repo.ICreateSubscription
}

func CreateSubscription(
	db *gorm.DB,
	params ICreateSubscriptionParams,
) (err error) {

	// Место для бизнес правил

	_, err = repo.CreateSubscription(db, params)
	return
}
