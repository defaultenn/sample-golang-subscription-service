package storage

import (
	"context"
	"test_task/internal/config"
	"test_task/internal/entity"
	"time"

	"github.com/rs/zerolog"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	ginLogger "gorm.io/gorm/logger"
)

type GormLogger struct {
	ctx context.Context
}

// LogMode реализует интерфейс logger.Interface
func (l *GormLogger) LogMode(level ginLogger.LogLevel) ginLogger.Interface {
	switch level {
	case ginLogger.Error:
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	case ginLogger.Warn:
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	case ginLogger.Info:
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case ginLogger.Silent:
		zerolog.SetGlobalLevel(zerolog.NoLevel)
	}
	return l
}

// Info логирование информационных сообщений
func (l *GormLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	zerolog.Ctx(ctx).Info().Msgf(msg, data...)
}

// Warn логирование предупреждений
func (l *GormLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	zerolog.Ctx(ctx).Warn().Msgf(msg, data...)
}

// Error логирование ошибок
func (l *GormLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	zerolog.Ctx(ctx).Error().Msgf(msg, data...)
}

// Trace логирование SQL запросов
func (l *GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	elapsed := time.Since(begin)

	sql, rows := fc()

	if err != nil {
		zerolog.Ctx(ctx).Error().Fields(
			map[string]any{
				"error":   err.Error(),
				"query":   sql,
				"rows":    rows,
				"latency": elapsed.String(),
				"source":  "database",
			},
		).Msg("bad database query")
	} else {
		zerolog.Ctx(ctx).Info().Fields(
			map[string]any{
				"source":  "database",
				"rows":    rows,
				"query":   sql,
				"latency": elapsed.String(),
			},
		).Msg("success database query")
	}
}

func NewDatabase(
	cfg config.DatabaseConfigInterface,
) (*gorm.DB, error) {

	db, err := gorm.Open(
		postgres.Open(cfg.GetDatabaseDSN()),
		&gorm.Config{
			Logger:         &GormLogger{},
			TranslateError: false,
			NowFunc: func() time.Time {
				return time.Now().Local()
			},
			FullSaveAssociations: false,
		},
	)

	if err != nil {
		return nil, err
	}

	// Достаточно ли таких миграций? Впрочем добавлю еще в виде файла.
	if err := db.AutoMigrate(
		&entity.Subscription{},
	); err != nil {
		panic(err)
	}

	return db, nil
}
