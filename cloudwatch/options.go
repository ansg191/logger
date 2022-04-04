package cloudwatch

import (
	"github.com/aws/aws-sdk-go/aws/credentials"
	"go.uber.org/zap/zapcore"
)

type OptionsStruct struct {
	name    string
	creds   *credentials.Credentials
	region  string
	isAsync bool
	level   zapcore.Level
}

type Options func(*OptionsStruct)

func WithName(name string) Options {
	return func(opt *OptionsStruct) {
		opt.name = name
	}
}

func WithStaticCredentials(accessKey, secretKey string) Options {
	return func(opt *OptionsStruct) {
		opt.creds = credentials.NewStaticCredentials(accessKey, secretKey, "")
	}
}

func WithEnvCredentials() Options {
	return func(opt *OptionsStruct) {
		opt.creds = credentials.NewEnvCredentials()
	}
}

func WithRegion(region string) Options {
	return func(opt *OptionsStruct) {
		opt.region = region
	}
}

func withAsync(async bool) Options {
	return func(opt *OptionsStruct) {
		opt.isAsync = async
	}
}

func WithLevel(level zapcore.Level) Options {
	return func(opt *OptionsStruct) {
		opt.level = level
	}
}
