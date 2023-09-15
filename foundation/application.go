package foundation

import (
	"fmt"
	SysLog "log"
	"os"
	"sync"

	"github.com/fwidjaya20/go-framework/config"
	"github.com/fwidjaya20/go-framework/console"
	ContractConfig "github.com/fwidjaya20/go-framework/contracts/config"
	ContractConsole "github.com/fwidjaya20/go-framework/contracts/console"
	"github.com/fwidjaya20/go-framework/contracts/foundation"
	ContractLog "github.com/fwidjaya20/go-framework/contracts/log"
	ContractSchedule "github.com/fwidjaya20/go-framework/contracts/schedule"
	"github.com/fwidjaya20/go-framework/log"
	"github.com/fwidjaya20/go-framework/schedule"
	"github.com/golang-module/carbon/v2"
)

var (
	App foundation.Application
)

type _Application struct {
	bindings  sync.Map
	instances sync.Map
}

func init() {
	app := &_Application{}

	baseServiceProvider := app.getBaseServiceProvider()

	app.registerServiceProviders(baseServiceProvider)
	app.bootServiceProviders(baseServiceProvider)

	App = app
}

func NewApplication() foundation.Application {
	return App
}

func (app *_Application) Boot() {
	configuredServiceProviders := app.GetConfig().Get("app.providers").([]foundation.ServiceProvider)
	configuredTz := app.GetConfig().GetString("app.timezone", carbon.UTC)

	app.registerServiceProviders(configuredServiceProviders)
	app.bootServiceProviders(configuredServiceProviders)

	app.GetConsole().Run(os.Args, true)
	carbon.SetTimezone(configuredTz)
}

func (app *_Application) Get(key any) (any, error) {
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

func (app *_Application) GetConfig() ContractConfig.Config {
	instance, err := app.Get(config.Binding)
	if nil != err {
		SysLog.Fatalln(err.Error())
		return nil
	}
	return instance.(ContractConfig.Config)
}

func (app *_Application) GetConsole() ContractConsole.Console {
	instance, err := app.Get(console.Binding)
	if nil != err {
		SysLog.Fatalln(err.Error())
		return nil
	}
	return instance.(ContractConsole.Console)
}

func (app *_Application) GetLogger() ContractLog.Logger {
	instance, err := app.Get(log.Binding)
	if nil != err {
		SysLog.Fatalln(err.Error())
		return nil
	}
	return instance.(ContractLog.Logger)
}

func (app *_Application) GetSchedule() ContractSchedule.Schedule {
	instance, err := app.Get(schedule.Binding)
	if nil != err {
		SysLog.Fatalln(err.Error())
		return nil
	}
	return instance.(ContractSchedule.Schedule)
}

func (app *_Application) Singleton(key any, callback func(app foundation.Application) (any, error)) {
	app.bindings.Store(key, callback)
}

func (app *_Application) bootServiceProviders(providers []foundation.ServiceProvider) {
	for _, it := range providers {
		it.Boot(app)
	}
}

func (app *_Application) getBaseServiceProvider() []foundation.ServiceProvider {
	return []foundation.ServiceProvider{
		&config.ServiceProvider{},
		&console.ServiceProvider{},
	}
}

func (app *_Application) registerServiceProviders(providers []foundation.ServiceProvider) {
	for _, it := range providers {
		it.Register(app)
	}
}
