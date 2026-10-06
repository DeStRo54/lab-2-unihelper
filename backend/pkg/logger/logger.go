package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"homewormanager/internal/config"
)

func InitLogger(cfg *config.Config) *zap.SugaredLogger {
	config := zap.Config{
		Level:            zap.NewAtomicLevelAt(zap.InfoLevel),
		Development:      false,
		Encoding:         cfg.EncodingMode,
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:      "time",
			LevelKey:     "level",
			MessageKey:   "msg",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeLevel:  zapcore.CapitalColorLevelEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}

	var err error
	logger, err := config.Build()
	if err != nil {
		panic(err)
	}
	defer func() {
		logger.Sync()
	}()
	logger.Info("Logger initialized")

	return logger.Sugar()
}
