package foundation

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/fwidjaya20/go-framework/config"
	"github.com/fwidjaya20/go-framework/console"
	ContractConfig "github.com/fwidjaya20/go-framework/contracts/config"
	ContractConsole "github.com/fwidjaya20/go-framework/contracts/console"
	"github.com/fwidjaya20/go-framework/contracts/foundation"
	"github.com/golang-module/carbon/v2"
)

var (
	App foundation.Application
)

type application struct {
	bindings  sync.Map
	instances sync.Map
}

func init() {
	app := &application{}

	baseServiceProvider := app.getBaseServiceProvider()

	app.registerServiceProviders(baseServiceProvider)
	app.bootServiceProviders(baseServiceProvider)

	App = app
}

func NewApplication() foundation.Application {
	return App
}

func (app *application) Boot() {
	configuredServiceProviders := app.GetConfig().Get("app.providers").([]foundation.ServiceProvider)
	configuredTz := app.GetConfig().GetString("app.timezone", carbon.UTC)

	app.registerServiceProviders(configuredServiceProviders)
	app.bootServiceProviders(configuredServiceProviders)

	app.GetConsole().Run(os.Args, true)
	carbon.SetTimezone(configuredTz)
}

func (app *application) Get(key any) (any, error) {
	binding, ok := app.bindings.Load(key)
	if !ok {
		return nil, fmt.Errorf("binding was not found: %+v", key)
	}

	if instance, ok := app.instances.Load(key); ok {
		return instance, nil
	}

	bindingImpl, err := binding.(func(app foundation.Application) (any, error))(app)
	if nil != err {
		return nil, err
	}

	app.instances.Store(key, bindingImpl)

	return bindingImpl, nil
}

func (app *application) GetConfig() ContractConfig.Config {
	instance, err := app.Get(config.Binding)
	if nil != err {
		log.Fatalln(err.Error())
		return nil
	}
	return instance.(ContractConfig.Config)
}

func (app *application) GetConsole() ContractConsole.Console {
	instance, err := app.Get(console.Binding)
	if nil != err {
		log.Fatalln(err.Error())
		return nil
	}
	return instance.(ContractConsole.Console)
}

func (app *application) Singleton(key any, callback func(app foundation.Application) (any, error)) {
	app.bindings.Store(key, callback)
}

func (app *application) bootServiceProviders(providers []foundation.ServiceProvider) {
	for _, it := range providers {
		it.Boot(app)
	}
}

func (app *application) getBaseServiceProvider() []foundation.ServiceProvider {
	return []foundation.ServiceProvider{
		&config.ServiceProvider{},
		&console.ServiceProvider{},
	}
}

func (app *application) registerServiceProviders(providers []foundation.ServiceProvider) {
	for _, it := range providers {
		it.Register(app)
	}
}
