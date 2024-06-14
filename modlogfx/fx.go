package modlogfx

import (
	"go.uber.org/zap"
)

func ProvideLogProvider(logger *zap.Logger) *LogProvider {
	return &LogProvider{
		rootLogger: logger,
	}
}
