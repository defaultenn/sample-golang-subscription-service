package routes

import (
	"net/http"
	"test_task/internal/controller/dto"
	"test_task/internal/controller/usecase"
	"test_task/internal/erroring"

	"github.com/gin-gonic/gin"
)

type ReadSubscriptionResult struct {
	Data *SubscriptionResultItem `json:"data"`
}

// @Summary     Retrieve подписки
// @Description Возвращает конкретную подписку
// @ID          subscription_retrieve
// @Tags  	    subscriptions
// @Accept      json
// @Produce     json
// @Param       query-params query dto.ReadSubscription true "Входные параметры"
// @Success     200 {object} ReadSubscriptionResult
// @Failure     422 {object} erroring.HTTPRequestValidationError
// @Failure     500 {object} erroring.HTTPInternalServerError
// @Router      /subscription [get]
func (router *SubscriptionRouter) ReadSubscriptionRoute(ctx *gin.Context) {
	params := &dto.ReadSubscription{}

	if err := ctx.BindQuery(params); err != nil {
		erroring.Handle(ctx, err)
		return
	}

	sub, err := usecase.ReadSubscription(
		router.Resources.GetDatabase(),
		params,
	)

	if err != nil {
		erroring.Handle(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, sub)
}
