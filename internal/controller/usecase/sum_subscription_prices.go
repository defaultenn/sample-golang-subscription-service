package usecase

import (
	"test_task/internal/repo"

	"gorm.io/gorm"
)

type ISumSubscriptionPricesParams interface {
	repo.IOverralPriceSum
}

func SumSubscriptionPrices(
	db *gorm.DB,
	params ISumSubscriptionPricesParams,
) (uint, error) {

	// Место для бизнес правил

	return repo.OverralPriceSum(db, params)
}
