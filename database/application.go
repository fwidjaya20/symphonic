package database

import (
	"github.com/fwidjaya20/symphonic/contracts/config"
	"github.com/fwidjaya20/symphonic/contracts/database"
	"github.com/fwidjaya20/symphonic/database/driver"
	"gorm.io/gorm"
)

type Application struct {
	instance *gorm.DB
}

func NewApplication(config config.Config) database.Database {
	dbDriver := driver.GetDatabaseDriver(config)

	orm := dbDriver.GetInstance()

	return &Application{
		instance: orm,
	}
}

func (a *Application) GetSession() *gorm.DB {
	return a.instance
}
