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

	opt := options{
		file:      defaultFile,
		envPrefix: defaultEnvPrefix,
		replacer:  defaultStringReplacer,
	}

	for _, o := range opts {
		o.apply(&opt)
	}

	err := createDefaultConfigFile(opt.file)
	if err != nil {
		return err
	}

	viper.SetConfigFile(opt.file)
	if opt.autoEnv {
		viper.SetEnvPrefix(opt.envPrefix)
		viper.SetEnvKeyReplacer(opt.replacer.Replacer)
		viper.AutomaticEnv()
	}

	for k, v := range opt.defaultValues {
		viper.SetDefault(k, v)
	}

	if opt.flags != nil {
		viper.BindPFlags(opt.flags)
	}

	if opt.watch {
		viper.WatchConfig()
	}

	return viper.ReadInConfig()
}

func createDefaultConfigFile(path string) error {
	dir, _ := filepath.Split(path)
	if err := os.MkdirAll(dir, 0750); err != nil {
		return err
	}

	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}

	defer f.Close()

	return nil
}
