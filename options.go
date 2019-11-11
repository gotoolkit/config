package config

import (
	"github.com/spf13/pflag"
)

type options struct {
	file          string
	envPrefix     string
	defaultValues map[string]interface{}
	flags         *pflag.FlagSet
	watch         bool
	autoEnv       bool
	enableFile    bool
	replacer      *Replacer
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

func WithAutoEnv(enable bool) Option {
	return optionFunc(func(o *options) {
		o.autoEnv = enable
	})
}

func WithStringReplacer(replacer *Replacer) Option {
	return optionFunc(func(o *options) {
		o.replacer = replacer
	})
}
