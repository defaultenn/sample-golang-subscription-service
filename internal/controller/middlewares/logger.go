package middlewares

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

func Logger(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {

		requestLogger := zerolog.Ctx(ctx).With()

		if requestID := c.GetString(RequestIDHeaderName); requestID != "" {
			requestLogger = requestLogger.Str("request_id", requestID)
		}

		finalLogger := requestLogger.Logger()

		newCtx := finalLogger.WithContext(c.Request.Context())
		c.Request = c.Request.WithContext(newCtx)

		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery
		start := time.Now()
		c.Next()

		if raw != "" {
			path = path + "?" + raw
		}

		msg := c.Errors.ByType(gin.ErrorTypePrivate).String()

		fields := map[string]any{
			"path":        path,
			"_msg":        msg,
			"body_size":   c.Writer.Size(),
			"status_code": c.Writer.Status(),
			"method":      c.Request.Method,
			"client_ip":   c.ClientIP(),
			"latency":     time.Since(start).String(),
			"user_agent":  c.Request.UserAgent(),
		}

		if msg == "" {
			msg = "success http response"
		} else {
			fields["error"] = msg
			msg = "bad http response"
		}

		finalLogger.Info().Fields(fields).Msg(msg)
	}
}
