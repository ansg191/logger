package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func SetupLogger(level zapcore.Level) (*zap.Logger, error) {
	var logger *zap.Logger
	var err error
	if Production {
		cfg := zap.NewProductionConfig()
		cfg.Level.SetLevel(level)
		logger, err = cfg.Build()
	} else {
		cfg := zap.NewDevelopmentConfig()
		cfg.Level.SetLevel(level)
		logger, err = cfg.Build()
	}
	if err != nil {
		return nil, err
	}

	_ = zap.ReplaceGlobals(logger)

	return logger, nil
}
