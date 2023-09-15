package log

type Logger interface {
	Debug(args ...any)
	Debugf(format string, args ...any)
	Error(args ...any)
	Errorf(format string, args ...any)
	Fatal(args ...any)
	Fatalf(format string, args ...any)
	Info(args ...any)
	Infof(format string, args ...any)
	Panic(args ...any)
	Panicf(format string, args ...any)
	Warning(args ...any)
	Warningf(format string, args ...any)
}
