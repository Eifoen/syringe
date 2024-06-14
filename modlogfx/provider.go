package modlogfx

import "go.uber.org/zap"

const moduleKey = "module"

type LogProvider struct {
	rootLogger *zap.Logger
}

func (p *LogProvider) Logger() *zap.Logger {
	return p.rootLogger
}

func (p *LogProvider) ModuleLogger(module string) *zap.Logger {
	if p.rootLogger == nil {
		return nil
	}

	return p.rootLogger.With(zap.String(moduleKey, module))
}
