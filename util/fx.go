package util

import (
	"github.com/eifoen/syringe/modlogfx"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

func BasicTestFxOption() fx.Option {
	return fx.Options(
		fx.Supply(zap.NewNop()),
		modlogfx.Module,
		fx.WithLogger(func(lp *modlogfx.LogProvider) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: lp.ModuleLogger("fx")}
		}))
}
