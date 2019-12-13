package config

import (
	"io"
	"strings"

	"github.com/spf13/pflag"
)

type options struct {
	configType    ConfigType
	reader        io.Reader
	file          string
	envPrefix     string
	defaultValues map[string]interface{}
	flags         *pflag.FlagSet
	watch         bool
	autoEnv       bool
	enableFile    bool
	replacer      *strings.Replacer
}

type Option interface {
	apply(*options)
}

type optionFunc func(*options)

func (f optionFunc) apply(o *options) {
	f(o)
}

func WithFile(file string) Option {
	return optionFunc(func(o *options) {
		o.file = file
	})
}

func WithEnv(prefix string) Option {
	return optionFunc(func(o *options) {
		o.envPrefix = prefix
		o.autoEnv = true
	})
}

func WithPFlags(flags *pflag.FlagSet) Option {
	return optionFunc(func(o *options) {
		o.flags = flags
	})
}

func WithDefault(defaultValues map[string]interface{}) Option {
	return optionFunc(func(o *options) {
		o.defaultValues = defaultValues
	})
}

func WithWatchEnable(enable bool) Option {
	return optionFunc(func(o *options) {
		o.watch = enable
	})
}

func WithStringReplacer(replacer *strings.Replacer) Option {
	return optionFunc(func(o *options) {
		o.replacer = replacer
	})
}

func WithReader(reader io.Reader, configType ConfigType) Option {
	return optionFunc(func(o *options) {
		o.reader = reader
		o.configType = configType
	})
}
