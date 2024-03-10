package logging

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Log *zap.Logger = nil

func New() (*zap.Logger, error) {
	cfg := zap.Config{
		Encoding:         "json",
		OutputPaths:      []string{"stdout", "log"},
		Level:            zap.NewAtomicLevelAt(zap.DebugLevel),
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey: "msg",
			LevelKey:   "level",
			TimeKey:    "ts",
			EncodeTime: zapcore.ISO8601TimeEncoder,
		},
	}

	return cfg.Build()
}
