package config

import (
	"log"
	"os"

	"github.com/spf13/cast"
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

func (app *Application) Get(name string, defaultValue ...any) any {
	if app.viper.IsSet(name) {
		return app.viper.Get(name)
	}

	if len(defaultValue) > 0 {
		return defaultValue[0]
	}

	return nil
}

func (app *Application) GetInt(name string, defaultValue ...int) int {
	return cast.ToInt(app.Get(name, defaultValue))
}

func (app *Application) GetInt8(name string, defaultValue ...int8) int8 {
	return cast.ToInt8(app.Get(name, defaultValue))
}

func (app *Application) GetInt16(name string, defaultValue ...int16) int16 {
	return cast.ToInt16(app.Get(name, defaultValue))
}

func (app *Application) GetInt32(name string, defaultValue ...int32) int32 {
	return cast.ToInt32(app.Get(name, defaultValue))
}

func (app *Application) GetInt64(name string, defaultValue ...int64) int64 {
	return cast.ToInt64(app.Get(name, defaultValue))
}

func (app *Application) GetString(name string, defaultValue ...string) string {
	return cast.ToString(app.Get(name, defaultValue))
}

func (app *Application) Inspect() any {
	return app.viper.AllSettings()
}
