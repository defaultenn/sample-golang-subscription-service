package usecase

import (
	"test_task/internal/entity"
	"test_task/internal/repo"

	"gorm.io/gorm"
)

type ICreateSubscriptionParams interface {
	repo.ICreateSubscription
}

func CreateSubscription(
	db *gorm.DB,
	params ICreateSubscriptionParams,
) (*entity.Subscription, error) {

	// Место для бизнес правил

	return repo.CreateSubscription(db, params)
}
