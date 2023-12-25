package config

import (
	"github.com/fwidjaya20/symphonic/contracts/foundation"
)

const Binding = "symphonic.config"

type ServiceProvider struct{}

func (c *ServiceProvider) Boot(_ foundation.Application) {}

func (c *ServiceProvider) Register(app foundation.Application) {
	app.Instance(Binding, func(app foundation.Application) (any, error) {
		return NewApplication(".env"), nil
	})
}
