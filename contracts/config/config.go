package config

type Config interface {
	Add(name string, configuration any)
	Get(name string, defaultValue ...any) any
	GetArrayString(name string, delimiter string, defaultValues ...string) []string
	GetInt(name string, defaultValue ...int) int
	GetInt8(name string, defaultValue ...int8) int8
	GetInt16(name string, defaultValue ...int16) int16
	GetInt32(name string, defaultValue ...int32) int32
	GetInt64(name string, defaultValue ...int64) int64
	GetString(name string, defaultValue ...string) string
	Inspect() any
}
