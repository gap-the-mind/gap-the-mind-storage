package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// CreateLogger creates a new logger
func CreateLogger() *zap.SugaredLogger {

	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	logger, _ := config.Build()

	return logger.Sugar()
}
