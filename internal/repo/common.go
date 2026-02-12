package repo

import (
	"fmt"
	"reflect"
	"test_task/internal/constants"
	"test_task/internal/entity"

	"gorm.io/gorm"
)

func PaginateScope(page int) func(tx *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		if page == 0 {
			page = 1
		}
		offset := (page - 1) * constants.PageSize
		return tx.Offset(offset).Limit(constants.PageSize)
	}
}

func FindByIDScope(id uint) func(*gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		return tx.Where("id = ?", id)
	}
}

func DefaultScope(db *gorm.DB) *gorm.DB {
	return db
}

func OmitWhereScope(field string, op string, value any) func(*gorm.DB) *gorm.DB {
	if reflect.ValueOf(value).IsZero() || field == "" {
		return DefaultScope
	}

	return func(tx *gorm.DB) *gorm.DB {
		return tx.Where(fmt.Sprintf("%s %s ?", field, op), value)
	}
}

func SubscriptionListFilter(db *gorm.DB, params IListSubscription) *gorm.DB {
	return db.Model(&entity.Subscription{}).Scopes(
		OmitWhereScope("user_id", "=", params.GetUserID()),
		OmitWhereScope("service_name", "=", params.GetServiceName()),
		OmitWhereScope("start_date", ">=", params.GetStartDate()),
		OmitWhereScope("end_date", "<=", params.GetEndDate()),
	)
}
