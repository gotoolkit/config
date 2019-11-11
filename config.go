package config

import (
	"strings"

	"github.com/spf13/viper"
)

var (
	defaultStringReplacer = strings.NewReplacer(".", "_")
)

const (
	defaultPath      = "./config"
	defaultEnvPrefix = "GOTOOLKIT"
)

func Setup(opts ...Option) error {

	options := options{
		path:           defaultPath,
		envPrefix:      defaultEnvPrefix,
		stringReplacer: defaultStringReplacer,
	}

	for _, o := range opts {
		o.apply(&options)
	}
	if len(options.name) > 0 {
		viper.SetConfigFile(options.name)
	}

	viper.AddConfigPath(options.path)

	if options.autoEnv {
		viper.SetEnvPrefix(options.envPrefix)
		viper.SetEnvKeyReplacer(options.stringReplacer)
		viper.AutomaticEnv()
	}

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
