package repo

import (
	"test_task/internal/entity"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IGetSubscription interface {
	GetSubscriptionID() uint
}

func GetSubscription(db *gorm.DB, params IGetSubscription) (*entity.Subscription, error) {

	sub := &entity.Subscription{}

	if err := db.Scopes(
		FindByIDScope(params.GetSubscriptionID()),
	).First(&sub).Error; err != nil {
		return nil, err
	}

	return sub, nil
}

type ICreateSubscription interface {
	GetServiceName() string
	GetPrice() uint
	GetUserID() uuid.UUID
	GetStartDate() time.Time
	GetEndDate() *time.Time
}

func CreateSubscription(db *gorm.DB, params ICreateSubscription) (*entity.Subscription, error) {
	sub := &entity.Subscription{
		ServiceName: params.GetServiceName(),
		Price:       params.GetPrice(),
		UserID:      params.GetUserID(),
		StartDate:   params.GetStartDate(),
		EndDate:     params.GetEndDate(),
	}

	if err := db.Save(sub).Error; err != nil {
		return nil, err
	}

	return sub, nil
}

type IListSubscription interface {
	GetServiceName() string
	GetUserID() *uuid.UUID
	GetStartDate() *time.Time
	GetEndDate() *time.Time
}

type IPaginatable interface {
	GetPage() int
}

func ListSubscriptions(
	db *gorm.DB,
	params interface {
		IListSubscription
		IPaginatable
	},
) ([]*entity.Subscription, error) {

	var subs []*entity.Subscription

	if err := SubscriptionListFilter(db, params).Scopes(
		PaginateScope(params.GetPage()),
	).Find(&subs).Error; err != nil {
		return nil, err
	}

	return subs, nil
}

type IDeleteSubscription interface {
	GetSubscriptionID() uint
}

func DeleteSubscription(db *gorm.DB, params IDeleteSubscription) error {
	return db.Scopes(
		FindByIDScope(params.GetSubscriptionID()),
	).Delete(&entity.Subscription{}).Error
}

type IUpdateSubscription interface {
	GetSubscriptionID() uint
	GetStartDate() *time.Time
	GetEndDate() *time.Time
}

func UpdateSubscription(db *gorm.DB, params IUpdateSubscription) (*entity.Subscription, error) {
	sub := &entity.Subscription{
		Model: gorm.Model{
			ID: params.GetSubscriptionID(),
		},
	}

	sub.SetEndDate(params.GetEndDate())
	if startDate := params.GetStartDate(); startDate != nil {
		sub.SetStartDate(*params.GetStartDate())
	}

	if err := db.Save(sub).Error; err != nil {
		return nil, err
	}

	return sub, nil
}

type IOverralPriceSum interface {
	IListSubscription
}

func OverralPriceSum(db *gorm.DB, params IOverralPriceSum) (uint, error) {
	var sum uint

	if err := SubscriptionListFilter(db, params).Select("COALESCE(SUM(price), 0)").Scan(&sum).Error; err != nil {
		return 0, err
	}

	return sum, nil
}
