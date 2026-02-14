package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Subscription struct {
	gorm.Model

	ServiceName string    `json:"service_name"`
	Price       uint      `json:"price"`
	UserID      uuid.UUID `json:"user_id"`

	StartDate time.Time  `json:"start_date"`
	EndDate   *time.Time `json:"end_date"`
}

func (s *Subscription) GetSubscriptionID() uint {
	return s.ID
}

func (s *Subscription) GetServiceName() string {
	return s.ServiceName
}

func (s *Subscription) SetServiceName(value string) {
	s.ServiceName = value
}

func (s *Subscription) GetUserID() uuid.UUID {
	return s.UserID
}

func (s *Subscription) SetUserID(value uuid.UUID) {
	s.UserID = value
}

func (s *Subscription) GetPrice() uint {
	return s.Price
}

func (s *Subscription) SetPrice(value uint) {
	s.Price = value
}

func (s *Subscription) GetStartDate() time.Time {
	return time.Time(s.StartDate)
}

func (s *Subscription) SetStartDate(value time.Time) {
	s.StartDate = value
}

func (s *Subscription) GetEndDate() *time.Time {
	return s.EndDate
}

func (s *Subscription) SetEndDate(value *time.Time) {
	s.EndDate = value
}
