package config

type ConfigType string

const (
	YAML       ConfigType = "yaml"
	JSON                  = "json"
	HCL                   = "hcl"
	ENV                   = "env"
	PROPERTIES            = "properties"
)

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
