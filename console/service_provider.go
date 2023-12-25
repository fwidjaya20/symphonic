package console

import (
	"github.com/fwidjaya20/symphonic/contracts/foundation"
)

const Binding = "symphonic.console"

type ServiceProvider struct{}

func (provider *ServiceProvider) Boot(_ foundation.Application) {}

func (provider *ServiceProvider) Register(app foundation.Application) {
	app.Instance(Binding, func(app foundation.Application) (any, error) {
		return NewApplication(), nil
	})
}
