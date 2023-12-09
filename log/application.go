package log

import (
	"encoding/json"
	"io"
	"os"

	"github.com/fwidjaya20/symphonic/contracts/log"
	LabStackLog "github.com/labstack/gommon/log"
	"github.com/sirupsen/logrus"
)

type Application struct {
	log *logrus.Logger
}

func (a *Application) Formatter() logrus.Formatter {
	return a.log.Formatter
}

func (a *Application) SetFormatter(formatter logrus.Formatter) {
	a.log.SetFormatter(formatter)
}

func (a *Application) Level() LabStackLog.Lvl {
	return toEchoLevel(a.log.Level)
}

func (a *Application) SetLevel(v LabStackLog.Lvl) {
	a.log.SetLevel(toLogrusLevel(v))
}

func (a *Application) Output() io.Writer {
	return a.log.Out
}

func (a *Application) SetOutput(w io.Writer) {
	a.log.SetOutput(w)
}

func (a *Application) Prefix() string {
	return ""
}

func (a *Application) SetPrefix(_ string) {
	return
}

func (a *Application) SetHeader(_ string) {
	return
}

func (a *Application) WithFields(fields logrus.Fields) *logrus.Entry {
	return a.log.WithFields(fields)
}

func (a *Application) Print(i ...interface{}) {
	a.log.Print(i...)
}

func (a *Application) Printf(format string, args ...interface{}) {
	a.log.Printf(format, args...)
}

func (a *Application) Printj(j LabStackLog.JSON) {
	b, err := json.Marshal(j)
	if err != nil {
		panic(err)
	}
	a.log.Println(string(b))
}

func (a *Application) Debug(i ...interface{}) {
	a.log.Debug(i...)
}

func (a *Application) Debugf(format string, args ...interface{}) {
	a.log.Debugf(format, args...)
}

func (a *Application) Debugj(j LabStackLog.JSON) {
	b, err := json.Marshal(j)
	if err != nil {
		panic(err)
	}
	a.log.Debugln(string(b))
}

func (a *Application) Info(i ...interface{}) {
	a.log.Info(i...)
}

func (a *Application) Infof(format string, args ...interface{}) {
	a.log.Infof(format, args...)
}

func (a *Application) Infoj(j LabStackLog.JSON) {
	b, err := json.Marshal(j)
	if err != nil {
		panic(err)
	}
	a.log.Infoln(string(b))
}

func (a *Application) Warn(i ...interface{}) {
	a.log.Warn(i...)
}

func (a *Application) Warnf(format string, args ...interface{}) {
	a.log.Warnf(format, args)
}

func (a *Application) Warnj(j LabStackLog.JSON) {
	b, err := json.Marshal(j)
	if err != nil {
		panic(err)
	}
	a.log.Warnln(string(b))
}

func (a *Application) Error(i ...interface{}) {
	a.log.Error(i...)
}

func (a *Application) Errorf(format string, args ...interface{}) {
	a.log.Errorf(format, args...)
}

func (a *Application) Errorj(j LabStackLog.JSON) {
	b, err := json.Marshal(j)
	if err != nil {
		panic(err)
	}
	a.log.Errorln(string(b))
}

func (a *Application) Fatal(i ...interface{}) {
	a.log.Fatal(i...)
}

func (a *Application) Fatalf(format string, args ...interface{}) {
	a.log.Fatalf(format, args)
}

func (a *Application) Fatalj(j LabStackLog.JSON) {
	b, err := json.Marshal(j)
	if err != nil {
		panic(err)
	}
	a.log.Fatalln(string(b))
}

func (a *Application) Panic(i ...interface{}) {
	a.log.Panic(i...)
}

func (a *Application) Panicf(format string, args ...interface{}) {
	a.log.Panicf(format, args...)
}

func (a *Application) Panicj(j LabStackLog.JSON) {
	b, err := json.Marshal(j)
	if err != nil {
		panic(err)
	}
	a.log.Panicln(string(b))
}

func NewApplication() log.Logger {
	formatter := new(logrus.TextFormatter)
	formatter.TimestampFormat = "2006-01-02 15:04:05"
	formatter.FullTimestamp = true

	logger := logrus.New()
	logger.SetFormatter(formatter)
	logger.SetLevel(logrus.DebugLevel)
	logger.SetOutput(os.Stdout)
	logger.SetReportCaller(true)

	return &Application{
		log: logger,
	}
}
