package config

import (
	"strings"

	"github.com/spf13/viper"
)

const (
	defaultPath      = "./config"
	defaultEnvPrefix = "GOTOOLKIT"
)

type Configer interface {
	Parse()
}

func Setup(opts ...Option) error {

	options := options{
		path:      defaultPath,
		envPrefix: defaultEnvPrefix,
	}

	for _, o := range opts {
		o.apply(&options)
	}
	if len(options.name) > 0 {
		viper.SetConfigFile(options.name)
	}

	viper.AddConfigPath(options.path)

	viper.SetEnvPrefix(options.envPrefix)
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	for k, v := range options.defaultValues {
		viper.SetDefault(k, v)
	}

	if options.flags != nil {
		viper.BindPFlags(options.flags)
	}

	if options.watch {
		viper.WatchConfig()
	}

	return viper.ReadInConfig()
}
