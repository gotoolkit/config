package config

type Configurator interface {
	Get(key string) interface{}
	GetInt(key string) int
	GetBool(key string) bool
	GetFloat64(key string) float64
	GetString(key string) string
	GetIntSlice(key string) []int
	GetStringSlice(key string) []string
	GetStringMap(key string) map[string]interface{}
}
