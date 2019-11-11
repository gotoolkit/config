package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

var (
	defaultStringReplacer = NewReplacer(".", "_")
)

const (
	defaultName      = "config.json"
	defaultPath      = "./config"
	defaultEnvPrefix = "GOTOOLKIT"
)

func Setup(opts ...Option) error {

	options := options{
		name:      defaultName,
		path:      defaultPath,
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
	path := filepath.Join(options.path, options.name)
	// viper.AddConfigPath(options.path)
	viper.SetConfigFile(path)

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
	path := filepath.Join(defaultPath, defaultName)
	log.Println(path)
	dir, _ := filepath.Split(path)
	log.Println(dir)
	if err := os.MkdirAll(dir, 0750); err != nil {
		return err
	}

	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.Write([]byte("{}"))
	if err != nil {
		return err
	}
	return nil
}
