package lib

import (
	"context"
	"os"

	"github.com/google/uuid"

	"github.com/rs/zerolog"
)

var Logger *zerolog.Logger

func NewLogger(isDebug bool) {
	logLevel := zerolog.InfoLevel
	if isDebug {
		logLevel = zerolog.DebugLevel
	}

	zerolog.SetGlobalLevel(logLevel)
	zerolog.LevelFieldName = "level"
	zerolog.TimestampFieldName = "timestamp"
	zerolog.MessageFieldName = "message"

	l := zerolog.New(os.Stderr).
		Hook(logIDHook{}).
		With().
		Timestamp().
		Caller().Logger()
	Logger = &l
}

// unique Log ID
type logIDHook struct{}

func (h logIDHook) Run(e *zerolog.Event, level zerolog.Level, msg string) {
	e.Str("id", uuid.New().String())
}

func LogDebug(ctx context.Context) *zerolog.Event {
	return logWithContext(ctx, Logger.Debug())
}

func LogInfo(ctx context.Context) *zerolog.Event {
	return logWithContext(ctx, Logger.Info())
}

func LogWarn(ctx context.Context) *zerolog.Event {
	return logWithContext(ctx, Logger.Warn())
}

func LogError(ctx context.Context, err error) *zerolog.Event {
	return logWithContext(ctx, Logger.Error().Err(err))
}

func logWithContext(ctx context.Context, evt *zerolog.Event) *zerolog.Event {
	if ctx == nil {
		return evt
	}

	getCtxValue := func(ctx context.Context, key string) string {
		if v := ctx.Value(key); v != nil {
			return v.(string)
		}
		return ""
	}

	return evt.
		Str("req_id", getCtxValue(ctx, "req_id")).
		Str("tenant_id", getCtxValue(ctx, "tenant_id")).
		Str("app_name", getCtxValue(ctx, "app_name")).
		Str("app_id", getCtxValue(ctx, "app_id"))
}
