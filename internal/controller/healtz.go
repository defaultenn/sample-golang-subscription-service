package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitHealthz(engine *gin.Engine) {
	engine.GET("/healthz", func(c *gin.Context) { c.JSON(http.StatusOK, map[string]interface{}{"status": "ok"}) })
}
