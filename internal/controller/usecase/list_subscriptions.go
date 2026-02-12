package usecase

import (
	"test_task/internal/entity"
	"test_task/internal/repo"

	"gorm.io/gorm"
)

type IListSubscriptionParams interface {
	repo.IListSubscription
	repo.IPaginatable
}

func ListSubscriptions(
	db *gorm.DB,
	params IListSubscriptionParams,
) (ss []*entity.Subscription, err error) {

	// Место для бизнес правил

	ss, err = repo.ListSubscriptions(db, params)
	return
}
