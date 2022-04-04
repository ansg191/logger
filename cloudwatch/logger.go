package cloudwatch

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/bahadirbb/zapcloudwatch"
	"go.uber.org/zap"
)

func SetupCloudwatchLogger(group, stream string, opts ...Options) (*zap.Logger, error) {
	options := &OptionsStruct{
		name:    "",
		creds:   nil,
		region:  "",
		isAsync: true,
		level:   zap.DebugLevel,
	}
	for _, opt := range opts {
		opt(options)
	}

	cfg := aws.NewConfig().
		WithRegion(options.region).
		WithCredentials(options.creds)

	cwHook, err := zapcloudwatch.NewCloudwatchHook(
		group,
		stream,
		options.isAsync,
		cfg,
		options.level,
	).GetHook()
	if err != nil {
		return nil, err
	}

	config := zap.NewDevelopmentConfig()
	config.Encoding = "json"
	logger, _ := config.Build()
	logger = logger.WithOptions(zap.Hooks(cwHook)).Named(options.name)
	return logger, nil
}
