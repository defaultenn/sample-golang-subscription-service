package routes

import (
	"net/http"
	"test_task/internal/controller/dto"
	"test_task/internal/controller/usecase"
	"test_task/internal/erroring"

	"github.com/gin-gonic/gin"
)

// @Summary     Калькуляция суммы стоимости всех подписок
// @Description Высчитывает сумму стоимости всех подписок за выбранный период, по названию подписки для конкретного пользователя
// @ID          subscription_overral_sum
// @Tags  	    subscriptions
// @Accept      json
// @Produce     json
// @Param       query-params query dto.SumSubscriptionPricesParams true "Входные параметры"
// @Success     200 {object} dto.SumSubscriptionPricesResult
// @Failure     422 {object} erroring.HTTPRequestValidationError
// @Failure     500 {object} erroring.HTTPInternalServerError
// @Router      /subscriptions/overral_sum [get]
func (router *SubscriptionRouter) SumSubscriptionPrices(ctx *gin.Context) {
	params := &dto.SumSubscriptionPricesParams{}

	if err := ctx.BindQuery(params); err != nil {
		erroring.Handle(ctx, err)
		return
	}

	sum, err := usecase.SumSubscriptionPrices(
		router.Resources.GetDatabase(),
		params,
	)

	if err != nil {
		erroring.Handle(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, &dto.SumSubscriptionPricesResult{
		Sum: sum,
	})
}
