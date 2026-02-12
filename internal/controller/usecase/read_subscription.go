package usecase

import (
	"test_task/internal/entity"
	"test_task/internal/repo"

	"gorm.io/gorm"
)

type IReadSubscription interface {
	repo.IGetSubscription
}

func ReadSubscription(
	db *gorm.DB,
	params IReadSubscription,
) (s *entity.Subscription, err error) {

	// Место для бизнес правил

	s, err = repo.GetSubscription(db, params)
	return
}
