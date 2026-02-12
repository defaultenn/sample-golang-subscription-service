package routes

import (
	"net/http"
	"test_task/internal/controller/dto"
	"test_task/internal/controller/usecase"
	"test_task/internal/erroring"

	"github.com/gin-gonic/gin"
)

// @Summary     Выдача списка подписок пользователя
// @Description Выдает список подписок пользователя с пагинацией, количество элементов на странице - 20
// @ID          subscription_list
// @Tags  	    subscriptions
// @Accept      json
// @Produce     json
// @Param       query-params query dto.ListSubscriptions true "Входные параметры"
// @Success     200 {object} dto.ListSubscriptionsResult
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

	result := &dto.ListSubscriptionsResult{
		Data: make([]*dto.SubscriptionResultItem, 0, len(subs)),
	}

	for _, sub := range subs {
		subItem := dto.SubscriptionResultItem{
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
