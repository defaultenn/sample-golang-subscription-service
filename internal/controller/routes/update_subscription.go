package routes

import (
	"net/http"
	"test_task/internal/controller/dto"
	"test_task/internal/controller/usecase"
	"test_task/internal/erroring"

	"github.com/gin-gonic/gin"
)

// @Summary     Обновление конкретной подписки пользователя
// @Description Обновляет конкретную подписку пользователя
// @ID          subscription_update
// @Tags  	    subscriptions
// @Accept      json
// @Produce     json
// @Success     200
// @Param       json-string body dto.UpdateSubscription true "Входные параметры"
// @Failure     422 {object} erroring.HTTPRequestValidationError
// @Failure     500 {object} erroring.HTTPInternalServerError
// @Router      /subscription [patch]
func (router *SubscriptionRouter) UpdateSubscriptionRoute(ctx *gin.Context) {
	params := &dto.UpdateSubscription{}

	if err := ctx.BindJSON(params); err != nil {
		erroring.Handle(ctx, err)
		return
	}

	if err := usecase.UpdateSubscription(
		router.Resources.GetDatabase(),
		params,
	); err != nil {
		erroring.Handle(ctx, err)
		return
	}

	ctx.Status(http.StatusOK)
}
