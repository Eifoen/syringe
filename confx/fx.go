package confx

import "go.uber.org/fx"

func ProvideStaticConfig[C any](config C) fx.Option {
	return fx.Supply(
		fx.Annotate(
			NewStaticProvider(config),
			fx.As(new(Provider[C]))))
}
