package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"test_task/internal/config"
	"test_task/internal/controller"
	"test_task/internal/controller/middlewares"
	"test_task/internal/controller/routes"
	"test_task/internal/storage"
	"test_task/pkg/httpserver"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

func main() {
	cfg := config.NewConfig()

	zerolog.SetGlobalLevel(cfg.GetLogConfig().GetLevel())
	logger := zerolog.New(os.Stdout).
		With().
		Caller().
		Timestamp().
		Logger()
	zerolog.DefaultContextLogger = &logger
	ctx := logger.WithContext(context.Background())

	resources := storage.NewStorage(cfg.GetDatabaseConfig())

	// if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
	// 	v.RegisterValidation("month_year", func(fl validator.FieldLevel) bool {
	// 		value := fl.Field().Int()

	// 		if value == 0 {
	// 			return true
	// 		}
	// 		// _, err := time.Parse("01-2006", value)
	// 		fmt.Printf("%s\n", value)
	// 		return false
	// 	})
	// }

	engine := gin.New()
	engine.NoRoute(routes.NoRouteHandler)
	engine.NoMethod(routes.NoMethodHandler)
	engine.HandleMethodNotAllowed = true
	engine.Use(middlewares.RequestID())
	engine.Use(middlewares.Logger(ctx))
	engine.Use(gin.Recovery())
	engine.Use(cors.New(cors.Config{
		AllowOrigins:  cfg.GetHTTPConfig().GetCORSOrigins(),
		AllowHeaders:  []string{"Origin", "Authorization", "Accept", "User-Agent", "Content-Type"},
		AllowMethods:  []string{"GET", "POST", "PATCH", "DELETE"},
		ExposeHeaders: []string{"Content-Length", "Date", "Content-Type"},
	}))

	router := routes.NewSubscriptionRouter(resources)
	router.InitRoutes(engine)

	controller.InitHealthz(engine)
	controller.InitSwagger(engine)

	httpServer := httpserver.New(
		ctx,
		engine,
		httpserver.Port(cfg.GetHTTPConfig().GetPort()),
		httpserver.ReadTimeout(cfg.GetHTTPConfig().GetReadTimeout()),
		httpserver.WriteTimeout(cfg.GetHTTPConfig().GetWriteTimeout()),
		httpserver.ShutdownTimeout(cfg.GetHTTPConfig().GetShutdownTimeout()),
	)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	select {
	case signal := <-interrupt:
		logger.Debug().Msg("signal: " + signal.String())
	case err := <-httpServer.Notify():
		logger.Debug().Msg(fmt.Sprintf("httpServer.Notify: %s", err.Error()))
	}

	if err := httpServer.Shutdown(ctx); err != nil {
		logger.Debug().Msg(fmt.Sprintf("httpServer.Shutdown: %s", err.Error()))
	}

	logger.Debug().Msg("Server stops gracefully.")
}
