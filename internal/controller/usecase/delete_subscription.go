package usecase

import (
	"test_task/internal/repo"

	"gorm.io/gorm"
)

type IDeleteSubscription interface {
	GetSubscriptionID() uint
}

func DeleteSubscription(
	db *gorm.DB,
	params IDeleteSubscription,
) (err error) {

	// Место для бизнес правил

	err = repo.DeleteSubscription(db, params)
	return nil
}
