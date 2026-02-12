package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const (
	RequestIDHeaderName = "X-Request-ID"
)

func RequestID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		requestID := ctx.GetHeader(RequestIDHeaderName)

		if requestID == "" {
			requestID = uuid.New().String()
		}

		ctx.Set(RequestIDHeaderName, requestID)
		ctx.Writer.Header().Set("X-Request-ID", requestID)

		ctx.Next()
	}
}
