package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

type Application struct {
	viper *viper.Viper
}

func NewApplication(filePath string) *Application {
	if _, err := os.Stat(filePath); nil != err {
		log.Fatal("Environment file was not found.")
	}

	app := &Application{
		viper: viper.New(),
	}

	app.viper.SetConfigFile("env")
	app.viper.SetConfigFile(filePath)

	if err := app.viper.ReadInConfig(); nil != err {
		log.Fatalln(err.Error())
	}

	app.viper.AutomaticEnv()

	return app
}

func (app *Application) Add(name string, configuration any) {
	app.viper.Set(name, configuration)
}

func (app *Application) Env(name string, defaultValue ...any) any {
	if app.viper.IsSet(name) {
		return app.viper.Get(name)
	}

	if len(defaultValue) > 0 {
		return defaultValue[0]
	}

	return nil
}

func (app *Application) Inspect() any {
	return app.viper.AllSettings()
}
