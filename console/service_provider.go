package console

import (
	"github.com/fwidjaya20/go-framework/contracts/foundation"
)

const Binding = "go_framework.console"

type ServiceProvider struct{}

func (provider *ServiceProvider) Boot(app foundation.Application) {}

func (provider *ServiceProvider) Register(app foundation.Application) {
	app.Instance(Binding, func(app foundation.Application) (any, error) {
		return NewApplication(), nil
	})
}
