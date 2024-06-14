package httpserverfx

import (
	"net/http"
	"testing"

	"github.com/eifoen/syringe/confx"
	"github.com/eifoen/syringe/util"
	"github.com/stretchr/testify/assert"
	"go.uber.org/fx"
)

func TestSimpleHttpServer(t *testing.T) {
	app := fx.New(
		util.BasicTestFxOption(),
		fx.Supply(fx.Annotate(http.NewServeMux(), fx.As(new(http.Handler)))),
		Module,
	)

	assert.NoError(t, app.Err())
}

func TestSimpleHttpServerWithC(t *testing.T) {
	srv := &http.Server{
		Addr: "test",
	}
	mux := http.NewServeMux()

	app := fx.New(
		util.BasicTestFxOption(),
		fx.Supply(fx.Annotate(mux, fx.As(new(http.Handler)))),
		confx.ProvideStaticConfig(&ServerConfig{
			Server: srv,
		}),
		Module,
	)

	assert.NoError(t, app.Err())
	assert.Equal(t, mux, srv.Handler)
}
