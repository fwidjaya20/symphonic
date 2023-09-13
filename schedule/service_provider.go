package schedule

import (
	"github.com/fwidjaya20/go-framework/contracts/foundation"
)

const Binding = "go_framework.schedule"

type ServiceProvider struct{}

func (sp *ServiceProvider) Boot(app foundation.Application) {}

func (sp *ServiceProvider) Register(app foundation.Application) {
	app.Singleton(Binding, func(app foundation.Application) (any, error) {
		return NewApplication(), nil
	})
}
