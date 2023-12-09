package log

import (
	LabStackLog "github.com/labstack/gommon/log"
	"github.com/sirupsen/logrus"
)

func toLogrusLevel(level LabStackLog.Lvl) logrus.Level {
	switch level {
	case LabStackLog.DEBUG:
		return logrus.DebugLevel
	case LabStackLog.INFO:
		return logrus.InfoLevel
	case LabStackLog.WARN:
		return logrus.WarnLevel
	case LabStackLog.ERROR:
		return logrus.ErrorLevel
	}
	return logrus.InfoLevel
}

func toEchoLevel(level logrus.Level) LabStackLog.Lvl {
	switch level {
	case logrus.DebugLevel:
		return LabStackLog.DEBUG
	case logrus.InfoLevel:
		return LabStackLog.INFO
	case logrus.WarnLevel:
		return LabStackLog.WARN
	case logrus.ErrorLevel:
		return LabStackLog.ERROR
	}
	return LabStackLog.OFF
}
