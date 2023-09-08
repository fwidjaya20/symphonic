package foundation

import (
	"github.com/fwidjaya20/go-framework/contracts/foundation"
	"github.com/fwidjaya20/go-framework/database"
)

type Application struct {
	services []foundation.ServiceProvider
}

func NewApplication() foundation.Application {
	return &Application{
		services: []foundation.ServiceProvider{
			database.NewDatabaseServiceProvider(),
		},
	}
}

func (app *Application) Boot() {
	for _, it := range app.services {
		it.Boot()
	}
}
