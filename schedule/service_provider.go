package schedule

import (
	"github.com/fwidjaya20/symphonic/contracts/foundation"
)

const Binding = "symphonic.schedule"

type ServiceProvider struct{}

func (sp *ServiceProvider) Boot(_ foundation.Application) {}

func (sp *ServiceProvider) Register(app foundation.Application) {
	app.Instance(Binding, func(app foundation.Application) (any, error) {
		return NewApplication(app.GetLogger()), nil
	})
}
