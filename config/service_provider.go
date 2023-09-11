package config

import (
	"github.com/fwidjaya20/go-framework/contracts/foundation"
)

const Binding = "go_framework.config"

type ServiceProvider struct{}

func (c *ServiceProvider) Boot(app foundation.Application) {}

func (c *ServiceProvider) Register(app foundation.Application) {
	app.Singleton(Binding, func(app foundation.Application) (any, error) {
		return NewApplication(".env"), nil
	})
}
