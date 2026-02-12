package dto

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
)

type MonthYear time.Time

func (ct *MonthYear) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	if s == "" || s == "null" {
		return nil
	}

	t, err := time.Parse("01-2006", s)
	if err == nil {
		*ct = MonthYear(t)
		return nil
	}

	return fmt.Errorf("invalid time format: %s", s)
}

func (ct MonthYear) MarshalJSON() ([]byte, error) {
	t := time.Time(ct)
	return []byte(`"` + t.Format("01-2006") + `"`), nil
}

type CreateSubscription struct {
	ServiceName string     `json:"service_name" binding:"required"`
	Price       uint       `json:"price" binding:"required"`
	UserID      uuid.UUID  `json:"user_id" binding:"requred,uuid4"`
	StartDate   MonthYear  `json:"start_date" binding:"required"`
	EndDate     *MonthYear `json:"end_date" binding:"omitempty"`
}

func (cs *CreateSubscription) GetServiceName() string {
	return cs.ServiceName
}

func (cs *CreateSubscription) GetPrice() uint {
	return cs.Price
}

func (cs *CreateSubscription) GetUserID() uuid.UUID {
	return cs.UserID
}

func (cs *CreateSubscription) GetStartDate() time.Time {
	return time.Time(cs.StartDate)
}

func (cs *CreateSubscription) GetEndDate() *time.Time {
	if cs.EndDate == nil {
		return nil
	}
	t := time.Time(*cs.EndDate)
	if t.IsZero() {
		return nil
	}
	return &t
}

type ReadSubscription struct {
	ID uint `json:"id" form:"id" binding:"gt=0"`
}

func (rs *ReadSubscription) GetSubscriptionID() uint {
	return rs.ID
}

type ReadSubscriptionResult struct {
	Data *SubscriptionResultItem `json:"data"`
}

type UpdateSubscription struct {
	ID        uint       `json:"id" binding:"gt=0"`
	StartDate *MonthYear `json:"start_date" binding:"omitempty"`
	EndDate   *MonthYear `json:"end_date" binding:"omitempty"`
}

func (us *UpdateSubscription) GetSubscriptionID() uint {
	return us.ID
}

func (us *UpdateSubscription) GetStartDate() *time.Time {
	if us.StartDate == nil {
		return nil
	}
	t := time.Time(*us.StartDate)
	if t.IsZero() {
		return nil
	}
	return &t
}

func (us *UpdateSubscription) GetEndDate() *time.Time {
	if us.EndDate == nil {
		return nil
	}
	t := time.Time(*us.EndDate)
	if t.IsZero() {
		return nil
	}
	return &t
}

type DeleteSubscription struct {
	ID uint `json:"id" form:"id" binding:"required,gt=0"`
}

func (ds *DeleteSubscription) GetSubscriptionID() uint {
	return ds.ID
}

type SubscriptionFilters struct {
	UserID      string     `json:"user_id" form:"user_id" binding:"omitempty,uuid4"`
	ServiceName string     `json:"service_name" form:"service_name" binding:"omitempty"`
	StartDate   *time.Time `json:"start_date" form:"start_date" binding:"omitempty" time_format:"01-2006"`
	EndDate     *time.Time `json:"end_date" form:"end_date" binding:"omitempty" time_format:"01-2006"`
}

func (sf *SubscriptionFilters) GetUserID() *uuid.UUID {

	id, err := uuid.Parse(sf.UserID)

	if err != nil {
		return nil
	}

	return &id
}

func (sf *SubscriptionFilters) GetServiceName() string {
	return sf.ServiceName
}

func (sf *SubscriptionFilters) GetStartDate() *time.Time {
	return sf.StartDate
}

func (sf *SubscriptionFilters) GetEndDate() *time.Time {
	return sf.EndDate
}

type ListSubscriptions struct {
	SubscriptionFilters
	Page int `json:"page" form:"page" binding:"omitempty,gt=0"`
}

func (ls *ListSubscriptions) GetPage() int {
	return ls.Page
}

type SubscriptionResultItem struct {
	ID          uint       `json:"id"`
	ServiceName string     `json:"service_name"`
	UserID      uuid.UUID  `json:"user_id"`
	Price       uint       `json:"price"`
	StartDate   MonthYear  `json:"start_date"`
	EndDate     *MonthYear `json:"end_date"`
}

type ListSubscriptionsResult struct {
	Data []*SubscriptionResultItem `json:"data"`
}

type SumSubscriptionPricesParams struct {
	SubscriptionFilters
}

type SumSubscriptionPricesResult struct {
	Sum uint `json:"sum"`
}
