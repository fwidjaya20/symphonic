package log

import (
	"io"

	"github.com/labstack/gommon/log"
	"github.com/sirupsen/logrus"
)

type Logger interface {
	Formatter() logrus.Formatter
	SetFormatter(formatter logrus.Formatter)
	Level() log.Lvl
	SetLevel(v log.Lvl)
	Output() io.Writer
	SetOutput(w io.Writer)
	Prefix() string
	SetPrefix(p string)
	SetHeader(h string)
	WithFields(fields logrus.Fields) *logrus.Entry
	Print(i ...interface{})
	Printf(format string, args ...interface{})
	Printj(j log.JSON)
	Debug(i ...interface{})
	Debugf(format string, args ...interface{})
	Debugj(j log.JSON)
	Info(i ...interface{})
	Infof(format string, args ...interface{})
	Infoj(j log.JSON)
	Warn(i ...interface{})
	Warnf(format string, args ...interface{})
	Warnj(j log.JSON)
	Error(i ...interface{})
	Errorf(format string, args ...interface{})
	Errorj(j log.JSON)
	Fatal(i ...interface{})
	Fatalj(j log.JSON)
	Fatalf(format string, args ...interface{})
	Panic(i ...interface{})
	Panicj(j log.JSON)
	Panicf(format string, args ...interface{})
}
