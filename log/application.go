package log

import (
	"os"

	"github.com/fwidjaya20/go-framework/contracts/log"
	"github.com/sirupsen/logrus"
)

type Application struct {
	instance *logrus.Logger
}

func NewApplication() log.Logger {
	formatter := new(logrus.TextFormatter)
	formatter.TimestampFormat = "2006-01-02 15:04:05"
	formatter.FullTimestamp = true

	logrusInstance := logrus.New()
	logrusInstance.SetFormatter(formatter)
	logrusInstance.SetLevel(logrus.DebugLevel)
	logrusInstance.SetOutput(os.Stdout)
	logrusInstance.SetReportCaller(true)

	return &Application{
		instance: logrusInstance,
	}
}

func (a *Application) Debug(args ...any) {
	a.instance.Debug(args)
}

func (a *Application) Debugf(format string, args ...any) {
	a.instance.Debugf(format, args)
}

func (a *Application) Error(args ...any) {
	a.instance.Error(args)
}

func (a *Application) Errorf(format string, args ...any) {
	a.instance.Errorf(format, args)
}

func (a *Application) Fatal(args ...any) {
	a.instance.Fatal(args)
}

func (a *Application) Fatalf(format string, args ...any) {
	a.instance.Fatalf(format, args)
}

func (a *Application) Info(args ...any) {
	a.instance.Info(args)
}

func (a *Application) Infof(format string, args ...any) {
	a.instance.Infof(format, args)
}

func (a *Application) Panic(args ...any) {
	a.instance.Panic(args)
}

func (a *Application) Panicf(format string, args ...any) {
	a.instance.Panicf(format, args)
}

func (a *Application) Warning(args ...any) {
	a.instance.Warning(args)
}

func (a *Application) Warningf(format string, args ...any) {
	a.instance.Warningf(format, args)
}
