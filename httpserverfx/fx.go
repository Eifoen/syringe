package httpserverfx

import (
	"context"
	"net/http"

	"github.com/eifoen/syringe/confx"
	"github.com/eifoen/syringe/modlogfx"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type HttpServerModuleParameters struct {
	fx.In

	LC fx.Lifecycle

	LogProvider    *modlogfx.LogProvider
	ConfigProvider confx.Provider[*ServerConfig] `optional:"true"`
	HttpHandler    http.Handler
}

func ProvideHttpServer(params HttpServerModuleParameters) *http.Server {
	logger := params.LogProvider.ModuleLogger(moduleName)

	logger.Info("configuring server..")
	if params.ConfigProvider == nil {
		params.ConfigProvider = confx.NewStaticProvider[*ServerConfig](defaultConfig)
		logger.Info("using default configuration")
	}

	config := params.ConfigProvider.Get()

	s := config.Server
	s.Handler = params.HttpHandler

	params.LC.Append(
		fx.Hook{
			OnStart: fxStartHook(logger, s),
			OnStop:  fxStopHook(logger, s),
		})

	logger.Info("server configured")

	return s
}

func fxStartHook(logger *zap.Logger, s *http.Server) func(context.Context) error {
	return func(ctx context.Context) error {
		logger.Info("server starting...")
		go s.ListenAndServe()
		logger.Info("server running")
		return nil
	}
}

func fxStopHook(logger *zap.Logger, s *http.Server) func(context.Context) error {
	return func(ctx context.Context) error {
		logger.Info("server stopping...")
		err := s.Shutdown(ctx)
		if err != nil {
			logger.Error(
				"unable to stop server",
				zap.Error(err))
			return err
		}
		logger.Info("server stopped")
		return nil
	}
}
