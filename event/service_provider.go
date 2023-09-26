package event

import "github.com/fwidjaya20/symphonic/contracts/foundation"

type ServiceProvider struct{}

const Binding = "symphonic.event"

func (sp *ServiceProvider) Boot(app foundation.Application) {}

func (sp *ServiceProvider) Register(app foundation.Application) {
	app.Instance(Binding, func(app foundation.Application) (any, error) {
		return NewApplication(app.GetConfig(), app.GetLogger()), nil
	})
}
