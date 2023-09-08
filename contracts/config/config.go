package config

type Config interface {
	Add(name string, configuration any)
	Env(name string, defaultValue ...any) any
	Inspect() any
}
