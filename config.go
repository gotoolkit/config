package config

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	*viper.Viper
}

func New(opts ...Option) (*Config, error) {
	var err error
	opt := options{
		reader:   strings.NewReader(""),
		replacer: NewReplacer(".", "_"),
	}

	for _, o := range opts {
		o.apply(&opt)
	}
	cfg := &Config{
		Viper: viper.New(),
	}
	if opt.autoEnv {
		cfg.Viper.SetEnvPrefix(opt.envPrefix)
		cfg.Viper.SetEnvKeyReplacer(opt.replacer)
		cfg.Viper.AutomaticEnv()
	}

	for k, v := range opt.defaultValues {
		cfg.Viper.SetDefault(k, v)
	}

	if opt.flags != nil {
		cfg.Viper.BindPFlags(opt.flags)
	}

	if len(opt.file) > 0 {
		err = createDefaultConfigFile(opt.file)
		if err != nil {
			return cfg, err
		}
		cfg.Viper.SetConfigFile(opt.file)

		if opt.watch {
			cfg.Viper.WatchConfig()
		}
		err = cfg.Viper.ReadInConfig()
		return cfg, err
	}

	cfg.Viper.SetConfigType(string(opt.configType))
	err = cfg.Viper.ReadConfig(opt.reader)
	return cfg, err

}

func createDefaultConfigFile(path string) error {
	dir, _ := filepath.Split(path)
	if len(dir) > 0 {
		if err := os.MkdirAll(dir, 0750); err != nil {
			return err
		}
	}

	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}

	defer f.Close()

	return nil
}
