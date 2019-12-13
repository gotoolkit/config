package config

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

func New(opts ...Option) (*viper.Viper, error) {
	var err error
	opt := options{
		reader:   strings.NewReader(""),
		replacer: NewReplacer(".", "_"),
	}

	for _, o := range opts {
		o.apply(&opt)
	}
	vp := viper.New()
	if opt.autoEnv {
		vp.SetEnvPrefix(opt.envPrefix)
		vp.SetEnvKeyReplacer(opt.replacer)
		vp.AutomaticEnv()
	}

	for k, v := range opt.defaultValues {
		vp.SetDefault(k, v)
	}

	if opt.flags != nil {
		vp.BindPFlags(opt.flags)
	}

	if opt.watch {
		vp.WatchConfig()
	}

	if len(opt.file) > 0 {
		err = createDefaultConfigFile(opt.file)
		if err != nil {
			return nil, err
		}
		vp.SetConfigFile(opt.file)
		err = vp.ReadInConfig()
		return vp, err
	}
	vp.SetConfigType(string(opt.configType))
	err = vp.ReadConfig(opt.reader)
	return vp, err

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
