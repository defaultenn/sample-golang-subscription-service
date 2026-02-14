package routes

import (
	"net/http"
	"test_task/internal/controller/dto"
	"test_task/internal/controller/usecase"
	"test_task/internal/erroring"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type SubscriptionResultItem struct {
	ID          uint           `json:"id"`
	ServiceName string         `json:"service_name"`
	UserID      uuid.UUID      `json:"user_id"`
	Price       uint           `json:"price"`
	StartDate   dto.MonthYear  `json:"start_date"`
	EndDate     *dto.MonthYear `json:"end_date"`
}

type ListSubscriptionsResult struct {
	Data []*SubscriptionResultItem `json:"data"`
}

// @Summary     Выдача списка подписок пользователя
// @Description Выдает список подписок пользователя с пагинацией, количество элементов на странице - 20
// @ID          subscription_list
// @Tags  	    subscriptions
// @Accept      json
// @Produce     json
// @Param       query-params query dto.ListSubscriptions true "Входные параметры"
// @Success     200 {object} ListSubscriptionsResult
// @Failure     422 {object} erroring.HTTPRequestValidationError
// @Failure     500 {object} erroring.HTTPInternalServerError
// @Router      /subscriptions [get]
func (router *SubscriptionRouter) ListSubscriptionsRoute(ctx *gin.Context) {

	params := &dto.ListSubscriptions{}

	if err := ctx.BindQuery(params); err != nil {
		erroring.Handle(ctx, err)
		return
	}

	subs, err := usecase.ListSubscriptions(
		router.Resources.GetDatabase(),
		params,
	)

	if err != nil {
		erroring.Handle(ctx, err)
		return
	}

	result := &ListSubscriptionsResult{
		Data: make([]*SubscriptionResultItem, 0, len(subs)),
	}

	for _, sub := range subs {
		subItem := SubscriptionResultItem{
			ID:          sub.ID,
			ServiceName: sub.ServiceName,
			Price:       sub.Price,
			UserID:      sub.UserID,
			StartDate:   dto.MonthYear(sub.StartDate),
		}

		if sub.EndDate != nil {
			endDate := dto.MonthYear(*sub.EndDate)
			subItem.EndDate = &endDate
		}

		result.Data = append(result.Data, &subItem)
	}

	ctx.JSON(http.StatusOK, result)
}
