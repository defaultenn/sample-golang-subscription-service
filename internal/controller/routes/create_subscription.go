package routes

import (
	"net/http"
	"test_task/internal/controller/dto"
	"test_task/internal/controller/usecase"
	"test_task/internal/erroring"

	"github.com/gin-gonic/gin"
)

// @Summary     Создание подписки для пользователя
// @Description Создает подписку для пользователя для указанного сервиса
// @ID          subscription_create
// @Tags  	    subscriptions
// @Accept      json
// @Produce     json
// @Success     200
// @Param       json-string body dto.CreateSubscription true "Входные параметры"
// @Failure     422 {object} erroring.HTTPRequestValidationError
// @Failure     500 {object} erroring.HTTPInternalServerError
// @Router      /subscription [post]
func (router *SubscriptionRouter) CreateSubscriptionRoute(ctx *gin.Context) {
	params := &dto.CreateSubscription{}

	if err := ctx.BindJSON(params); err != nil {
		erroring.Handle(ctx, err)
		return
	}

	if err := usecase.CreateSubscription(
		router.Resources.GetDatabase(),
		params,
	); err != nil {
		erroring.Handle(ctx, err)
		return
	}

	ctx.Status(http.StatusOK)
}
