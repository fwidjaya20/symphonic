package config

type Config interface {
	Add(name string, configuration any)
	Get(name string, defaultValue ...any) any
	GetString(name string, defaultValue ...string) string
	Inspect() any
}
