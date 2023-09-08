package foundation

import (
	"github.com/fwidjaya20/go-framework/config"
	ContractConfig "github.com/fwidjaya20/go-framework/contracts/config"
	"github.com/fwidjaya20/go-framework/contracts/foundation"
)

var (
	App foundation.Application
)

func init() {
	App = &Application{
		config: config.NewApplication(".env"),
	}
}

type Application struct {
	config ContractConfig.Config
}

func NewApplication() foundation.Application {
	return App
}

func (app *Application) Boot() {
	for _, it := range app.getServiceProviders() {
		it.Boot()
	}
}

func (app *Application) Config() ContractConfig.Config {
	return app.config
}

func (app *Application) getServiceProviders() []foundation.ServiceProvider {
	return app.config.Env("app.providers").([]foundation.ServiceProvider)
}
