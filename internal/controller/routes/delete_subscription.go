package routes

import (
	"net/http"
	"test_task/internal/controller/dto"
	"test_task/internal/controller/usecase"
	"test_task/internal/erroring"

	"github.com/gin-gonic/gin"
)

// @Summary     Удаление подписки пользователя
// @Description Удаляет подписку у пользователя
// @ID          subscription_delete
// @Tags  	    subscriptions
// @Accept      json
// @Produce     json
// @Success     200
// @Param       query-params query dto.DeleteSubscription true "Входные параметры"
// @Failure     422 {object} erroring.HTTPRequestValidationError
// @Failure     500 {object} erroring.HTTPInternalServerError
// @Router      /subscription [delete]
func (router *SubscriptionRouter) DeleteSubscriptionRoute(ctx *gin.Context) {
	params := &dto.DeleteSubscription{}

	if err := ctx.BindQuery(params); err != nil {
		erroring.Handle(ctx, err)
		return
	}

	if err := usecase.DeleteSubscription(
		router.Resources.GetDatabase(),
		params,
	); err != nil {
		erroring.Handle(ctx, err)
		return
	}

	ctx.Status(http.StatusOK)
}
