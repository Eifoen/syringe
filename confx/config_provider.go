package confx

type Provider[C any] interface {
	Get() C
}

type StaticProvider[C any] struct {
	config C
}

func (p *StaticProvider[C]) Get() C {
	return p.config
}

func NewStaticProvider[C any](config C) *StaticProvider[C] {
	return &StaticProvider[C]{
		config: config,
	}
}
