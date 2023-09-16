package log

import "github.com/fwidjaya20/go-framework/contracts/foundation"

const Binding = "go_framework.log"

type ServiceProvider struct{}

func (sp *ServiceProvider) Boot(app foundation.Application) {}

func (sp *ServiceProvider) Register(app foundation.Application) {
	app.Instance(Binding, func(app foundation.Application) (any, error) {
		return NewApplication(), nil
	})
}
