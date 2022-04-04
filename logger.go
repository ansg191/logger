package logger

import "go.uber.org/zap"

func SetupLogger() (*zap.Logger, error) {
	var logger *zap.Logger
	var err error
	if Production {
		logger, err = zap.NewProduction()
	} else {
		logger, err = zap.NewDevelopment()
	}
	if err != nil {
		return nil, err
	}

	_ = zap.ReplaceGlobals(logger)

	return logger, nil
}
