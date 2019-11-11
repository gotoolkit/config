package config

import (
	"github.com/spf13/viper"
)

const (
	defaultFileName  = "config.json"
	defaultPath      = "./"
	defaultEnvPrefix = "GOTOOLKIT_"
)

type Configer interface {
	Parse()
}

func Setup(opts ...Option) error {
	viper.SetDefault("config", "")
	options := options{
		name:      defaultFileName,
		path:      defaultPath,
		envPrefix: defaultEnvPrefix,
	}

	for _, o := range opts {
		o.apply(&options)
	}
	viper.SetConfigFile(options.name)
	viper.AddConfigPath(options.path)
	viper.SetEnvPrefix(options.envPrefix)

	for k, v := range options.defaultValues {
		viper.SetDefault(k, v)
	}

	viper.BindPFlags(options.flags)

	if options.watch {
		viper.WatchConfig()
	}

	return viper.ReadInConfig()
}
