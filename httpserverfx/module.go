package httpserverfx

import (
	"net/http"

	"go.uber.org/fx"
)

const moduleName string = "httpserver"

var Module = fx.Module(
	moduleName,
	fx.Provide(ProvideHttpServer),
	fx.Invoke(func(_ *http.Server) {}))
