package logger

import (
	"github.com/hiagomf/bank-api/server/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// SetupLogger initialize zap logger
func SetupLogger() (*zap.Logger, error) {
	c := config.GetConfig()
	var cfg zap.Config

	if c.Production {
		cfg = zap.NewProductionConfig()
		cfg.DisableStacktrace = true
		cfg.EncoderConfig = zap.NewProductionEncoderConfig()
	} else {
		cfg = zap.NewDevelopmentConfig()
		cfg.EncoderConfig = zap.NewDevelopmentEncoderConfig()
	}

	cfg.Encoding = "json"

	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	cfg.EncoderConfig.MessageKey = "message"
	cfg.EncoderConfig.LevelKey = "level"
	cfg.EncoderConfig.CallerKey = "caller"
	cfg.EncoderConfig.NameKey = "name"
	cfg.EncoderConfig.TimeKey = "time"
	cfg.EncoderConfig.StacktraceKey = "stack_trace"
	cfg.InitialFields = map[string]interface{}{
		"application": "bank-api",
		"version":     "0.1",
	}

	cfg.OutputPaths = []string{c.AccessLogDirectory, "stdout"}
	cfg.ErrorOutputPaths = []string{c.ErrorLogDirectory, "stderr"}

	return cfg.Build()
}

// PanicRecovery handles recovered panics
func PanicRecovery(p interface{}) (err error) {
	zap.S().Error(
		"PANIC detected: ",
		p,
	)

	return
}
