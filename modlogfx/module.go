package modlogfx

import "go.uber.org/fx"

const moduleName string = "modlog"

var Module = fx.Module(
	moduleName,
	fx.Provide(ProvideLogProvider))
