package config

import (
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

var (
	defaultStringReplacer = NewReplacer(".", "_")
)

const (
	defaultFile      = "./.env"
	defaultEnvPrefix = "GOTOOLKIT"
)

func Setup(opts ...Option) error {

	options := options{
		file:      defaultFile,
		envPrefix: defaultEnvPrefix,
		replacer:  defaultStringReplacer,
	}

	err := createDefaultConfigFile()
	if err != nil {
		return err
	}

	for _, o := range opts {
		o.apply(&options)
	}

	viper.SetConfigFile(options.file)
	if options.autoEnv {
		viper.SetEnvPrefix(options.envPrefix)
		viper.SetEnvKeyReplacer(options.replacer.Replacer)
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

func createDefaultConfigFile() error {
	dir, _ := filepath.Split(defaultFile)
	if err := os.MkdirAll(dir, 0750); err != nil {
		return err
	}

	f, err := os.OpenFile(defaultFile, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}

	defer f.Close()

	return nil
}
