package routes

import (
	"net/http"
	"test_task/internal/storage"

	"github.com/gin-gonic/gin"
)

type NoRouteResponse struct {
	Message string `json:"message" binding:"required"`
}

func NoRouteHandler(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusNotFound, &NoRouteResponse{
		Message: "такого пути не найдено",
	})
}

type NoMethodResponse struct {
	Message string `json:"message" binding:"required"`
}

func NoMethodHandler(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusMethodNotAllowed, &NoMethodResponse{
		Message: "метод для этого пути недоступен",
	})
}

type SubscriptionRouter struct {
	Resources storage.StorageInterface
}

func NewSubscriptionRouter(resources storage.StorageInterface) *SubscriptionRouter {
	return &SubscriptionRouter{
		Resources: resources,
	}
}

func (router *SubscriptionRouter) InitRoutes(engine *gin.Engine) {
	group := engine.Group("/subscription")
	{
		group.POST("", router.CreateSubscriptionRoute)
		group.GET("", router.ReadSubscriptionRoute)
		group.PATCH("", router.UpdateSubscriptionRoute)
		group.DELETE("", router.DeleteSubscriptionRoute)
	}

	group = engine.Group("/subscriptions")
	{
		group.GET("", router.ListSubscriptionsRoute)
		group.GET("/overral_sum", router.SumSubscriptionPrices)
	}
}
