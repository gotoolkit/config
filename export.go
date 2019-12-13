package config

import (
	"time"

	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
)

func YamlStringSettings() (string, error) {
	c := viper.AllSettings()
	bs, err := yaml.Marshal(c)
	if err != nil {
		return "", err
	}
	return string(bs), nil
}

func IsSet(key string) bool {
	return viper.IsSet(key)
}

func AllSettings() map[string]interface{} {
	return viper.AllSettings()
}

func Get(key string) interface{} {
	return viper.Get(key)
}

func GetString(key string) string {
	return viper.GetString(key)
}

func GetBool(key string) bool {
	return viper.GetBool(key)
}

func GetInt(key string) int {
	return viper.GetInt(key)
}

func GetInt32(key string) int32 {
	return viper.GetInt32(key)
}

func GetInt64(key string) int64 {
	return viper.GetInt64(key)
}

func GetUint(key string) uint {
	return viper.GetUint(key)
}

func GetUint32(key string) uint32 {
	return viper.GetUint32(key)
}

func GetUint64(key string) uint64 {
	return viper.GetUint64(key)
}

func GetFloat64(key string) float64 {
	return viper.GetFloat64(key)
}

func GetTime(key string) time.Time {
	return viper.GetTime(key)
}

func GetDuration(key string) time.Duration {
	return viper.GetDuration(key)
}

func GetIntSlice(key string) []int {
	return viper.GetIntSlice(key)
}

func GetStringSlice(key string) []string {
	return viper.GetStringSlice(key)
}

func GetStringMap(key string) map[string]interface{} {
	return viper.GetStringMap(key)
}

func GetStringMapString(key string) map[string]string {
	return viper.GetStringMapString(key)
}

func GetStringMapStringSlice(key string) map[string][]string {
	return viper.GetStringMapStringSlice(key)
}
